package warehouses

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
	Save(ctx context.Context, inputEntity *Warehouses) (*Warehouses, error)
	UpdateByID(ctx context.Context, update *UpdateDto, pkCode string) (*Warehouses, error)
	DeleteByID(ctx context.Context, pkCode string) error
	FindByID(ctx context.Context, pkCode string) (*Warehouses, error)
	FindAll(ctx context.Context) ([]Warehouses, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]Warehouses, pageable.Page, error)
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

func (r *repo) Save(ctx context.Context, inputEntity *Warehouses) (*Warehouses, error) {
	tag := "repository.Save"

	query := `		
		insert into bookstore_inventory.warehouses (
			code,
			name,
			address,
			timezone,
			active
		)
		values ($1, $2, $3, $4, $5)
		returning
			code,
			name,
			address,
			timezone,
			active
		`

	row := r.db.Pool.QueryRow(
		ctx, query,
		inputEntity.Code,
		inputEntity.Name,
		inputEntity.Address,
		inputEntity.Timezone,
		inputEntity.Active,
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *repo) UpdateByID(ctx context.Context, update *UpdateDto, pkCode string) (*Warehouses, error) {
	tag := "repository.UpdateByID"

	if update == nil {
		return nil, fmt.Errorf("%s: update is nil", tag)
	}

	query := `		
		update bookstore_inventory.warehouses
		set
			name     = coalesce($2, name),
			address  = coalesce($3, address),
			timezone = coalesce($4, timezone),
			active   = coalesce($5, active)
		where code = $1
		returning
			code,
			name,
			address,
			timezone,
			active
		`

	row := r.db.Pool.QueryRow(
		ctx, query,
		pkCode,
		update.Name,
		update.Address,
		update.Timezone,
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
		delete from only bookstore_inventory.warehouses
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

func (r *repo) FindByID(ctx context.Context, pkCode string) (*Warehouses, error) {
	tag := "repository.FindByID"

	query := `		
		select
			code,
			name,
			address,
			timezone,
			active
		from bookstore_inventory.warehouses
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

func (r *repo) FindAll(ctx context.Context) ([]Warehouses, error) {
	tag := "repository.FindAll"

	query := `		
		select
			code,
			name,
			address,
			timezone,
			active
		from bookstore_inventory.warehouses
		order by code
		`

	rows, err := r.db.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	scannedEntities := make([]Warehouses, 0)
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

func (r *repo) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]Warehouses, pageable.Page, error) {
	tag := "repository.FindAllPageable"

	// retrieve total count
	queryCnt := `select count(*) from bookstore_inventory.warehouses`
	var totalCount int
	if err := r.db.Pool.QueryRow(ctx, queryCnt).Scan(&totalCount); err != nil {
		return nil, pageable.Page{}, err
	}

	// init page
	page := pageable.CreatePage(pq, totalCount)

	// handle empty result
	if totalCount == 0 {
		return make([]Warehouses, 0), page, nil
	}

	// select entities
	query := `		
		select
			code,
			name,
			address,
			timezone,
			active
		from bookstore_inventory.warehouses
		order by code
		offset $1 limit $2
		`

	rows, err := r.db.Pool.Query(ctx, query, pq.GetOffset(), pq.GetLimit())
	if err != nil {
		return nil, pageable.Page{}, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	scannedEntities := make([]Warehouses, 0)
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
func scanFullRow(row pgx.Row) (*Warehouses, error) {
	var scannedEntity Warehouses
	err := row.Scan(
		&scannedEntity.Code,
		&scannedEntity.Name,
		&scannedEntity.Address,
		&scannedEntity.Timezone,
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
