package order_lines

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
	Save(ctx context.Context, inputEntity *OrderLines) (*OrderLines, error)
	UpdateByID(ctx context.Context, update *UpdateDto, pkOrderID int64, pkLineNo int) (*OrderLines, error)
	DeleteByID(ctx context.Context, pkOrderID int64, pkLineNo int) error
	FindByID(ctx context.Context, pkOrderID int64, pkLineNo int) (*OrderLines, error)
	FindAll(ctx context.Context) ([]OrderLines, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]OrderLines, pageable.Page, error)
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

func (r *repo) Save(ctx context.Context, inputEntity *OrderLines) (*OrderLines, error) {
	tag := "repository.Save"

	query := `		
		insert into bookstore_sales.order_lines (
			order_id,
			line_no,
			book_id,
			quantity,
			unit_price,
			discount_amount,
			note
		)
		values ($1, $2, $3, $4, $5, $6, $7)
		returning
			order_id,
			line_no,
			book_id,
			quantity,
			unit_price,
			discount_amount,
			note
		`

	row := r.db.Pool.QueryRow(
		ctx, query,
		inputEntity.OrderID,
		inputEntity.LineNo,
		inputEntity.BookID,
		inputEntity.Quantity,
		inputEntity.UnitPrice,
		inputEntity.DiscountAmount,
		inputEntity.Note,
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *repo) UpdateByID(ctx context.Context, update *UpdateDto, pkOrderID int64, pkLineNo int) (*OrderLines, error) {
	tag := "repository.UpdateByID"

	if update == nil {
		return nil, fmt.Errorf("%s: update is nil", tag)
	}

	query := `		
		update bookstore_sales.order_lines
		set
			book_id         = coalesce($3, book_id),
			quantity        = coalesce($4, quantity),
			unit_price      = coalesce($5, unit_price),
			discount_amount = coalesce($6, discount_amount),
			note            = coalesce($7, note)
		where order_id = $1 and line_no = $2
		returning
			order_id,
			line_no,
			book_id,
			quantity,
			unit_price,
			discount_amount,
			note
		`

	row := r.db.Pool.QueryRow(
		ctx, query,
		pkOrderID, pkLineNo,
		update.BookID,
		update.Quantity,
		update.UnitPrice,
		update.DiscountAmount,
		update.Note,
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *repo) DeleteByID(ctx context.Context, pkOrderID int64, pkLineNo int) error {
	tag := "repository.DeleteByID"

	query := `		
		delete from only bookstore_sales.order_lines
		where order_id = $1 and line_no = $2
		`

	cmdTag, err := r.db.Pool.Exec(ctx, query, pkOrderID, pkLineNo)
	if err != nil {
		return fmt.Errorf("%s: %w", tag, err)
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("%s: %w", tag, apperrors.ErrNotFound)
	}
	return nil
}

func (r *repo) FindByID(ctx context.Context, pkOrderID int64, pkLineNo int) (*OrderLines, error) {
	tag := "repository.FindByID"

	query := `		
		select
			order_id,
			line_no,
			book_id,
			quantity,
			unit_price,
			discount_amount,
			note
		from bookstore_sales.order_lines
		where order_id = $1 and line_no = $2
		order by order_id, line_no
		`

	row := r.db.Pool.QueryRow(ctx, query, pkOrderID, pkLineNo)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *repo) FindAll(ctx context.Context) ([]OrderLines, error) {
	tag := "repository.FindAll"

	query := `		
		select
			order_id,
			line_no,
			book_id,
			quantity,
			unit_price,
			discount_amount,
			note
		from bookstore_sales.order_lines
		order by order_id, line_no
		`

	rows, err := r.db.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	scannedEntities := make([]OrderLines, 0)
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

func (r *repo) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]OrderLines, pageable.Page, error) {
	tag := "repository.FindAllPageable"

	// retrieve total count
	queryCnt := `select count(*) from bookstore_sales.order_lines`
	var totalCount int
	if err := r.db.Pool.QueryRow(ctx, queryCnt).Scan(&totalCount); err != nil {
		return nil, pageable.Page{}, err
	}

	// init page
	page := pageable.CreatePage(pq, totalCount)

	// handle empty result
	if totalCount == 0 {
		return make([]OrderLines, 0), page, nil
	}

	// select entities
	query := `		
		select
			order_id,
			line_no,
			book_id,
			quantity,
			unit_price,
			discount_amount,
			note
		from bookstore_sales.order_lines
		order by order_id, line_no
		offset $1 limit $2
		`

	rows, err := r.db.Pool.Query(ctx, query, pq.GetOffset(), pq.GetLimit())
	if err != nil {
		return nil, pageable.Page{}, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	scannedEntities := make([]OrderLines, 0)
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
func scanFullRow(row pgx.Row) (*OrderLines, error) {
	var scannedEntity OrderLines
	err := row.Scan(
		&scannedEntity.OrderID,
		&scannedEntity.LineNo,
		&scannedEntity.BookID,
		&scannedEntity.Quantity,
		&scannedEntity.UnitPrice,
		&scannedEntity.DiscountAmount,
		&scannedEntity.Note,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperrors.ErrNotFound
		}
		return nil, err
	}
	return &scannedEntity, nil
}
