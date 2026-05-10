package import_batches

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
	Save(ctx context.Context, inputEntity *ImportBatches) (*ImportBatches, error)
	UpdateByID(ctx context.Context, update *UpdateDto, pkSourceName string, pkBatchNo int) (*ImportBatches, error)
	DeleteByID(ctx context.Context, pkSourceName string, pkBatchNo int) error
	FindByID(ctx context.Context, pkSourceName string, pkBatchNo int) (*ImportBatches, error)
	FindAll(ctx context.Context) ([]ImportBatches, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]ImportBatches, pageable.Page, error)
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

func (r *repo) Save(ctx context.Context, inputEntity *ImportBatches) (*ImportBatches, error) {
	tag := "repository.Save"

	query := `		
		insert into bookstore_import.import_batches (
			source_name,
			batch_no,
			started_at,
			finished_at,
			file_name,
			row_count,
			metadata
		)
		values ($1, $2, $3, $4, $5, $6, $7)
		returning
			source_name,
			batch_no,
			started_at,
			finished_at,
			file_name,
			row_count,
			metadata
		`

	row := r.db.Pool.QueryRow(
		ctx, query,
		inputEntity.SourceName,
		inputEntity.BatchNo,
		inputEntity.StartedAt,
		inputEntity.FinishedAt,
		inputEntity.FileName,
		inputEntity.RowCount,
		inputEntity.Metadata,
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *repo) UpdateByID(ctx context.Context, update *UpdateDto, pkSourceName string, pkBatchNo int) (*ImportBatches, error) {
	tag := "repository.UpdateByID"

	if update == nil {
		return nil, fmt.Errorf("%s: update is nil", tag)
	}

	query := `		
		update bookstore_import.import_batches
		set
			started_at  = coalesce($3, started_at),
			finished_at = coalesce($4, finished_at),
			file_name   = coalesce($5, file_name),
			row_count   = coalesce($6, row_count),
			metadata    = coalesce($7, metadata)
		where source_name = $1 and batch_no = $2
		returning
			source_name,
			batch_no,
			started_at,
			finished_at,
			file_name,
			row_count,
			metadata
		`

	row := r.db.Pool.QueryRow(
		ctx, query,
		pkSourceName, pkBatchNo,
		update.StartedAt,
		update.FinishedAt,
		update.FileName,
		update.RowCount,
		update.Metadata,
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *repo) DeleteByID(ctx context.Context, pkSourceName string, pkBatchNo int) error {
	tag := "repository.DeleteByID"

	query := `		
		delete from only bookstore_import.import_batches
		where source_name = $1 and batch_no = $2
		`

	cmdTag, err := r.db.Pool.Exec(ctx, query, pkSourceName, pkBatchNo)
	if err != nil {
		return fmt.Errorf("%s: %w", tag, err)
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("%s: %w", tag, apperrors.ErrNotFound)
	}
	return nil
}

func (r *repo) FindByID(ctx context.Context, pkSourceName string, pkBatchNo int) (*ImportBatches, error) {
	tag := "repository.FindByID"

	query := `		
		select
			source_name,
			batch_no,
			started_at,
			finished_at,
			file_name,
			row_count,
			metadata
		from bookstore_import.import_batches
		where source_name = $1 and batch_no = $2
		order by source_name, batch_no
		`

	row := r.db.Pool.QueryRow(ctx, query, pkSourceName, pkBatchNo)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *repo) FindAll(ctx context.Context) ([]ImportBatches, error) {
	tag := "repository.FindAll"

	query := `		
		select
			source_name,
			batch_no,
			started_at,
			finished_at,
			file_name,
			row_count,
			metadata
		from bookstore_import.import_batches
		order by source_name, batch_no
		`

	rows, err := r.db.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	scannedEntities := make([]ImportBatches, 0)
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

func (r *repo) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]ImportBatches, pageable.Page, error) {
	tag := "repository.FindAllPageable"

	// retrieve total count
	queryCnt := `select count(*) from bookstore_import.import_batches`
	var totalCount int
	if err := r.db.Pool.QueryRow(ctx, queryCnt).Scan(&totalCount); err != nil {
		return nil, pageable.Page{}, err
	}

	// init page
	page := pageable.CreatePage(pq, totalCount)

	// handle empty result
	if totalCount == 0 {
		return make([]ImportBatches, 0), page, nil
	}

	// select entities
	query := `		
		select
			source_name,
			batch_no,
			started_at,
			finished_at,
			file_name,
			row_count,
			metadata
		from bookstore_import.import_batches
		order by source_name, batch_no
		offset $1 limit $2
		`

	rows, err := r.db.Pool.Query(ctx, query, pq.GetOffset(), pq.GetLimit())
	if err != nil {
		return nil, pageable.Page{}, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	scannedEntities := make([]ImportBatches, 0)
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
func scanFullRow(row pgx.Row) (*ImportBatches, error) {
	var scannedEntity ImportBatches
	err := row.Scan(
		&scannedEntity.SourceName,
		&scannedEntity.BatchNo,
		&scannedEntity.StartedAt,
		&scannedEntity.FinishedAt,
		&scannedEntity.FileName,
		&scannedEntity.RowCount,
		&scannedEntity.Metadata,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperrors.ErrNotFound
		}
		return nil, err
	}
	return &scannedEntity, nil
}
