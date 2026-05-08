package impl

import (
	"context"
	"fmt"

	dbModel "go-project-template-v5/internal/api/orders/entity/postgres"

	"go-project-template-v5/internal/api/orders/repository"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/pkg/storage/postgres"

	"github.com/jackc/pgx/v5"
)

type ordersRepository struct {
	db *postgres.Postgres
}

var _ repository.OrdersRepository = &ordersRepository{}

func NewOrdersRepository(_ context.Context, db *postgres.Postgres) repository.OrdersRepository {
	return &ordersRepository{
		db: db,
	}
}

func (r *ordersRepository) Save(ctx context.Context, inputEntity *dbModel.Orders) (*dbModel.Orders, error) {
	tag := "ordersRepository.Save"

	query := `		
		insert into public.orders (
			user_id,
			description
		)
		values ($1, $2)
		returning
			record_id,
			user_id,
			description,
			created_at,
			updated_at,
			guid
		`

	row := r.db.Pool.QueryRow(ctx, query,
		inputEntity.UserID,
		inputEntity.Description,
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *ordersRepository) UpdateByID(ctx context.Context, inputEntity *dbModel.Orders, pkRecordID int) (*dbModel.Orders, error) {
	tag := "ordersRepository.UpdateByID"

	query := `		
		update public.orders
		set
			user_id     = coalesce(nullif($2, 0::int4), user_id),
			description = coalesce(nullif($3, ''), description)
		where record_id = $1
		returning
			record_id,
			user_id,
			description,
			created_at,
			updated_at,
			guid
		`

	row := r.db.Pool.QueryRow(ctx, query,
		pkRecordID,
		inputEntity.UserID,
		inputEntity.Description,
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *ordersRepository) DeleteByID(ctx context.Context, pkRecordID int) error {
	tag := "ordersRepository.DeleteByID"

	query := `		
		delete from only public.orders
		where record_id = $1
		`

	cmdTag, err := r.db.Pool.Exec(ctx, query, pkRecordID)
	if err != nil || cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("%s. no rows were deleted: %w", tag, err)
	}
	return nil
}

func (r *ordersRepository) FindByID(ctx context.Context, pkRecordID int) (*dbModel.Orders, error) {
	tag := "ordersRepository.FindByID"

	query := `		
		select
			record_id,
			user_id,
			description,
			created_at,
			updated_at,
			guid
		from public.orders
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

func (r *ordersRepository) FindAll(ctx context.Context) ([]dbModel.Orders, error) {
	tag := "ordersRepository.FindAll"

	query := `		
		select
			record_id,
			user_id,
			description,
			created_at,
			updated_at,
			guid
		from public.orders
		order by record_id
		`

	rows, err := r.db.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	var scannedEntities []dbModel.Orders
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

func (r *ordersRepository) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.Orders, pageable.Page, error) {
	tag := "ordersRepository.FindAllPageable"

	// retrieve total count
	queryCnt := `select count(*) from public.orders`
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
			user_id,
			description,
			created_at,
			updated_at,
			guid
		from public.orders
		order by record_id
		offset $1 limit $2
		`

	rows, err := r.db.Pool.Query(ctx, query, pq.GetOffset(), pq.GetLimit())
	if err != nil {
		return nil, pageable.Page{}, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	var scannedEntities []dbModel.Orders
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
func scanFullRow(row pgx.Row) (*dbModel.Orders, error) {
	var scannedEntity dbModel.Orders
	err := row.Scan(
		&scannedEntity.RecordID,
		&scannedEntity.UserID,
		&scannedEntity.Description,
		&scannedEntity.CreatedAt,
		&scannedEntity.UpdatedAt,
		&scannedEntity.GUID,
	)
	if err != nil {
		return nil, err
	}
	return &scannedEntity, nil
}
