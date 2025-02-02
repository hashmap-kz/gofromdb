package impl

import (
	"context"
	"fmt"

	dbModel "go-project-template-v5/internal/api/category/entity/postgres"

	"go-project-template-v5/internal/api/category/repository"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/pkg/storage/postgres"

	"github.com/jackc/pgx/v5"
)

type categoryRepository struct {
	db *postgres.Postgres
}

var _ repository.CategoryRepository = &categoryRepository{}

func NewCategoryRepository(_ context.Context, db *postgres.Postgres) repository.CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}

func (r *categoryRepository) Save(ctx context.Context, inputEntity *dbModel.Category) (*dbModel.Category, error) {
	tag := "categoryRepository.Save"

	query := `		
		insert into public.category (
			name,
			parent_id
		)
		values ($1, $2)
		returning
			record_id,
			name,
			parent_id,
			created_at,
			updated_at,
			guid
		`

	row := r.db.Pool.QueryRow(ctx, query,
		inputEntity.Name,
		inputEntity.ParentID,
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *categoryRepository) UpdateByID(ctx context.Context, pkRecordID int, inputEntity *dbModel.Category) (*dbModel.Category, error) {
	tag := "categoryRepository.UpdateByID"

	query := `		
		update public.category
		set 
			name = coalesce(nullif($2, ''), name),
			parent_id = coalesce(nullif($3, 0::int4), parent_id)
		where record_id = $1
		returning 
			record_id,
			name,
			parent_id,
			created_at,
			updated_at,
			guid
		`

	row := r.db.Pool.QueryRow(ctx, query,
		pkRecordID,
		inputEntity.Name,
		inputEntity.ParentID,
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *categoryRepository) DeleteByID(ctx context.Context, pkRecordID int) error {
	tag := "categoryRepository.DeleteByID"

	query := `		
		delete from only public.category
		where record_id = $1
		`

	cmdTag, err := r.db.Pool.Exec(ctx, query, pkRecordID)
	if err != nil || cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("%s. no rows were deleted: %w", tag, err)
	}
	return nil
}

func (r *categoryRepository) FindByID(ctx context.Context, pkRecordID int) (*dbModel.Category, error) {
	tag := "categoryRepository.FindByID"

	query := `		
		select
			record_id,
			name,
			parent_id,
			created_at,
			updated_at,
			guid
		from public.category
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

func (r *categoryRepository) FindAll(ctx context.Context) ([]dbModel.Category, error) {
	tag := "categoryRepository.FindAll"

	query := `		
		select
			record_id,
			name,
			parent_id,
			created_at,
			updated_at,
			guid
		from public.category
		order by record_id
		`

	rows, err := r.db.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	var scannedEntities []dbModel.Category
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

func (r *categoryRepository) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.Category, pageable.Page, error) {
	tag := "categoryRepository.FindAllPageable"

	// retrieve total count
	queryCnt := `select count(*) from public.category`
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
			name,
			parent_id,
			created_at,
			updated_at,
			guid
		from public.category
		order by record_id
		offset $1 limit $2
		`

	rows, err := r.db.Pool.Query(ctx, query, pq.GetOffset(), pq.GetLimit())
	if err != nil {
		return nil, pageable.Page{}, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	var scannedEntities []dbModel.Category
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
func scanFullRow(row pgx.Row) (*dbModel.Category, error) {
	var scannedEntity dbModel.Category
	err := row.Scan(
		&scannedEntity.RecordID,
		&scannedEntity.Name,
		&scannedEntity.ParentID,
		&scannedEntity.CreatedAt,
		&scannedEntity.UpdatedAt,
		&scannedEntity.Guid,
	)
	if err != nil {
		return nil, err
	}
	return &scannedEntity, nil
}
