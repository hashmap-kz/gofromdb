package impl

import (
	"context"
	"fmt"

	dbModel "go-project-template-v5/internal/api/steps/entity/postgres"

	"go-project-template-v5/internal/api/steps/repository"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/pkg/storage/postgres"

	"github.com/jackc/pgx/v5"
)

type stepsRepository struct {
	db *postgres.Postgres
}

var _ repository.StepsRepository = &stepsRepository{}

func NewStepsRepository(_ context.Context, db *postgres.Postgres) repository.StepsRepository {
	return &stepsRepository{
		db: db,
	}
}

func (r *stepsRepository) Save(ctx context.Context, inputEntity *dbModel.Steps) (*dbModel.Steps, error) {
	tag := "stepsRepository.Save"

	query := `		
		insert into public.steps (
			step_name
		)
		values ($1)
		returning
			record_id,
			step_name,
			created_at,
			updated_at,
			guid
		`

	row := r.db.Pool.QueryRow(ctx, query,
		inputEntity.StepName,
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *stepsRepository) UpdateByID(ctx context.Context, inputEntity *dbModel.Steps, pkRecordID int) (*dbModel.Steps, error) {
	tag := "stepsRepository.UpdateByID"

	query := `		
		update public.steps
		set 
			step_name = coalesce(nullif($2, ''), step_name)
		where record_id = $1
		returning 
			record_id,
			step_name,
			created_at,
			updated_at,
			guid
		`

	row := r.db.Pool.QueryRow(ctx, query,
		pkRecordID,
		inputEntity.StepName,
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *stepsRepository) DeleteByID(ctx context.Context, pkRecordID int) error {
	tag := "stepsRepository.DeleteByID"

	query := `		
		delete from only public.steps
		where record_id = $1
		`

	cmdTag, err := r.db.Pool.Exec(ctx, query, pkRecordID)
	if err != nil || cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("%s. no rows were deleted: %w", tag, err)
	}
	return nil
}

func (r *stepsRepository) FindByID(ctx context.Context, pkRecordID int) (*dbModel.Steps, error) {
	tag := "stepsRepository.FindByID"

	query := `		
		select
			record_id,
			step_name,
			created_at,
			updated_at,
			guid
		from public.steps
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

func (r *stepsRepository) FindAll(ctx context.Context) ([]dbModel.Steps, error) {
	tag := "stepsRepository.FindAll"

	query := `		
		select
			record_id,
			step_name,
			created_at,
			updated_at,
			guid
		from public.steps
		order by record_id
		`

	rows, err := r.db.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	var scannedEntities []dbModel.Steps
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

func (r *stepsRepository) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.Steps, pageable.Page, error) {
	tag := "stepsRepository.FindAllPageable"

	// retrieve total count
	queryCnt := `select count(*) from public.steps`
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
			step_name,
			created_at,
			updated_at,
			guid
		from public.steps
		order by record_id
		offset $1 limit $2
		`

	rows, err := r.db.Pool.Query(ctx, query, pq.GetOffset(), pq.GetLimit())
	if err != nil {
		return nil, pageable.Page{}, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	var scannedEntities []dbModel.Steps
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
func scanFullRow(row pgx.Row) (*dbModel.Steps, error) {
	var scannedEntity dbModel.Steps
	err := row.Scan(
		&scannedEntity.RecordID,
		&scannedEntity.StepName,
		&scannedEntity.CreatedAt,
		&scannedEntity.UpdatedAt,
		&scannedEntity.Guid,
	)
	if err != nil {
		return nil, err
	}
	return &scannedEntity, nil
}
