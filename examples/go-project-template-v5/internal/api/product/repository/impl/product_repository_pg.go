package impl

import (
	"context"
	"fmt"

	dbModel "go-project-template-v5/internal/api/product/entity/postgres"

	"go-project-template-v5/internal/api/product/repository"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/pkg/storage/postgres"
)

type productRepository struct {
	db *postgres.Postgres
}

var _ repository.ProductRepository = &productRepository{}

func NewProductRepository(_ context.Context, db *postgres.Postgres) repository.ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) Save(ctx context.Context, inputEntity *dbModel.Product) (*dbModel.Product, error) {
	tag := "productRepository.Save"

	query := `		
		insert into public.product (
			category_id,
			name,
			description
		)
		values ($1, $2, $3)
		returning
			record_id,
			category_id,
			name,
			description,
			created_at,
			updated_at,
			guid
		`

	var scannedEntity dbModel.Product
	err := r.db.Pool.QueryRow(ctx, query,
		inputEntity.CategoryID,
		inputEntity.Name,
		inputEntity.Description,
	).Scan(
		&scannedEntity.RecordID,
		&scannedEntity.CategoryID,
		&scannedEntity.Name,
		&scannedEntity.Description,
		&scannedEntity.CreatedAt,
		&scannedEntity.UpdatedAt,
		&scannedEntity.Guid,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return &scannedEntity, nil
}

func (r *productRepository) Update(ctx context.Context, inputEntity *dbModel.Product) (*dbModel.Product, error) {
	tag := "productRepository.Update"

	query := `		
		update public.product
		set 
			category_id = $2,
			name = $3,
			description = $4
		where record_id = $1
		returning 
			record_id,
			category_id,
			name,
			description,
			created_at,
			updated_at,
			guid
		`

	var scannedEntity dbModel.Product
	err := r.db.Pool.QueryRow(ctx, query,
		inputEntity.RecordID,
		inputEntity.CategoryID,
		inputEntity.Name,
		inputEntity.Description,
	).Scan(
		&scannedEntity.RecordID,
		&scannedEntity.CategoryID,
		&scannedEntity.Name,
		&scannedEntity.Description,
		&scannedEntity.CreatedAt,
		&scannedEntity.UpdatedAt,
		&scannedEntity.Guid,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return &scannedEntity, nil
}

func (r *productRepository) Delete(ctx context.Context, id int) error {
	tag := "productRepository.Delete"

	query := `		
		delete from only public.product
		where record_id = $1
		`

	cmdTag, err := r.db.Pool.Exec(ctx, query, id)
	if err != nil || cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("%s. no rows deleted for id: %v, %w", tag, id, err)
	}
	return nil
}

func (r *productRepository) GetByID(ctx context.Context, id int) (*dbModel.Product, error) {
	tag := "productRepository.GetByID"

	query := `		
		select
			record_id,
			category_id,
			name,
			description,
			created_at,
			updated_at,
			guid
		from public.product
		where record_id = $1
		order by record_id
		`

	var scannedEntity dbModel.Product
	err := r.db.Pool.QueryRow(ctx, query, id).Scan(
		&scannedEntity.RecordID,
		&scannedEntity.CategoryID,
		&scannedEntity.Name,
		&scannedEntity.Description,
		&scannedEntity.CreatedAt,
		&scannedEntity.UpdatedAt,
		&scannedEntity.Guid,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return &scannedEntity, nil
}

func (r *productRepository) GetAll(ctx context.Context) ([]dbModel.Product, error) {
	tag := "productRepository.GetAll"

	query := `		
		select
			record_id,
			category_id,
			name,
			description,
			created_at,
			updated_at,
			guid
		from public.product
		order by record_id
		`

	rows, err := r.db.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	var scannedEntities []dbModel.Product
	for rows.Next() {
		var scannedEntity dbModel.Product
		err = rows.Scan(
			&scannedEntity.RecordID,
			&scannedEntity.CategoryID,
			&scannedEntity.Name,
			&scannedEntity.Description,
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

func (r *productRepository) GetAllPaginated(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.Product, pageable.Page, error) {
	tag := "productRepository.GetAllPaginated"

	// retrieve total count
	queryCnt := `		select count(record_id) from public.product`
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
			category_id,
			name,
			description,
			created_at,
			updated_at,
			guid
		from public.product
		order by record_id
		offset $1 limit $2
		`

	rows, err := r.db.Pool.Query(ctx, query, pq.GetOffset(), pq.GetLimit())
	if err != nil {
		return nil, pageable.Page{}, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	var scannedEntities []dbModel.Product
	for rows.Next() {
		var scannedEntity dbModel.Product
		err = rows.Scan(
			&scannedEntity.RecordID,
			&scannedEntity.CategoryID,
			&scannedEntity.Name,
			&scannedEntity.Description,
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
