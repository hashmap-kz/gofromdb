package stock_levels

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
	Save(ctx context.Context, inputEntity *StockLevels) (*StockLevels, error)
	UpdateByID(ctx context.Context, inputEntity *StockLevels, pkWarehouseCode string, pkBookID int64) (*StockLevels, error)
	DeleteByID(ctx context.Context, pkWarehouseCode string, pkBookID int64) error
	FindByID(ctx context.Context, pkWarehouseCode string, pkBookID int64) (*StockLevels, error)
	FindAll(ctx context.Context) ([]StockLevels, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]StockLevels, pageable.Page, error)
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

func (r *repo) Save(ctx context.Context, inputEntity *StockLevels) (*StockLevels, error) {
	tag := "repository.Save"

	query := `		
		insert into bookstore_inventory.stock_levels (
			warehouse_code,
			book_id,
			available_qty,
			reserved_qty,
			reorder_threshold,
			last_counted_at
		)
		values ($1, $2, $3, $4, $5, $6)
		returning
			warehouse_code,
			book_id,
			available_qty,
			reserved_qty,
			reorder_threshold,
			last_counted_at
		`

	row := r.db.Pool.QueryRow(
		ctx, query,
		inputEntity.WarehouseCode,
		inputEntity.BookID,
		inputEntity.AvailableQty,
		inputEntity.ReservedQty,
		inputEntity.ReorderThreshold,
		inputEntity.LastCountedAt,
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *repo) UpdateByID(ctx context.Context, inputEntity *StockLevels, pkWarehouseCode string, pkBookID int64) (*StockLevels, error) {
	tag := "repository.UpdateByID"

	query := `		
		update bookstore_inventory.stock_levels
		set
			available_qty     = coalesce(nullif($3, 0::int4), available_qty),
			reserved_qty      = coalesce(nullif($4, 0::int4), reserved_qty),
			reorder_threshold = coalesce(nullif($5, 0::int4), reorder_threshold),
			last_counted_at   = coalesce(nullif($6, '0001-01-01 00:00:00'::timestamp), last_counted_at)
		where warehouse_code = $1 and book_id = $2
		returning
			warehouse_code,
			book_id,
			available_qty,
			reserved_qty,
			reorder_threshold,
			last_counted_at
		`

	row := r.db.Pool.QueryRow(
		ctx, query,
		pkWarehouseCode, pkBookID,
		inputEntity.AvailableQty,
		inputEntity.ReservedQty,
		inputEntity.ReorderThreshold,
		inputEntity.LastCountedAt,
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *repo) DeleteByID(ctx context.Context, pkWarehouseCode string, pkBookID int64) error {
	tag := "repository.DeleteByID"

	query := `		
		delete from only bookstore_inventory.stock_levels
		where warehouse_code = $1 and book_id = $2
		`

	cmdTag, err := r.db.Pool.Exec(ctx, query, pkWarehouseCode, pkBookID)
	if err != nil {
		return fmt.Errorf("%s: %w", tag, err)
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("%s: %w", tag, apperrors.ErrNotFound)
	}
	return nil
}

func (r *repo) FindByID(ctx context.Context, pkWarehouseCode string, pkBookID int64) (*StockLevels, error) {
	tag := "repository.FindByID"

	query := `		
		select
			warehouse_code,
			book_id,
			available_qty,
			reserved_qty,
			reorder_threshold,
			last_counted_at
		from bookstore_inventory.stock_levels
		where warehouse_code = $1 and book_id = $2
		order by warehouse_code, book_id
		`

	row := r.db.Pool.QueryRow(ctx, query, pkWarehouseCode, pkBookID)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *repo) FindAll(ctx context.Context) ([]StockLevels, error) {
	tag := "repository.FindAll"

	query := `		
		select
			warehouse_code,
			book_id,
			available_qty,
			reserved_qty,
			reorder_threshold,
			last_counted_at
		from bookstore_inventory.stock_levels
		order by warehouse_code, book_id
		`

	rows, err := r.db.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	scannedEntities := make([]StockLevels, 0)
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

func (r *repo) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]StockLevels, pageable.Page, error) {
	tag := "repository.FindAllPageable"

	// retrieve total count
	queryCnt := `select count(*) from bookstore_inventory.stock_levels`
	var totalCount int
	if err := r.db.Pool.QueryRow(ctx, queryCnt).Scan(&totalCount); err != nil {
		return nil, pageable.Page{}, err
	}

	// init page
	page := pageable.CreatePage(pq, totalCount)

	// handle empty result
	if totalCount == 0 {
		return make([]StockLevels, 0), page, nil
	}

	// select entities
	query := `		
		select
			warehouse_code,
			book_id,
			available_qty,
			reserved_qty,
			reorder_threshold,
			last_counted_at
		from bookstore_inventory.stock_levels
		order by warehouse_code, book_id
		offset $1 limit $2
		`

	rows, err := r.db.Pool.Query(ctx, query, pq.GetOffset(), pq.GetLimit())
	if err != nil {
		return nil, pageable.Page{}, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	scannedEntities := make([]StockLevels, 0)
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
func scanFullRow(row pgx.Row) (*StockLevels, error) {
	var scannedEntity StockLevels
	err := row.Scan(
		&scannedEntity.WarehouseCode,
		&scannedEntity.BookID,
		&scannedEntity.AvailableQty,
		&scannedEntity.ReservedQty,
		&scannedEntity.ReorderThreshold,
		&scannedEntity.LastCountedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperrors.ErrNotFound
		}
		return nil, err
	}
	return &scannedEntity, nil
}
