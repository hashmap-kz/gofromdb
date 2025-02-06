package impl

import (
	"context"
	"fmt"

	dbModel "go-project-template-v5/internal/api/purchases/entity/postgres"

	"go-project-template-v5/internal/api/purchases/repository"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/pkg/storage/postgres"

	"github.com/jackc/pgx/v5"
)

type purchasesRepository struct {
	db *postgres.Postgres
}

var _ repository.PurchasesRepository = &purchasesRepository{}

func NewPurchasesRepository(_ context.Context, db *postgres.Postgres) repository.PurchasesRepository {
	return &purchasesRepository{
		db: db,
	}
}

func (r *purchasesRepository) Save(ctx context.Context, inputEntity *dbModel.Purchases) (*dbModel.Purchases, error) {
	tag := "purchasesRepository.Save"

	query := `		
		insert into public.purchases (
			customer_id,
			description
		)
		values ($1, $2)
		returning
			record_id,
			customer_id,
			description,
			created_at,
			updated_at,
			guid
		`

	row := r.db.Pool.QueryRow(ctx, query,
		inputEntity.CustomerID,
		inputEntity.Description,
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *purchasesRepository) UpdateByID(ctx context.Context, inputEntity *dbModel.Purchases, pkRecordID int) (*dbModel.Purchases, error) {
	tag := "purchasesRepository.UpdateByID"

	query := `		
		update public.purchases
		set 
			customer_id = coalesce(nullif($2, 0::int4), customer_id),
			description = coalesce(nullif($3, ''), description)
		where record_id = $1
		returning 
			record_id,
			customer_id,
			description,
			created_at,
			updated_at,
			guid
		`

	row := r.db.Pool.QueryRow(ctx, query,
		pkRecordID,
		inputEntity.CustomerID,
		inputEntity.Description,
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *purchasesRepository) DeleteByID(ctx context.Context, pkRecordID int) error {
	tag := "purchasesRepository.DeleteByID"

	query := `		
		delete from only public.purchases
		where record_id = $1
		`

	cmdTag, err := r.db.Pool.Exec(ctx, query, pkRecordID)
	if err != nil || cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("%s. no rows were deleted: %w", tag, err)
	}
	return nil
}

func (r *purchasesRepository) FindByID(ctx context.Context, pkRecordID int) (*dbModel.Purchases, error) {
	tag := "purchasesRepository.FindByID"

	query := `		
		select
			record_id,
			customer_id,
			description,
			created_at,
			updated_at,
			guid
		from public.purchases
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

func (r *purchasesRepository) FindAll(ctx context.Context) ([]dbModel.Purchases, error) {
	tag := "purchasesRepository.FindAll"

	query := `		
		select
			record_id,
			customer_id,
			description,
			created_at,
			updated_at,
			guid
		from public.purchases
		order by record_id
		`

	rows, err := r.db.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	var scannedEntities []dbModel.Purchases
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

func (r *purchasesRepository) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.Purchases, pageable.Page, error) {
	tag := "purchasesRepository.FindAllPageable"

	// retrieve total count
	queryCnt := `select count(*) from public.purchases`
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
			customer_id,
			description,
			created_at,
			updated_at,
			guid
		from public.purchases
		order by record_id
		offset $1 limit $2
		`

	rows, err := r.db.Pool.Query(ctx, query, pq.GetOffset(), pq.GetLimit())
	if err != nil {
		return nil, pageable.Page{}, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	var scannedEntities []dbModel.Purchases
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
func scanFullRow(row pgx.Row) (*dbModel.Purchases, error) {
	var scannedEntity dbModel.Purchases
	err := row.Scan(
		&scannedEntity.RecordID,
		&scannedEntity.CustomerID,
		&scannedEntity.Description,
		&scannedEntity.CreatedAt,
		&scannedEntity.UpdatedAt,
		&scannedEntity.Guid,
	)
	if err != nil {
		return nil, err
	}
	return &scannedEntity, nil
}
