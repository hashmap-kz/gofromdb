package impl

import (
	"context"
	"fmt"

	dbModel "go-project-template-v5/internal/api/customer_orders/entity/postgres"

	"go-project-template-v5/internal/api/customer_orders/repository"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/pkg/storage/postgres"

	"github.com/jackc/pgx/v5"
)

type customerOrdersRepository struct {
	db *postgres.Postgres
}

var _ repository.CustomerOrdersRepository = &customerOrdersRepository{}

func NewCustomerOrdersRepository(_ context.Context, db *postgres.Postgres) repository.CustomerOrdersRepository {
	return &customerOrdersRepository{
		db: db,
	}
}

func (r *customerOrdersRepository) Save(ctx context.Context, inputEntity *dbModel.CustomerOrders) (*dbModel.CustomerOrders, error) {
	tag := "customerOrdersRepository.Save"

	query := `		
		insert into public.customer_orders (
			client_id,
			description
		)
		values ($1, $2)
		returning
			record_id,
			client_id,
			description,
			created_at,
			updated_at,
			guid
		`

	row := r.db.Pool.QueryRow(ctx, query,
		inputEntity.ClientID,
		inputEntity.Description,
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *customerOrdersRepository) UpdateByID(ctx context.Context, inputEntity *dbModel.CustomerOrders, pkRecordID int) (*dbModel.CustomerOrders, error) {
	tag := "customerOrdersRepository.UpdateByID"

	query := `		
		update public.customer_orders
		set 
			client_id   = coalesce(nullif($2, 0::int4), client_id),
			description = coalesce(nullif($3, ''), description)
		where record_id = $1
		returning 
			record_id,
			client_id,
			description,
			created_at,
			updated_at,
			guid
		`

	row := r.db.Pool.QueryRow(ctx, query,
		pkRecordID,
		inputEntity.ClientID,
		inputEntity.Description,
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *customerOrdersRepository) DeleteByID(ctx context.Context, pkRecordID int) error {
	tag := "customerOrdersRepository.DeleteByID"

	query := `		
		delete from only public.customer_orders
		where record_id = $1
		`

	cmdTag, err := r.db.Pool.Exec(ctx, query, pkRecordID)
	if err != nil || cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("%s. no rows were deleted: %w", tag, err)
	}
	return nil
}

func (r *customerOrdersRepository) FindByID(ctx context.Context, pkRecordID int) (*dbModel.CustomerOrders, error) {
	tag := "customerOrdersRepository.FindByID"

	query := `		
		select
			record_id,
			client_id,
			description,
			created_at,
			updated_at,
			guid
		from public.customer_orders
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

func (r *customerOrdersRepository) FindAll(ctx context.Context) ([]dbModel.CustomerOrders, error) {
	tag := "customerOrdersRepository.FindAll"

	query := `		
		select
			record_id,
			client_id,
			description,
			created_at,
			updated_at,
			guid
		from public.customer_orders
		order by record_id
		`

	rows, err := r.db.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	var scannedEntities []dbModel.CustomerOrders
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

func (r *customerOrdersRepository) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.CustomerOrders, pageable.Page, error) {
	tag := "customerOrdersRepository.FindAllPageable"

	// retrieve total count
	queryCnt := `select count(*) from public.customer_orders`
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
			client_id,
			description,
			created_at,
			updated_at,
			guid
		from public.customer_orders
		order by record_id
		offset $1 limit $2
		`

	rows, err := r.db.Pool.Query(ctx, query, pq.GetOffset(), pq.GetLimit())
	if err != nil {
		return nil, pageable.Page{}, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	var scannedEntities []dbModel.CustomerOrders
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
func scanFullRow(row pgx.Row) (*dbModel.CustomerOrders, error) {
	var scannedEntity dbModel.CustomerOrders
	err := row.Scan(
		&scannedEntity.RecordID,
		&scannedEntity.ClientID,
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
