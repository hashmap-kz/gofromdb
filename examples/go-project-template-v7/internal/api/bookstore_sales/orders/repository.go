package orders

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
	Save(ctx context.Context, inputEntity *BookstoreSalesOrders) (*BookstoreSalesOrders, error)
	UpdateByID(ctx context.Context, inputEntity *BookstoreSalesOrders, pkOrderID int64) (*BookstoreSalesOrders, error)
	DeleteByID(ctx context.Context, pkOrderID int64) error
	FindByID(ctx context.Context, pkOrderID int64) (*BookstoreSalesOrders, error)
	FindAll(ctx context.Context) ([]BookstoreSalesOrders, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]BookstoreSalesOrders, pageable.Page, error)
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

func (r *repo) Save(ctx context.Context, inputEntity *BookstoreSalesOrders) (*BookstoreSalesOrders, error) {
	tag := "repository.Save"

	query := `		
		insert into bookstore_sales.orders (
			customer_id,
			status,
			placed_at,
			paid_at,
			cancelled_at,
			comment
		)
		values ($1, $2, $3, $4, $5, $6)
		returning
			order_id,
			customer_id,
			status,
			placed_at,
			paid_at,
			cancelled_at,
			comment
		`

	row := r.db.Pool.QueryRow(
		ctx, query,
		inputEntity.CustomerID,
		inputEntity.Status,
		inputEntity.PlacedAt,
		inputEntity.PaidAt,
		inputEntity.CancelledAt,
		inputEntity.Comment,
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *repo) UpdateByID(ctx context.Context, inputEntity *BookstoreSalesOrders, pkOrderID int64) (*BookstoreSalesOrders, error) {
	tag := "repository.UpdateByID"

	query := `		
		update bookstore_sales.orders
		set
			customer_id  = coalesce(nullif($2, '00000000-0000-0000-0000-000000000000'::uuid), customer_id),
			status       = coalesce(nullif($3, ''), status),
			placed_at    = coalesce(nullif($4, '0001-01-01 00:00:00'::timestamptz), placed_at),
			paid_at      = coalesce(nullif($5, '0001-01-01 00:00:00'::timestamptz), paid_at),
			cancelled_at = coalesce(nullif($6, '0001-01-01 00:00:00'::timestamptz), cancelled_at),
			comment      = coalesce(nullif($7, ''), comment)
		where order_id = $1
		returning
			order_id,
			customer_id,
			status,
			placed_at,
			paid_at,
			cancelled_at,
			comment
		`

	row := r.db.Pool.QueryRow(
		ctx, query,
		pkOrderID,
		inputEntity.CustomerID,
		inputEntity.Status,
		inputEntity.PlacedAt,
		inputEntity.PaidAt,
		inputEntity.CancelledAt,
		inputEntity.Comment,
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *repo) DeleteByID(ctx context.Context, pkOrderID int64) error {
	tag := "repository.DeleteByID"

	query := `		
		delete from only bookstore_sales.orders
		where order_id = $1
		`

	cmdTag, err := r.db.Pool.Exec(ctx, query, pkOrderID)
	if err != nil {
		return fmt.Errorf("%s: %w", tag, err)
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("%s: %w", tag, apperrors.ErrNotFound)
	}
	return nil
}

func (r *repo) FindByID(ctx context.Context, pkOrderID int64) (*BookstoreSalesOrders, error) {
	tag := "repository.FindByID"

	query := `		
		select
			order_id,
			customer_id,
			status,
			placed_at,
			paid_at,
			cancelled_at,
			comment
		from bookstore_sales.orders
		where order_id = $1
		order by order_id
		`

	row := r.db.Pool.QueryRow(ctx, query, pkOrderID)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *repo) FindAll(ctx context.Context) ([]BookstoreSalesOrders, error) {
	tag := "repository.FindAll"

	query := `		
		select
			order_id,
			customer_id,
			status,
			placed_at,
			paid_at,
			cancelled_at,
			comment
		from bookstore_sales.orders
		order by order_id
		`

	rows, err := r.db.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	scannedEntities := make([]BookstoreSalesOrders, 0)
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

func (r *repo) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]BookstoreSalesOrders, pageable.Page, error) {
	tag := "repository.FindAllPageable"

	// retrieve total count
	queryCnt := `select count(*) from bookstore_sales.orders`
	var totalCount int
	if err := r.db.Pool.QueryRow(ctx, queryCnt).Scan(&totalCount); err != nil {
		return nil, pageable.Page{}, err
	}

	// init page
	page := pageable.CreatePage(pq, totalCount)

	// handle empty result
	if totalCount == 0 {
		return make([]BookstoreSalesOrders, 0), page, nil
	}

	// select entities
	query := `		
		select
			order_id,
			customer_id,
			status,
			placed_at,
			paid_at,
			cancelled_at,
			comment
		from bookstore_sales.orders
		order by order_id
		offset $1 limit $2
		`

	rows, err := r.db.Pool.Query(ctx, query, pq.GetOffset(), pq.GetLimit())
	if err != nil {
		return nil, pageable.Page{}, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	scannedEntities := make([]BookstoreSalesOrders, 0)
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
func scanFullRow(row pgx.Row) (*BookstoreSalesOrders, error) {
	var scannedEntity BookstoreSalesOrders
	err := row.Scan(
		&scannedEntity.OrderID,
		&scannedEntity.CustomerID,
		&scannedEntity.Status,
		&scannedEntity.PlacedAt,
		&scannedEntity.PaidAt,
		&scannedEntity.CancelledAt,
		&scannedEntity.Comment,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperrors.ErrNotFound
		}
		return nil, err
	}
	return &scannedEntity, nil
}
