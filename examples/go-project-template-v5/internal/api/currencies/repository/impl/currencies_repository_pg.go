package impl

import (
	"context"
	"fmt"

	dbModel "go-project-template-v5/internal/api/currencies/entity/postgres"

	"go-project-template-v5/internal/api/currencies/repository"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/pkg/storage/postgres"

	"github.com/jackc/pgx/v5"
)

type currenciesRepository struct {
	db *postgres.Postgres
}

var _ repository.CurrenciesRepository = &currenciesRepository{}

func NewCurrenciesRepository(_ context.Context, db *postgres.Postgres) repository.CurrenciesRepository {
	return &currenciesRepository{
		db: db,
	}
}

func (r *currenciesRepository) Save(ctx context.Context, inputEntity *dbModel.Currencies) (*dbModel.Currencies, error) {
	tag := "currenciesRepository.Save"

	query := `		
		insert into public.currencies (
			currency_code,
			currency_value
		)
		values ($1, $2)
		returning
			record_id,
			currency_code,
			currency_value,
			created_at,
			updated_at,
			guid
		`

	row := r.db.Pool.QueryRow(ctx, query,
		inputEntity.CurrencyCode,
		inputEntity.CurrencyValue,
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *currenciesRepository) UpdateByID(ctx context.Context, inputEntity *dbModel.Currencies, pkRecordID int) (*dbModel.Currencies, error) {
	tag := "currenciesRepository.UpdateByID"

	query := `		
		update public.currencies
		set 
			currency_code  = coalesce(nullif($2, ''), currency_code),
			currency_value = coalesce(nullif($3, ''), currency_value)
		where record_id = $1
		returning 
			record_id,
			currency_code,
			currency_value,
			created_at,
			updated_at,
			guid
		`

	row := r.db.Pool.QueryRow(ctx, query,
		pkRecordID,
		inputEntity.CurrencyCode,
		inputEntity.CurrencyValue,
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *currenciesRepository) DeleteByID(ctx context.Context, pkRecordID int) error {
	tag := "currenciesRepository.DeleteByID"

	query := `		
		delete from only public.currencies
		where record_id = $1
		`

	cmdTag, err := r.db.Pool.Exec(ctx, query, pkRecordID)
	if err != nil || cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("%s. no rows were deleted: %w", tag, err)
	}
	return nil
}

func (r *currenciesRepository) FindByID(ctx context.Context, pkRecordID int) (*dbModel.Currencies, error) {
	tag := "currenciesRepository.FindByID"

	query := `		
		select
			record_id,
			currency_code,
			currency_value,
			created_at,
			updated_at,
			guid
		from public.currencies
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

func (r *currenciesRepository) FindAll(ctx context.Context) ([]dbModel.Currencies, error) {
	tag := "currenciesRepository.FindAll"

	query := `		
		select
			record_id,
			currency_code,
			currency_value,
			created_at,
			updated_at,
			guid
		from public.currencies
		order by record_id
		`

	rows, err := r.db.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	var scannedEntities []dbModel.Currencies
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

func (r *currenciesRepository) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.Currencies, pageable.Page, error) {
	tag := "currenciesRepository.FindAllPageable"

	// retrieve total count
	queryCnt := `select count(*) from public.currencies`
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
			currency_code,
			currency_value,
			created_at,
			updated_at,
			guid
		from public.currencies
		order by record_id
		offset $1 limit $2
		`

	rows, err := r.db.Pool.Query(ctx, query, pq.GetOffset(), pq.GetLimit())
	if err != nil {
		return nil, pageable.Page{}, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	var scannedEntities []dbModel.Currencies
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
func scanFullRow(row pgx.Row) (*dbModel.Currencies, error) {
	var scannedEntity dbModel.Currencies
	err := row.Scan(
		&scannedEntity.RecordID,
		&scannedEntity.CurrencyCode,
		&scannedEntity.CurrencyValue,
		&scannedEntity.CreatedAt,
		&scannedEntity.UpdatedAt,
		&scannedEntity.Guid,
	)
	if err != nil {
		return nil, err
	}
	return &scannedEntity, nil
}
