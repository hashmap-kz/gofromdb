package impl

import (
	"context"
	"fmt"

	dbModel "go-project-template-v5/internal/api/category/entity/postgres"

	"go-project-template-v5/internal/api/category/repository"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/pkg/storage/postgres"
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

	var scannedEntity dbModel.Category
	err := r.db.Pool.QueryRow(ctx, query,
		inputEntity.Name,
		inputEntity.ParentID,
	).Scan(
		&scannedEntity.RecordID,
		&scannedEntity.Name,
		&scannedEntity.ParentID,
		&scannedEntity.CreatedAt,
		&scannedEntity.UpdatedAt,
		&scannedEntity.Guid,
	)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return &scannedEntity, nil
}

func (r *categoryRepository) Update(ctx context.Context, inputEntity *dbModel.Category) (*dbModel.Category, error) {
	tag := "categoryRepository.Update"

	query := `		
		update public.category
		set 
			name = $2,
			parent_id = $3
		where record_id = $1
		returning 
			record_id,
			name,
			parent_id,
			created_at,
			updated_at,
			guid
		`

	var scannedEntity dbModel.Category
	err := r.db.Pool.QueryRow(ctx, query,
		inputEntity.RecordID,
		inputEntity.Name,
		inputEntity.ParentID,
	).Scan(
		&scannedEntity.RecordID,
		&scannedEntity.Name,
		&scannedEntity.ParentID,
		&scannedEntity.CreatedAt,
		&scannedEntity.UpdatedAt,
		&scannedEntity.Guid,
	)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return &scannedEntity, nil
}

func (r *categoryRepository) Delete(ctx context.Context, id int) error {
	tag := "categoryRepository.Delete"

	query := `		
		delete from only public.category
		where record_id = $1
		`

	cmdTag, err := r.db.Pool.Exec(ctx, query, id)
	if err != nil || cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("%s. no rows deleted for id: %v, %w", tag, id, err)
	}
	return nil
}

func (r *categoryRepository) GetByID(ctx context.Context, id int) (*dbModel.Category, error) {
	tag := "categoryRepository.GetByID"

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

	var scannedEntity dbModel.Category
	err := r.db.Pool.QueryRow(ctx, query, id).Scan(
		&scannedEntity.RecordID,
		&scannedEntity.Name,
		&scannedEntity.ParentID,
		&scannedEntity.CreatedAt,
		&scannedEntity.UpdatedAt,
		&scannedEntity.Guid,
	)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return &scannedEntity, nil
}

func (r *categoryRepository) GetAll(ctx context.Context) ([]dbModel.Category, error) {
	tag := "categoryRepository.GetAll"

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
		var scannedEntity dbModel.Category
		err = rows.Scan(
			&scannedEntity.RecordID,
			&scannedEntity.Name,
			&scannedEntity.ParentID,
			&scannedEntity.CreatedAt,
			&scannedEntity.UpdatedAt,
			&scannedEntity.Guid,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", tag, err)
		}
		scannedEntities = append(scannedEntities, scannedEntity)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return scannedEntities, nil
}

func (r *categoryRepository) GetAllPaginated(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.Category, pageable.Page, error) {
	tag := "categoryRepository.GetAllPaginated"

	// retrieve total count
	queryCnt := `select count(record_id) from public.category`
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
		var scannedEntity dbModel.Category
		err = rows.Scan(
			&scannedEntity.RecordID,
			&scannedEntity.Name,
			&scannedEntity.ParentID,
			&scannedEntity.CreatedAt,
			&scannedEntity.UpdatedAt,
			&scannedEntity.Guid,
		)
		if err != nil {
			return nil, pageable.Page{}, fmt.Errorf("%s: %w", tag, err)
		}
		scannedEntities = append(scannedEntities, scannedEntity)
	}

	if rows.Err() != nil {
		return nil, page, rows.Err()
	}
	return scannedEntities, page, nil
}
