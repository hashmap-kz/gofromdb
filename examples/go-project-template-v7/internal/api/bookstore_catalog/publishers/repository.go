package publishers

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
	Save(ctx context.Context, inputEntity *Publishers) (*Publishers, error)
	UpdateByID(ctx context.Context, update *UpdateDto, pkCode string) (*Publishers, error)
	DeleteByID(ctx context.Context, pkCode string) error
	FindByID(ctx context.Context, pkCode string) (*Publishers, error)
	FindAll(ctx context.Context) ([]Publishers, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]Publishers, pageable.Page, error)
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

func (r *repo) Save(ctx context.Context, inputEntity *Publishers) (*Publishers, error) {
	tag := "repository.Save"

	query := `		
		insert into bookstore_catalog.publishers (
			code,
			name,
			country_code,
			website,
			founded_on,
			active
		)
		values ($1, $2, $3, $4, $5, $6)
		returning
			code,
			name,
			country_code,
			website,
			founded_on,
			active
		`

	row := r.db.Pool.QueryRow(
		ctx, query,
		inputEntity.Code,
		inputEntity.Name,
		inputEntity.CountryCode,
		inputEntity.Website,
		inputEntity.FoundedOn,
		inputEntity.Active,
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *repo) UpdateByID(ctx context.Context, update *UpdateDto, pkCode string) (*Publishers, error) {
	tag := "repository.UpdateByID"

	if update == nil {
		return nil, fmt.Errorf("%s: update is nil", tag)
	}

	query := `		
		update bookstore_catalog.publishers
		set
			name         = coalesce($2, name),
			country_code = coalesce($3, country_code),
			website      = coalesce($4, website),
			founded_on   = coalesce($5, founded_on),
			active       = coalesce($6, active)
		where code = $1
		returning
			code,
			name,
			country_code,
			website,
			founded_on,
			active
		`

	row := r.db.Pool.QueryRow(
		ctx, query,
		pkCode,
		update.Name,
		update.CountryCode,
		update.Website,
		update.FoundedOn,
		update.Active,
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *repo) DeleteByID(ctx context.Context, pkCode string) error {
	tag := "repository.DeleteByID"

	query := `		
		delete from only bookstore_catalog.publishers
		where code = $1
		`

	cmdTag, err := r.db.Pool.Exec(ctx, query, pkCode)
	if err != nil {
		return fmt.Errorf("%s: %w", tag, err)
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("%s: %w", tag, apperrors.ErrNotFound)
	}
	return nil
}

func (r *repo) FindByID(ctx context.Context, pkCode string) (*Publishers, error) {
	tag := "repository.FindByID"

	query := `		
		select
			code,
			name,
			country_code,
			website,
			founded_on,
			active
		from bookstore_catalog.publishers
		where code = $1
		order by code
		`

	row := r.db.Pool.QueryRow(ctx, query, pkCode)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *repo) FindAll(ctx context.Context) ([]Publishers, error) {
	tag := "repository.FindAll"

	query := `		
		select
			code,
			name,
			country_code,
			website,
			founded_on,
			active
		from bookstore_catalog.publishers
		order by code
		`

	rows, err := r.db.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	scannedEntities := make([]Publishers, 0)
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

func (r *repo) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]Publishers, pageable.Page, error) {
	tag := "repository.FindAllPageable"

	// retrieve total count
	queryCnt := `select count(*) from bookstore_catalog.publishers`
	var totalCount int
	if err := r.db.Pool.QueryRow(ctx, queryCnt).Scan(&totalCount); err != nil {
		return nil, pageable.Page{}, err
	}

	// init page
	page := pageable.CreatePage(pq, totalCount)

	// handle empty result
	if totalCount == 0 {
		return make([]Publishers, 0), page, nil
	}

	// select entities
	query := `		
		select
			code,
			name,
			country_code,
			website,
			founded_on,
			active
		from bookstore_catalog.publishers
		order by code
		offset $1 limit $2
		`

	rows, err := r.db.Pool.Query(ctx, query, pq.GetOffset(), pq.GetLimit())
	if err != nil {
		return nil, pageable.Page{}, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	scannedEntities := make([]Publishers, 0)
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
func scanFullRow(row pgx.Row) (*Publishers, error) {
	var scannedEntity Publishers
	err := row.Scan(
		&scannedEntity.Code,
		&scannedEntity.Name,
		&scannedEntity.CountryCode,
		&scannedEntity.Website,
		&scannedEntity.FoundedOn,
		&scannedEntity.Active,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperrors.ErrNotFound
		}
		return nil, err
	}
	return &scannedEntity, nil
}
