package import_errors

import (
	"context"
	"fmt"
	"go-project-template-v7/pkg/pageable"
	"go-project-template-v7/pkg/storage/postgres"

	"github.com/jackc/pgx/v5"
)

type Repository interface {
	Save(ctx context.Context, inputEntity *ImportErrors) (*ImportErrors, error)
	FindAll(ctx context.Context) ([]ImportErrors, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]ImportErrors, pageable.Page, error)
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

func (r *repo) Save(ctx context.Context, inputEntity *ImportErrors) (*ImportErrors, error) {
	tag := "repository.Save"

	query := `		
		insert into bookstore_import.import_errors (
			source_name,
			batch_no,
			row_no,
			column_name,
			message,
			raw_payload
		)
		values ($1, $2, $3, $4, $5, $6)
		returning
			source_name,
			batch_no,
			row_no,
			column_name,
			message,
			raw_payload,
			created_at
		`

	row := r.db.Pool.QueryRow(
		ctx, query,
		inputEntity.SourceName,
		inputEntity.BatchNo,
		inputEntity.RowNo,
		inputEntity.ColumnName,
		inputEntity.Message,
		inputEntity.RawPayload,
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *repo) FindAll(ctx context.Context) ([]ImportErrors, error) {
	tag := "repository.FindAll"

	query := `		
		select
			source_name,
			batch_no,
			row_no,
			column_name,
			message,
			raw_payload,
			created_at
		from bookstore_import.import_errors
		
		`

	rows, err := r.db.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	var scannedEntities []ImportErrors
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

func (r *repo) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]ImportErrors, pageable.Page, error) {
	tag := "repository.FindAllPageable"

	// retrieve total count
	queryCnt := `select count(*) from bookstore_import.import_errors`
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
			source_name,
			batch_no,
			row_no,
			column_name,
			message,
			raw_payload,
			created_at
		from bookstore_import.import_errors
		
		offset $1 limit $2
		`

	rows, err := r.db.Pool.Query(ctx, query, pq.GetOffset(), pq.GetLimit())
	if err != nil {
		return nil, pageable.Page{}, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	var scannedEntities []ImportErrors
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
func scanFullRow(row pgx.Row) (*ImportErrors, error) {
	var scannedEntity ImportErrors
	err := row.Scan(
		&scannedEntity.SourceName,
		&scannedEntity.BatchNo,
		&scannedEntity.RowNo,
		&scannedEntity.ColumnName,
		&scannedEntity.Message,
		&scannedEntity.RawPayload,
		&scannedEntity.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &scannedEntity, nil
}
