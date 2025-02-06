package impl

import (
	"context"
	"fmt"

	dbModel "go-project-template-v5/internal/api/product_prices/entity/postgres"

	"go-project-template-v5/internal/api/product_prices/repository"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/pkg/storage/postgres"

	"github.com/jackc/pgx/v5"
)

type productPricesRepository struct {
	db *postgres.Postgres
}

var _ repository.ProductPricesRepository = &productPricesRepository{}

func NewProductPricesRepository(_ context.Context, db *postgres.Postgres) repository.ProductPricesRepository {
	return &productPricesRepository{
		db: db,
	}
}

func (r *productPricesRepository) Save(ctx context.Context, inputEntity *dbModel.ProductPrices) (*dbModel.ProductPrices, error) {
	tag := "productPricesRepository.Save"

	query := `		
		insert into public.product_prices (
			product_price_period,
			product_id,
			product_price
		)
		values ($1, $2, $3)
		returning
			record_id,
			product_price_period,
			product_id,
			product_price,
			created_at,
			updated_at,
			guid
		`

	row := r.db.Pool.QueryRow(ctx, query,
		inputEntity.ProductPricePeriod,
		inputEntity.ProductID,
		inputEntity.ProductPrice,
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *productPricesRepository) UpdateByID(ctx context.Context, inputEntity *dbModel.ProductPrices, pkRecordID int) (*dbModel.ProductPrices, error) {
	tag := "productPricesRepository.UpdateByID"

	query := `		
		update public.product_prices
		set 
			product_price_period = coalesce(nullif($2, 'empty'::daterange), product_price_period),
			product_id           = coalesce(nullif($3, 0::int4), product_id),
			product_price        = coalesce(nullif($4, 0::numeric), product_price)
		where record_id = $1
		returning 
			record_id,
			product_price_period,
			product_id,
			product_price,
			created_at,
			updated_at,
			guid
		`

	row := r.db.Pool.QueryRow(ctx, query,
		pkRecordID,
		inputEntity.ProductPricePeriod,
		inputEntity.ProductID,
		inputEntity.ProductPrice,
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *productPricesRepository) DeleteByID(ctx context.Context, pkRecordID int) error {
	tag := "productPricesRepository.DeleteByID"

	query := `		
		delete from only public.product_prices
		where record_id = $1
		`

	cmdTag, err := r.db.Pool.Exec(ctx, query, pkRecordID)
	if err != nil || cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("%s. no rows were deleted: %w", tag, err)
	}
	return nil
}

func (r *productPricesRepository) FindByID(ctx context.Context, pkRecordID int) (*dbModel.ProductPrices, error) {
	tag := "productPricesRepository.FindByID"

	query := `		
		select
			record_id,
			product_price_period,
			product_id,
			product_price,
			created_at,
			updated_at,
			guid
		from public.product_prices
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

func (r *productPricesRepository) FindAll(ctx context.Context) ([]dbModel.ProductPrices, error) {
	tag := "productPricesRepository.FindAll"

	query := `		
		select
			record_id,
			product_price_period,
			product_id,
			product_price,
			created_at,
			updated_at,
			guid
		from public.product_prices
		order by record_id
		`

	rows, err := r.db.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	var scannedEntities []dbModel.ProductPrices
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

func (r *productPricesRepository) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.ProductPrices, pageable.Page, error) {
	tag := "productPricesRepository.FindAllPageable"

	// retrieve total count
	queryCnt := `select count(*) from public.product_prices`
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
			product_price_period,
			product_id,
			product_price,
			created_at,
			updated_at,
			guid
		from public.product_prices
		order by record_id
		offset $1 limit $2
		`

	rows, err := r.db.Pool.Query(ctx, query, pq.GetOffset(), pq.GetLimit())
	if err != nil {
		return nil, pageable.Page{}, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	var scannedEntities []dbModel.ProductPrices
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
func scanFullRow(row pgx.Row) (*dbModel.ProductPrices, error) {
	var scannedEntity dbModel.ProductPrices
	err := row.Scan(
		&scannedEntity.RecordID,
		&scannedEntity.ProductPricePeriod,
		&scannedEntity.ProductID,
		&scannedEntity.ProductPrice,
		&scannedEntity.CreatedAt,
		&scannedEntity.UpdatedAt,
		&scannedEntity.Guid,
	)
	if err != nil {
		return nil, err
	}
	return &scannedEntity, nil
}
