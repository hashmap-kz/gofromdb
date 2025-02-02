package impl

import (
	"context"
	"fmt"

	dbModel "go-project-template-v5/internal/api/buy/entity/postgres"

	"go-project-template-v5/internal/api/buy/repository"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/pkg/storage/postgres"

	"github.com/jackc/pgx/v5"
)

type buyRepository struct {
	db *postgres.Postgres
}

var _ repository.BuyRepository = &buyRepository{}

func NewBuyRepository(_ context.Context, db *postgres.Postgres) repository.BuyRepository {
	return &buyRepository{
		db: db,
	}
}

func (r *buyRepository) Save(ctx context.Context, inputEntity *dbModel.Buy) (*dbModel.Buy, error) {
	tag := "buyRepository.Save"

	query := `		
		insert into public.buy (
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

func (r *buyRepository) UpdateByID(ctx context.Context, inputEntity *dbModel.Buy, pkRecordID int) (*dbModel.Buy, error) {
	tag := "buyRepository.UpdateByID"

	query := `		
		update public.buy
		set 
			client_id = coalesce(nullif($2, 0::int4), client_id),
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

func (r *buyRepository) DeleteByID(ctx context.Context, pkRecordID int) error {
	tag := "buyRepository.DeleteByID"

	query := `		
		delete from only public.buy
		where record_id = $1
		`

	cmdTag, err := r.db.Pool.Exec(ctx, query, pkRecordID)
	if err != nil || cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("%s. no rows were deleted: %w", tag, err)
	}
	return nil
}

func (r *buyRepository) FindByID(ctx context.Context, pkRecordID int) (*dbModel.Buy, error) {
	tag := "buyRepository.FindByID"

	query := `		
		select
			record_id,
			client_id,
			description,
			created_at,
			updated_at,
			guid
		from public.buy
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

func (r *buyRepository) FindAll(ctx context.Context) ([]dbModel.Buy, error) {
	tag := "buyRepository.FindAll"

	query := `		
		select
			record_id,
			client_id,
			description,
			created_at,
			updated_at,
			guid
		from public.buy
		order by record_id
		`

	rows, err := r.db.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	var scannedEntities []dbModel.Buy
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

func (r *buyRepository) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.Buy, pageable.Page, error) {
	tag := "buyRepository.FindAllPageable"

	// retrieve total count
	queryCnt := `select count(*) from public.buy`
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
		from public.buy
		order by record_id
		offset $1 limit $2
		`

	rows, err := r.db.Pool.Query(ctx, query, pq.GetOffset(), pq.GetLimit())
	if err != nil {
		return nil, pageable.Page{}, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	var scannedEntities []dbModel.Buy
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
func scanFullRow(row pgx.Row) (*dbModel.Buy, error) {
	var scannedEntity dbModel.Buy
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
