package stock_events

import (
	"context"
	"fmt"
	"go-project-template-v7/pkg/pageable"
	"go-project-template-v7/pkg/storage/postgres"

	"github.com/jackc/pgx/v5"
)

type Repository interface {
	Save(ctx context.Context, inputEntity *StockEvents) (*StockEvents, error)
	FindAll(ctx context.Context) ([]StockEvents, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]StockEvents, pageable.Page, error)
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

func (r *repo) Save(ctx context.Context, inputEntity *StockEvents) (*StockEvents, error) {
	tag := "repository.Save"

	query := `		
		insert into bookstore_inventory.stock_events (
			happened_at,
			warehouse_code,
			book_id,
			delta_qty,
			reason,
			payload
		)
		values ($1, $2, $3, $4, $5, $6)
		returning
			happened_at,
			warehouse_code,
			book_id,
			delta_qty,
			reason,
			payload
		`

	row := r.db.Pool.QueryRow(
		ctx, query,
		inputEntity.HappenedAt,
		inputEntity.WarehouseCode,
		inputEntity.BookID,
		inputEntity.DeltaQty,
		inputEntity.Reason,
		inputEntity.Payload,
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *repo) FindAll(ctx context.Context) ([]StockEvents, error) {
	tag := "repository.FindAll"

	query := `		
		select
			happened_at,
			warehouse_code,
			book_id,
			delta_qty,
			reason,
			payload
		from bookstore_inventory.stock_events
		
		`

	rows, err := r.db.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	var scannedEntities []StockEvents
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

func (r *repo) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]StockEvents, pageable.Page, error) {
	tag := "repository.FindAllPageable"

	// retrieve total count
	queryCnt := `select count(*) from bookstore_inventory.stock_events`
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
			happened_at,
			warehouse_code,
			book_id,
			delta_qty,
			reason,
			payload
		from bookstore_inventory.stock_events
		
		offset $1 limit $2
		`

	rows, err := r.db.Pool.Query(ctx, query, pq.GetOffset(), pq.GetLimit())
	if err != nil {
		return nil, pageable.Page{}, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	var scannedEntities []StockEvents
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
func scanFullRow(row pgx.Row) (*StockEvents, error) {
	var scannedEntity StockEvents
	err := row.Scan(
		&scannedEntity.HappenedAt,
		&scannedEntity.WarehouseCode,
		&scannedEntity.BookID,
		&scannedEntity.DeltaQty,
		&scannedEntity.Reason,
		&scannedEntity.Payload,
	)
	if err != nil {
		return nil, err
	}
	return &scannedEntity, nil
}
