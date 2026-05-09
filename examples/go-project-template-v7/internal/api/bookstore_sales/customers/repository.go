package customers

import (
	"context"
	"fmt"
	"go-project-template-v7/pkg/pageable"
	"go-project-template-v7/pkg/storage/postgres"

	"github.com/jackc/pgx/v5"
)

type Repository interface {
	Save(ctx context.Context, inputEntity *Customers) (*Customers, error)
	UpdateByID(ctx context.Context, inputEntity *Customers, pkCustomerID string) (*Customers, error)
	DeleteByID(ctx context.Context, pkCustomerID string) error
	FindByID(ctx context.Context, pkCustomerID string) (*Customers, error)
	FindAll(ctx context.Context) ([]Customers, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]Customers, pageable.Page, error)
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

func (r *repo) Save(ctx context.Context, inputEntity *Customers) (*Customers, error) {
	tag := "repository.Save"

	query := `		
		insert into bookstore_sales.customers (
			customer_id,
			email,
			full_name,
			phone,
			marketing_opt_in,
			registered_at
		)
		values ($1, $2, $3, $4, $5, $6)
		returning
			customer_id,
			email,
			full_name,
			phone,
			marketing_opt_in,
			registered_at
		`

	row := r.db.Pool.QueryRow(
		ctx, query,
		inputEntity.CustomerID,
		inputEntity.Email,
		inputEntity.FullName,
		inputEntity.Phone,
		inputEntity.MarketingOptIn,
		inputEntity.RegisteredAt,
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *repo) UpdateByID(ctx context.Context, inputEntity *Customers, pkCustomerID string) (*Customers, error) {
	tag := "repository.UpdateByID"

	query := `		
		update bookstore_sales.customers
		set
			email            = coalesce(nullif($2, ''), email),
			full_name        = coalesce(nullif($3, ''), full_name),
			phone            = coalesce(nullif($4, ''), phone),
			marketing_opt_in = coalesce(nullif($5, 'false'::bool), marketing_opt_in),
			registered_at    = coalesce(nullif($6, '0001-01-01 00:00:00'::timestamptz), registered_at)
		where customer_id = $1
		returning
			customer_id,
			email,
			full_name,
			phone,
			marketing_opt_in,
			registered_at
		`

	row := r.db.Pool.QueryRow(
		ctx, query,
		pkCustomerID,
		inputEntity.Email,
		inputEntity.FullName,
		inputEntity.Phone,
		inputEntity.MarketingOptIn,
		inputEntity.RegisteredAt,
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *repo) DeleteByID(ctx context.Context, pkCustomerID string) error {
	tag := "repository.DeleteByID"

	query := `		
		delete from only bookstore_sales.customers
		where customer_id = $1
		`

	cmdTag, err := r.db.Pool.Exec(ctx, query, pkCustomerID)
	if err != nil || cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("%s. no rows were deleted: %w", tag, err)
	}
	return nil
}

func (r *repo) FindByID(ctx context.Context, pkCustomerID string) (*Customers, error) {
	tag := "repository.FindByID"

	query := `		
		select
			customer_id,
			email,
			full_name,
			phone,
			marketing_opt_in,
			registered_at
		from bookstore_sales.customers
		where customer_id = $1
		order by customer_id
		`

	row := r.db.Pool.QueryRow(ctx, query, pkCustomerID)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *repo) FindAll(ctx context.Context) ([]Customers, error) {
	tag := "repository.FindAll"

	query := `		
		select
			customer_id,
			email,
			full_name,
			phone,
			marketing_opt_in,
			registered_at
		from bookstore_sales.customers
		order by customer_id
		`

	rows, err := r.db.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	var scannedEntities []Customers
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

func (r *repo) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]Customers, pageable.Page, error) {
	tag := "repository.FindAllPageable"

	// retrieve total count
	queryCnt := `select count(*) from bookstore_sales.customers`
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
			customer_id,
			email,
			full_name,
			phone,
			marketing_opt_in,
			registered_at
		from bookstore_sales.customers
		order by customer_id
		offset $1 limit $2
		`

	rows, err := r.db.Pool.Query(ctx, query, pq.GetOffset(), pq.GetLimit())
	if err != nil {
		return nil, pageable.Page{}, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	var scannedEntities []Customers
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
func scanFullRow(row pgx.Row) (*Customers, error) {
	var scannedEntity Customers
	err := row.Scan(
		&scannedEntity.CustomerID,
		&scannedEntity.Email,
		&scannedEntity.FullName,
		&scannedEntity.Phone,
		&scannedEntity.MarketingOptIn,
		&scannedEntity.RegisteredAt,
	)
	if err != nil {
		return nil, err
	}
	return &scannedEntity, nil
}
