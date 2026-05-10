package discount_codes

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
	Save(ctx context.Context, inputEntity *DiscountCodes) (*DiscountCodes, error)
	UpdateByID(ctx context.Context, inputEntity *DiscountCodes, pkCode string) (*DiscountCodes, error)
	DeleteByID(ctx context.Context, pkCode string) error
	FindByID(ctx context.Context, pkCode string) (*DiscountCodes, error)
	FindAll(ctx context.Context) ([]DiscountCodes, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]DiscountCodes, pageable.Page, error)
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

func (r *repo) Save(ctx context.Context, inputEntity *DiscountCodes) (*DiscountCodes, error) {
	tag := "repository.Save"

	query := `		
		insert into bookstore_sales.discount_codes (
			code,
			description,
			percent_off,
			valid_period,
			max_uses,
			active
		)
		values ($1, $2, $3, $4, $5, $6)
		returning
			code,
			description,
			percent_off,
			valid_period,
			max_uses,
			active
		`

	row := r.db.Pool.QueryRow(
		ctx, query,
		inputEntity.Code,
		inputEntity.Description,
		inputEntity.PercentOff,
		inputEntity.ValidPeriod,
		inputEntity.MaxUses,
		inputEntity.Active,
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *repo) UpdateByID(ctx context.Context, inputEntity *DiscountCodes, pkCode string) (*DiscountCodes, error) {
	tag := "repository.UpdateByID"

	query := `		
		update bookstore_sales.discount_codes
		set
			description  = coalesce(nullif($2, ''), description),
			percent_off  = coalesce(nullif($3, 0::numeric), percent_off),
			valid_period = coalesce(nullif($4, 'empty'::daterange), valid_period),
			max_uses     = coalesce(nullif($5, 0::int4), max_uses),
			active       = coalesce(nullif($6, 'false'::bool), active)
		where code = $1
		returning
			code,
			description,
			percent_off,
			valid_period,
			max_uses,
			active
		`

	row := r.db.Pool.QueryRow(
		ctx, query,
		pkCode,
		inputEntity.Description,
		inputEntity.PercentOff,
		inputEntity.ValidPeriod,
		inputEntity.MaxUses,
		inputEntity.Active,
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
		delete from only bookstore_sales.discount_codes
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

func (r *repo) FindByID(ctx context.Context, pkCode string) (*DiscountCodes, error) {
	tag := "repository.FindByID"

	query := `		
		select
			code,
			description,
			percent_off,
			valid_period,
			max_uses,
			active
		from bookstore_sales.discount_codes
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

func (r *repo) FindAll(ctx context.Context) ([]DiscountCodes, error) {
	tag := "repository.FindAll"

	query := `		
		select
			code,
			description,
			percent_off,
			valid_period,
			max_uses,
			active
		from bookstore_sales.discount_codes
		order by code
		`

	rows, err := r.db.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	scannedEntities := make([]DiscountCodes, 0)
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

func (r *repo) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]DiscountCodes, pageable.Page, error) {
	tag := "repository.FindAllPageable"

	// retrieve total count
	queryCnt := `select count(*) from bookstore_sales.discount_codes`
	var totalCount int
	if err := r.db.Pool.QueryRow(ctx, queryCnt).Scan(&totalCount); err != nil {
		return nil, pageable.Page{}, err
	}

	// init page
	page := pageable.CreatePage(pq, totalCount)

	// handle empty result
	if totalCount == 0 {
		return make([]DiscountCodes, 0), page, nil
	}

	// select entities
	query := `		
		select
			code,
			description,
			percent_off,
			valid_period,
			max_uses,
			active
		from bookstore_sales.discount_codes
		order by code
		offset $1 limit $2
		`

	rows, err := r.db.Pool.Query(ctx, query, pq.GetOffset(), pq.GetLimit())
	if err != nil {
		return nil, pageable.Page{}, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	scannedEntities := make([]DiscountCodes, 0)
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
func scanFullRow(row pgx.Row) (*DiscountCodes, error) {
	var scannedEntity DiscountCodes
	err := row.Scan(
		&scannedEntity.Code,
		&scannedEntity.Description,
		&scannedEntity.PercentOff,
		&scannedEntity.ValidPeriod,
		&scannedEntity.MaxUses,
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
