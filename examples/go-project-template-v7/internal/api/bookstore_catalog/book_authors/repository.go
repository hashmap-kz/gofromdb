package book_authors

import (
	"context"
	"errors"
	"fmt"
	"go-project-template-v7/pkg/apperrors"
	"go-project-template-v7/pkg/pageable"
	"go-project-template-v7/pkg/storage/postgres"

	"github.com/jackc/pgx/v5"
)

type Repository interface {
	Save(ctx context.Context, inputEntity *BookAuthors) (*BookAuthors, error)
	UpdateByID(ctx context.Context, update *UpdateDto, pkBookID int64, pkAuthorID string) (*BookAuthors, error)
	DeleteByID(ctx context.Context, pkBookID int64, pkAuthorID string) error
	FindByID(ctx context.Context, pkBookID int64, pkAuthorID string) (*BookAuthors, error)
	FindAll(ctx context.Context) ([]BookAuthors, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]BookAuthors, pageable.Page, error)
}

type repo struct {
	db *postgres.Postgres
}

var _ Repository = &repo{}

func NewRepository(_ context.Context, db *postgres.Postgres) Repository {
	return &repo{
		db: db,
	}
}

func (r *repo) Save(ctx context.Context, inputEntity *BookAuthors) (*BookAuthors, error) {
	tag := "repository.Save"

	query := `		
		insert into bookstore_catalog.book_authors (
			book_id,
			author_id,
			contribution_order,
			role,
			notes
		)
		values ($1, $2, $3, $4, $5)
		returning
			book_id,
			author_id,
			contribution_order,
			role,
			notes
		`

	row := r.db.Pool.QueryRow(
		ctx, query,
		inputEntity.BookID,
		inputEntity.AuthorID,
		inputEntity.ContributionOrder,
		inputEntity.Role,
		inputEntity.Notes,
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *repo) UpdateByID(ctx context.Context, update *UpdateDto, pkBookID int64, pkAuthorID string) (*BookAuthors, error) {
	tag := "repository.UpdateByID"

	if update == nil {
		return nil, fmt.Errorf("%s: update is nil", tag)
	}

	query := `		
		update bookstore_catalog.book_authors
		set
			contribution_order = coalesce($3, contribution_order),
			role               = coalesce($4, role),
			notes              = coalesce($5, notes)
		where book_id = $1 and author_id = $2
		returning
			book_id,
			author_id,
			contribution_order,
			role,
			notes
		`

	row := r.db.Pool.QueryRow(
		ctx, query,
		pkBookID, pkAuthorID,
		update.ContributionOrder,
		update.Role,
		update.Notes,
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *repo) DeleteByID(ctx context.Context, pkBookID int64, pkAuthorID string) error {
	tag := "repository.DeleteByID"

	query := `		
		delete from only bookstore_catalog.book_authors
		where book_id = $1 and author_id = $2
		`

	cmdTag, err := r.db.Pool.Exec(ctx, query, pkBookID, pkAuthorID)
	if err != nil {
		return fmt.Errorf("%s: %w", tag, err)
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("%s: %w", tag, apperrors.ErrNotFound)
	}
	return nil
}

func (r *repo) FindByID(ctx context.Context, pkBookID int64, pkAuthorID string) (*BookAuthors, error) {
	tag := "repository.FindByID"

	query := `		
		select
			book_id,
			author_id,
			contribution_order,
			role,
			notes
		from bookstore_catalog.book_authors
		where book_id = $1 and author_id = $2
		order by book_id, author_id
		`

	row := r.db.Pool.QueryRow(ctx, query, pkBookID, pkAuthorID)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *repo) FindAll(ctx context.Context) ([]BookAuthors, error) {
	tag := "repository.FindAll"

	query := `		
		select
			book_id,
			author_id,
			contribution_order,
			role,
			notes
		from bookstore_catalog.book_authors
		order by book_id, author_id
		`

	rows, err := r.db.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	scannedEntities := make([]BookAuthors, 0)
	for rows.Next() {
		scannedEntity, err := scanFullRow(rows)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", tag, err)
		}
		scannedEntities = append(scannedEntities, *scannedEntity)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return scannedEntities, nil
}

func (r *repo) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]BookAuthors, pageable.Page, error) {
	tag := "repository.FindAllPageable"

	// retrieve total count
	queryCnt := `select count(*) from bookstore_catalog.book_authors`
	var totalCount int
	if err := r.db.Pool.QueryRow(ctx, queryCnt).Scan(&totalCount); err != nil {
		return nil, pageable.Page{}, err
	}

	// init page
	page := pageable.CreatePage(pq, totalCount)

	// handle empty result
	if totalCount == 0 {
		return make([]BookAuthors, 0), page, nil
	}

	// select entities
	query := `		
		select
			book_id,
			author_id,
			contribution_order,
			role,
			notes
		from bookstore_catalog.book_authors
		order by book_id, author_id
		offset $1 limit $2
		`

	rows, err := r.db.Pool.Query(ctx, query, pq.GetOffset(), pq.GetLimit())
	if err != nil {
		return nil, pageable.Page{}, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	scannedEntities := make([]BookAuthors, 0)
	for rows.Next() {
		scannedEntity, err := scanFullRow(rows)
		if err != nil {
			return nil, pageable.Page{}, fmt.Errorf("%s: %w", tag, err)
		}
		scannedEntities = append(scannedEntities, *scannedEntity)
	}

	if rows.Err() != nil {
		return nil, page, rows.Err()
	}
	return scannedEntities, page, nil
}

// scan utils

// scanFullRow is expected to scan all columns from a table.
// For simplicity, most methods scan the entire row of the table into the result entity.
// You should adapt methods as needed (e.g., if business logic requires returning only an ID after an UPDATE).
func scanFullRow(row pgx.Row) (*BookAuthors, error) {
	var scannedEntity BookAuthors
	err := row.Scan(
		&scannedEntity.BookID,
		&scannedEntity.AuthorID,
		&scannedEntity.ContributionOrder,
		&scannedEntity.Role,
		&scannedEntity.Notes,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperrors.ErrNotFound
		}
		return nil, err
	}
	return &scannedEntity, nil
}
