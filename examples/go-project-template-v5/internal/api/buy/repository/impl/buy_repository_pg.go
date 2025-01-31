package impl

import (
	"context"
	"fmt"

	dbModel "go-project-template-v5/internal/api/buy/entity/postgres"

	"go-project-template-v5/internal/api/buy/repository"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/pkg/storage/postgres"
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

	var scannedEntity dbModel.Buy
	err := r.db.Pool.QueryRow(ctx, query,
		inputEntity.ClientID,
		inputEntity.Description,
	).Scan(
		&scannedEntity.RecordID,
		&scannedEntity.ClientID,
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

func (r *buyRepository) UpdateByID(ctx context.Context, entityId int, inputEntity *dbModel.Buy) (*dbModel.Buy, error) {
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

	var scannedEntity dbModel.Buy
	err := r.db.Pool.QueryRow(ctx, query,
		entityId,
		inputEntity.ClientID,
		inputEntity.Description,
	).Scan(
		&scannedEntity.RecordID,
		&scannedEntity.ClientID,
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

func (r *buyRepository) DeleteByID(ctx context.Context, entityId int) error {
	tag := "buyRepository.DeleteByID"

	query := `		
		delete from only public.buy
		where record_id = $1
		`

	cmdTag, err := r.db.Pool.Exec(ctx, query, entityId)
	if err != nil || cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("%s. no rows deleted for id: %v, %w", tag, entityId, err)
	}
	return nil
}

func (r *buyRepository) FindByID(ctx context.Context, entityId int) (*dbModel.Buy, error) {
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

	var scannedEntity dbModel.Buy
	err := r.db.Pool.QueryRow(ctx, query, entityId).Scan(
		&scannedEntity.RecordID,
		&scannedEntity.ClientID,
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
		var scannedEntity dbModel.Buy
		err = rows.Scan(
			&scannedEntity.RecordID,
			&scannedEntity.ClientID,
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

func (r *buyRepository) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.Buy, pageable.Page, error) {
	tag := "buyRepository.FindAllPageable"

	// retrieve total count
	queryCnt := `select count(record_id) from public.buy`
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
		var scannedEntity dbModel.Buy
		err = rows.Scan(
			&scannedEntity.RecordID,
			&scannedEntity.ClientID,
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
