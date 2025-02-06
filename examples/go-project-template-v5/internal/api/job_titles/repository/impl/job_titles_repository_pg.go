package impl

import (
	"context"
	"fmt"

	dbModel "go-project-template-v5/internal/api/job_titles/entity/postgres"

	"go-project-template-v5/internal/api/job_titles/repository"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/pkg/storage/postgres"

	"github.com/jackc/pgx/v5"
)

type jobTitlesRepository struct {
	db *postgres.Postgres
}

var _ repository.JobTitlesRepository = &jobTitlesRepository{}

func NewJobTitlesRepository(_ context.Context, db *postgres.Postgres) repository.JobTitlesRepository {
	return &jobTitlesRepository{
		db: db,
	}
}

func (r *jobTitlesRepository) Save(ctx context.Context, inputEntity *dbModel.JobTitles) (*dbModel.JobTitles, error) {
	tag := "jobTitlesRepository.Save"

	query := `		
		insert into public.job_titles (
			title_name
		)
		values ($1)
		returning
			record_id,
			title_name,
			created_at,
			updated_at,
			guid
		`

	row := r.db.Pool.QueryRow(ctx, query,
		inputEntity.TitleName,
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *jobTitlesRepository) UpdateByID(ctx context.Context, inputEntity *dbModel.JobTitles, pkRecordID int) (*dbModel.JobTitles, error) {
	tag := "jobTitlesRepository.UpdateByID"

	query := `		
		update public.job_titles
		set 
			title_name = coalesce(nullif($2, ''), title_name)
		where record_id = $1
		returning 
			record_id,
			title_name,
			created_at,
			updated_at,
			guid
		`

	row := r.db.Pool.QueryRow(ctx, query,
		pkRecordID,
		inputEntity.TitleName,
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *jobTitlesRepository) DeleteByID(ctx context.Context, pkRecordID int) error {
	tag := "jobTitlesRepository.DeleteByID"

	query := `		
		delete from only public.job_titles
		where record_id = $1
		`

	cmdTag, err := r.db.Pool.Exec(ctx, query, pkRecordID)
	if err != nil || cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("%s. no rows were deleted: %w", tag, err)
	}
	return nil
}

func (r *jobTitlesRepository) FindByID(ctx context.Context, pkRecordID int) (*dbModel.JobTitles, error) {
	tag := "jobTitlesRepository.FindByID"

	query := `		
		select
			record_id,
			title_name,
			created_at,
			updated_at,
			guid
		from public.job_titles
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

func (r *jobTitlesRepository) FindAll(ctx context.Context) ([]dbModel.JobTitles, error) {
	tag := "jobTitlesRepository.FindAll"

	query := `		
		select
			record_id,
			title_name,
			created_at,
			updated_at,
			guid
		from public.job_titles
		order by record_id
		`

	rows, err := r.db.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	var scannedEntities []dbModel.JobTitles
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

func (r *jobTitlesRepository) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.JobTitles, pageable.Page, error) {
	tag := "jobTitlesRepository.FindAllPageable"

	// retrieve total count
	queryCnt := `select count(*) from public.job_titles`
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
			title_name,
			created_at,
			updated_at,
			guid
		from public.job_titles
		order by record_id
		offset $1 limit $2
		`

	rows, err := r.db.Pool.Query(ctx, query, pq.GetOffset(), pq.GetLimit())
	if err != nil {
		return nil, pageable.Page{}, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	var scannedEntities []dbModel.JobTitles
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
func scanFullRow(row pgx.Row) (*dbModel.JobTitles, error) {
	var scannedEntity dbModel.JobTitles
	err := row.Scan(
		&scannedEntity.RecordID,
		&scannedEntity.TitleName,
		&scannedEntity.CreatedAt,
		&scannedEntity.UpdatedAt,
		&scannedEntity.Guid,
	)
	if err != nil {
		return nil, err
	}
	return &scannedEntity, nil
}
