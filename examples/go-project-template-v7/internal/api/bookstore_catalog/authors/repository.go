package authors

import (
	"context"
	"fmt"
	"go-project-template-v7/pkg/pageable"
	"go-project-template-v7/pkg/storage/postgres"

	"github.com/jackc/pgx/v5"
)

type Repository interface {
	Save(ctx context.Context, inputEntity *Authors) (*Authors, error)
	UpdateByID(ctx context.Context, inputEntity *Authors, pkAuthorID string) (*Authors, error)
	DeleteByID(ctx context.Context, pkAuthorID string) error
	FindByID(ctx context.Context, pkAuthorID string) (*Authors, error)
	FindAll(ctx context.Context) ([]Authors, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]Authors, pageable.Page, error)
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

func (r *repo) Save(ctx context.Context, inputEntity *Authors) (*Authors, error) {
	tag := "repository.Save"

	query := `		
		insert into bookstore_catalog.authors (
			author_id,
			display_name,
			legal_name,
			biography,
			metadata,
			active,
			born_on
		)
		values ($1, $2, $3, $4, $5, $6, $7)
		returning
			author_id,
			display_name,
			legal_name,
			biography,
			metadata,
			active,
			born_on
		`

	row := r.db.Pool.QueryRow(
		ctx, query,
		inputEntity.AuthorID,
		inputEntity.DisplayName,
		inputEntity.LegalName,
		inputEntity.Biography,
		inputEntity.Metadata,
		inputEntity.Active,
		inputEntity.BornOn,
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *repo) UpdateByID(ctx context.Context, inputEntity *Authors, pkAuthorID string) (*Authors, error) {
	tag := "repository.UpdateByID"

	query := `		
		update bookstore_catalog.authors
		set
			display_name = coalesce(nullif($2, ''), display_name),
			legal_name   = coalesce(nullif($3, ''), legal_name),
			biography    = coalesce(nullif($4, ''), biography),
			metadata     = coalesce(nullif($5, '{}'::jsonb), metadata),
			active       = coalesce(nullif($6, 'false'::bool), active),
			born_on      = coalesce(nullif($7, '0001-01-01 00:00:00'::date), born_on)
		where author_id = $1
		returning
			author_id,
			display_name,
			legal_name,
			biography,
			metadata,
			active,
			born_on
		`

	row := r.db.Pool.QueryRow(
		ctx, query,
		pkAuthorID,
		inputEntity.DisplayName,
		inputEntity.LegalName,
		inputEntity.Biography,
		inputEntity.Metadata,
		inputEntity.Active,
		inputEntity.BornOn,
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *repo) DeleteByID(ctx context.Context, pkAuthorID string) error {
	tag := "repository.DeleteByID"

	query := `		
		delete from only bookstore_catalog.authors
		where author_id = $1
		`

	cmdTag, err := r.db.Pool.Exec(ctx, query, pkAuthorID)
	if err != nil || cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("%s. no rows were deleted: %w", tag, err)
	}
	return nil
}

func (r *repo) FindByID(ctx context.Context, pkAuthorID string) (*Authors, error) {
	tag := "repository.FindByID"

	query := `		
		select
			author_id,
			display_name,
			legal_name,
			biography,
			metadata,
			active,
			born_on
		from bookstore_catalog.authors
		where author_id = $1
		order by author_id
		`

	row := r.db.Pool.QueryRow(ctx, query, pkAuthorID)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *repo) FindAll(ctx context.Context) ([]Authors, error) {
	tag := "repository.FindAll"

	query := `		
		select
			author_id,
			display_name,
			legal_name,
			biography,
			metadata,
			active,
			born_on
		from bookstore_catalog.authors
		order by author_id
		`

	rows, err := r.db.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	var scannedEntities []Authors
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

func (r *repo) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]Authors, pageable.Page, error) {
	tag := "repository.FindAllPageable"

	// retrieve total count
	queryCnt := `select count(*) from bookstore_catalog.authors`
	var totalCount int
	if err := r.db.Pool.QueryRow(ctx, queryCnt).Scan(&totalCount); err != nil {
		return nil, pageable.Page{}, err
	}

	// init page
	page := pageable.CreatePage(pq, totalCount)

	// handle empty result
	if totalCount == 0 {
		return nil, page, nil
	}

	// select entities
	query := `		
		select
			author_id,
			display_name,
			legal_name,
			biography,
			metadata,
			active,
			born_on
		from bookstore_catalog.authors
		order by author_id
		offset $1 limit $2
		`

	rows, err := r.db.Pool.Query(ctx, query, pq.GetOffset(), pq.GetLimit())
	if err != nil {
		return nil, pageable.Page{}, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	var scannedEntities []Authors
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
func scanFullRow(row pgx.Row) (*Authors, error) {
	var scannedEntity Authors
	err := row.Scan(
		&scannedEntity.AuthorID,
		&scannedEntity.DisplayName,
		&scannedEntity.LegalName,
		&scannedEntity.Biography,
		&scannedEntity.Metadata,
		&scannedEntity.Active,
		&scannedEntity.BornOn,
	)
	if err != nil {
		return nil, err
	}
	return &scannedEntity, nil
}
