package order_items

import (
	"context"
	"fmt"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/pkg/storage/postgres"

	"github.com/jackc/pgx/v5"
)

type Repository interface {
	Save(ctx context.Context, inputEntity *OrderItems) (*OrderItems, error)
	UpdateByID(ctx context.Context, inputEntity *OrderItems, pkRecordID int) (*OrderItems, error)
	DeleteByID(ctx context.Context, pkRecordID int) error
	FindByID(ctx context.Context, pkRecordID int) (*OrderItems, error)
	FindAll(ctx context.Context) ([]OrderItems, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]OrderItems, pageable.Page, error)
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

func (r *repo) Save(ctx context.Context, inputEntity *OrderItems) (*OrderItems, error) {
	tag := "repository.Save"

	query := `		
		insert into public.order_items (
			order_id,
			product_id,
			quantity,
			price
		)
		values ($1, $2, $3, $4)
		returning
			record_id,
			order_id,
			product_id,
			quantity,
			price,
			created_at,
			updated_at,
			guid
		`

	row := r.db.Pool.QueryRow(ctx, query,
		inputEntity.OrderID,
		inputEntity.ProductID,
		inputEntity.Quantity,
		inputEntity.Price,
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *repo) UpdateByID(ctx context.Context, inputEntity *OrderItems, pkRecordID int) (*OrderItems, error) {
	tag := "repository.UpdateByID"

	query := `		
		update public.order_items
		set
			order_id   = coalesce(nullif($2, 0::int4), order_id),
			product_id = coalesce(nullif($3, 0::int4), product_id),
			quantity   = coalesce(nullif($4, 0::numeric), quantity),
			price      = coalesce(nullif($5, 0::numeric), price)
		where record_id = $1
		returning
			record_id,
			order_id,
			product_id,
			quantity,
			price,
			created_at,
			updated_at,
			guid
		`

	row := r.db.Pool.QueryRow(ctx, query,
		pkRecordID,
		inputEntity.OrderID,
		inputEntity.ProductID,
		inputEntity.Quantity,
		inputEntity.Price,
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *repo) DeleteByID(ctx context.Context, pkRecordID int) error {
	tag := "repository.DeleteByID"

	query := `		
		delete from only public.order_items
		where record_id = $1
		`

	cmdTag, err := r.db.Pool.Exec(ctx, query, pkRecordID)
	if err != nil || cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("%s. no rows were deleted: %w", tag, err)
	}
	return nil
}

func (r *repo) FindByID(ctx context.Context, pkRecordID int) (*OrderItems, error) {
	tag := "repository.FindByID"

	query := `		
		select
			record_id,
			order_id,
			product_id,
			quantity,
			price,
			created_at,
			updated_at,
			guid
		from public.order_items
		where record_id = $1
		order by record_id
		`

	row := r.db.Pool.QueryRow(ctx, query, pkRecordID)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *repo) FindAll(ctx context.Context) ([]OrderItems, error) {
	tag := "repository.FindAll"

	query := `		
		select
			record_id,
			order_id,
			product_id,
			quantity,
			price,
			created_at,
			updated_at,
			guid
		from public.order_items
		order by record_id
		`

	rows, err := r.db.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	var scannedEntities []OrderItems
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

func (r *repo) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]OrderItems, pageable.Page, error) {
	tag := "repository.FindAllPageable"

	// retrieve total count
	queryCnt := `select count(*) from public.order_items`
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
			record_id,
			order_id,
			product_id,
			quantity,
			price,
			created_at,
			updated_at,
			guid
		from public.order_items
		order by record_id
		offset $1 limit $2
		`

	rows, err := r.db.Pool.Query(ctx, query, pq.GetOffset(), pq.GetLimit())
	if err != nil {
		return nil, pageable.Page{}, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	var scannedEntities []OrderItems
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
func scanFullRow(row pgx.Row) (*OrderItems, error) {
	var scannedEntity OrderItems
	err := row.Scan(
		&scannedEntity.RecordID,
		&scannedEntity.OrderID,
		&scannedEntity.ProductID,
		&scannedEntity.Quantity,
		&scannedEntity.Price,
		&scannedEntity.CreatedAt,
		&scannedEntity.UpdatedAt,
		&scannedEntity.GUID,
	)
	if err != nil {
		return nil, err
	}
	return &scannedEntity, nil
}
