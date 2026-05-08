package impl

import (
	"context"
	"fmt"

	dbModel "go-project-template-v5/internal/api/no_pk/entity/postgres"

	"go-project-template-v5/internal/api/no_pk/repository"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/pkg/storage/postgres"

	"github.com/jackc/pgx/v5"
)

type noPkRepository struct {
	db *postgres.Postgres
}

var _ repository.NoPkRepository = &noPkRepository{}

func NewNoPkRepository(_ context.Context, db *postgres.Postgres) repository.NoPkRepository {
	return &noPkRepository{
		db: db,
	}
}

func (r *noPkRepository) Save(ctx context.Context, inputEntity *dbModel.NoPk) (*dbModel.NoPk, error) {
	tag := "noPkRepository.Save"

	query := `		
		insert into public.no_pk (
			event_time,
			message
		)
		values ($1, $2)
		returning
			event_time,
			message
		`

	row := r.db.Pool.QueryRow(ctx, query,
		inputEntity.EventTime,
		inputEntity.Message,
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *noPkRepository) FindAll(ctx context.Context) ([]dbModel.NoPk, error) {
	tag := "noPkRepository.FindAll"

	query := `		
		select
			event_time,
			message
		from public.no_pk
		
		`

	rows, err := r.db.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	var scannedEntities []dbModel.NoPk
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

func (r *noPkRepository) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.NoPk, pageable.Page, error) {
	tag := "noPkRepository.FindAllPageable"

	// retrieve total count
	queryCnt := `select count(*) from public.no_pk`
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
			event_time,
			message
		from public.no_pk
		
		offset $1 limit $2
		`

	rows, err := r.db.Pool.Query(ctx, query, pq.GetOffset(), pq.GetLimit())
	if err != nil {
		return nil, pageable.Page{}, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	var scannedEntities []dbModel.NoPk
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
func scanFullRow(row pgx.Row) (*dbModel.NoPk, error) {
	var scannedEntity dbModel.NoPk
	err := row.Scan(
		&scannedEntity.EventTime,
		&scannedEntity.Message,
	)
	if err != nil {
		return nil, err
	}
	return &scannedEntity, nil
}
