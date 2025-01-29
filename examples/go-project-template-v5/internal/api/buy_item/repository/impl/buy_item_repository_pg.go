package impl

import (
	"context"
	"fmt"

	dbModel "go-project-template-v5/internal/api/buy_item/entity/postgres"

	"go-project-template-v5/internal/api/buy_item/repository"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/pkg/storage/postgres"
)

type buyItemRepository struct {
	db *postgres.Postgres
}

var _ repository.BuyItemRepository = &buyItemRepository{}

func NewBuyItemRepository(_ context.Context, db *postgres.Postgres) repository.BuyItemRepository {
	return &buyItemRepository{
		db: db,
	}
}

func (r *buyItemRepository) Save(ctx context.Context, inputEntity *dbModel.BuyItem) (*dbModel.BuyItem, error) {
	tag := "buyItemRepository.Save"

	query := `		
		insert into public.buy_item (
			buy_id,
			product_id,
			quantity,
			price
		)
		values ($1, $2, $3, $4)
		returning
			record_id,
			buy_id,
			product_id,
			quantity,
			price,
			created_at,
			updated_at,
			guid
		`

	var scannedEntity dbModel.BuyItem
	err := r.db.Pool.QueryRow(ctx, query,
		inputEntity.BuyID,
		inputEntity.ProductID,
		inputEntity.Quantity,
		inputEntity.Price,
	).Scan(
		&scannedEntity.RecordID,
		&scannedEntity.BuyID,
		&scannedEntity.ProductID,
		&scannedEntity.Quantity,
		&scannedEntity.Price,
		&scannedEntity.CreatedAt,
		&scannedEntity.UpdatedAt,
		&scannedEntity.Guid,
	)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return &scannedEntity, nil
}

func (r *buyItemRepository) Update(ctx context.Context, inputEntity *dbModel.BuyItem) (*dbModel.BuyItem, error) {
	tag := "buyItemRepository.Update"

	query := `		
		update public.buy_item
		set 
			buy_id = $2,
			product_id = $3,
			quantity = $4,
			price = $5
		where record_id = $1
		returning 
			record_id,
			buy_id,
			product_id,
			quantity,
			price,
			created_at,
			updated_at,
			guid
		`

	var scannedEntity dbModel.BuyItem
	err := r.db.Pool.QueryRow(ctx, query,
		inputEntity.RecordID,
		inputEntity.BuyID,
		inputEntity.ProductID,
		inputEntity.Quantity,
		inputEntity.Price,
	).Scan(
		&scannedEntity.RecordID,
		&scannedEntity.BuyID,
		&scannedEntity.ProductID,
		&scannedEntity.Quantity,
		&scannedEntity.Price,
		&scannedEntity.CreatedAt,
		&scannedEntity.UpdatedAt,
		&scannedEntity.Guid,
	)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return &scannedEntity, nil
}

func (r *buyItemRepository) Delete(ctx context.Context, id int) error {
	tag := "buyItemRepository.Delete"

	query := `		
		delete from only public.buy_item
		where record_id = $1
		`

	cmdTag, err := r.db.Pool.Exec(ctx, query, id)
	if err != nil || cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("%s. no rows deleted for id: %v, %w", tag, id, err)
	}
	return nil
}

func (r *buyItemRepository) GetByID(ctx context.Context, id int) (*dbModel.BuyItem, error) {
	tag := "buyItemRepository.GetByID"

	query := `		
		select
			record_id,
			buy_id,
			product_id,
			quantity,
			price,
			created_at,
			updated_at,
			guid
		from public.buy_item
		where record_id = $1
		order by record_id
		`

	var scannedEntity dbModel.BuyItem
	err := r.db.Pool.QueryRow(ctx, query, id).Scan(
		&scannedEntity.RecordID,
		&scannedEntity.BuyID,
		&scannedEntity.ProductID,
		&scannedEntity.Quantity,
		&scannedEntity.Price,
		&scannedEntity.CreatedAt,
		&scannedEntity.UpdatedAt,
		&scannedEntity.Guid,
	)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return &scannedEntity, nil
}

func (r *buyItemRepository) GetAll(ctx context.Context) ([]dbModel.BuyItem, error) {
	tag := "buyItemRepository.GetAll"

	query := `		
		select
			record_id,
			buy_id,
			product_id,
			quantity,
			price,
			created_at,
			updated_at,
			guid
		from public.buy_item
		order by record_id
		`

	rows, err := r.db.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	var scannedEntities []dbModel.BuyItem
	for rows.Next() {
		var scannedEntity dbModel.BuyItem
		err = rows.Scan(
			&scannedEntity.RecordID,
			&scannedEntity.BuyID,
			&scannedEntity.ProductID,
			&scannedEntity.Quantity,
			&scannedEntity.Price,
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

func (r *buyItemRepository) GetAllPaginated(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.BuyItem, pageable.Page, error) {
	tag := "buyItemRepository.GetAllPaginated"

	// retrieve total count
	queryCnt := `select count(record_id) from public.buy_item`
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
			buy_id,
			product_id,
			quantity,
			price,
			created_at,
			updated_at,
			guid
		from public.buy_item
		order by record_id
		offset $1 limit $2
		`

	rows, err := r.db.Pool.Query(ctx, query, pq.GetOffset(), pq.GetLimit())
	if err != nil {
		return nil, pageable.Page{}, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	var scannedEntities []dbModel.BuyItem
	for rows.Next() {
		var scannedEntity dbModel.BuyItem
		err = rows.Scan(
			&scannedEntity.RecordID,
			&scannedEntity.BuyID,
			&scannedEntity.ProductID,
			&scannedEntity.Quantity,
			&scannedEntity.Price,
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
