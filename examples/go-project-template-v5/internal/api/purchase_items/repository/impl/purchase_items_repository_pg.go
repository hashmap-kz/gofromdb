package impl

import (
	"context"
	"fmt"

	dbModel "go-project-template-v5/internal/api/purchase_items/entity/postgres"

	"go-project-template-v5/internal/api/purchase_items/repository"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/pkg/storage/postgres"

	"github.com/jackc/pgx/v5"
)

type purchaseItemsRepository struct {
	db *postgres.Postgres
}

var _ repository.PurchaseItemsRepository = &purchaseItemsRepository{}

func NewPurchaseItemsRepository(_ context.Context, db *postgres.Postgres) repository.PurchaseItemsRepository {
	return &purchaseItemsRepository{
		db: db,
	}
}

func (r *purchaseItemsRepository) Save(ctx context.Context, inputEntity *dbModel.PurchaseItems) (*dbModel.PurchaseItems, error) {
	tag := "purchaseItemsRepository.Save"

	query := `		
		insert into public.purchase_items (
			purchase_id,
			product_id,
			quantity,
			price
		)
		values ($1, $2, $3, $4)
		returning
			record_id,
			purchase_id,
			product_id,
			quantity,
			price,
			created_at,
			updated_at,
			guid
		`

	row := r.db.Pool.QueryRow(ctx, query,
		inputEntity.PurchaseID,
		inputEntity.ProductID,
		inputEntity.Quantity,
		inputEntity.Price,
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *purchaseItemsRepository) UpdateByID(ctx context.Context, inputEntity *dbModel.PurchaseItems, pkRecordID int) (*dbModel.PurchaseItems, error) {
	tag := "purchaseItemsRepository.UpdateByID"

	query := `		
		update public.purchase_items
		set 
			purchase_id = coalesce(nullif($2, 0::int4), purchase_id),
			product_id  = coalesce(nullif($3, 0::int4), product_id),
			quantity    = coalesce(nullif($4, 0::int4), quantity),
			price       = coalesce(nullif($5, 0::numeric), price)
		where record_id = $1
		returning 
			record_id,
			purchase_id,
			product_id,
			quantity,
			price,
			created_at,
			updated_at,
			guid
		`

	row := r.db.Pool.QueryRow(ctx, query,
		pkRecordID,
		inputEntity.PurchaseID,
		inputEntity.ProductID,
		inputEntity.Quantity,
		inputEntity.Price,
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *purchaseItemsRepository) DeleteByID(ctx context.Context, pkRecordID int) error {
	tag := "purchaseItemsRepository.DeleteByID"

	query := `		
		delete from only public.purchase_items
		where record_id = $1
		`

	cmdTag, err := r.db.Pool.Exec(ctx, query, pkRecordID)
	if err != nil || cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("%s. no rows were deleted: %w", tag, err)
	}
	return nil
}

func (r *purchaseItemsRepository) FindByID(ctx context.Context, pkRecordID int) (*dbModel.PurchaseItems, error) {
	tag := "purchaseItemsRepository.FindByID"

	query := `		
		select
			record_id,
			purchase_id,
			product_id,
			quantity,
			price,
			created_at,
			updated_at,
			guid
		from public.purchase_items
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

func (r *purchaseItemsRepository) FindAll(ctx context.Context) ([]dbModel.PurchaseItems, error) {
	tag := "purchaseItemsRepository.FindAll"

	query := `		
		select
			record_id,
			purchase_id,
			product_id,
			quantity,
			price,
			created_at,
			updated_at,
			guid
		from public.purchase_items
		order by record_id
		`

	rows, err := r.db.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	var scannedEntities []dbModel.PurchaseItems
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

func (r *purchaseItemsRepository) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.PurchaseItems, pageable.Page, error) {
	tag := "purchaseItemsRepository.FindAllPageable"

	// retrieve total count
	queryCnt := `select count(*) from public.purchase_items`
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
			purchase_id,
			product_id,
			quantity,
			price,
			created_at,
			updated_at,
			guid
		from public.purchase_items
		order by record_id
		offset $1 limit $2
		`

	rows, err := r.db.Pool.Query(ctx, query, pq.GetOffset(), pq.GetLimit())
	if err != nil {
		return nil, pageable.Page{}, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	var scannedEntities []dbModel.PurchaseItems
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
func scanFullRow(row pgx.Row) (*dbModel.PurchaseItems, error) {
	var scannedEntity dbModel.PurchaseItems
	err := row.Scan(
		&scannedEntity.RecordID,
		&scannedEntity.PurchaseID,
		&scannedEntity.ProductID,
		&scannedEntity.Quantity,
		&scannedEntity.Price,
		&scannedEntity.CreatedAt,
		&scannedEntity.UpdatedAt,
		&scannedEntity.Guid,
	)
	if err != nil {
		return nil, err
	}
	return &scannedEntity, nil
}
