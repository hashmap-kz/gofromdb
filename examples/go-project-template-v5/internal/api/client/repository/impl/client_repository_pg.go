package impl

import (
	"context"
	"fmt"

	dbModel "go-project-template-v5/internal/api/client/entity/postgres"

	"go-project-template-v5/internal/api/client/repository"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/pkg/storage/postgres"

	"github.com/jackc/pgx/v5"
)

type clientRepository struct {
	db *postgres.Postgres
}

var _ repository.ClientRepository = &clientRepository{}

func NewClientRepository(_ context.Context, db *postgres.Postgres) repository.ClientRepository {
	return &clientRepository{
		db: db,
	}
}

func (r *clientRepository) Save(ctx context.Context, inputEntity *dbModel.Client) (*dbModel.Client, error) {
	tag := "clientRepository.Save"

	query := `		
		insert into public.client (
			email
		)
		values ($1)
		returning
			record_id,
			email,
			created_at,
			updated_at,
			guid
		`

	row := r.db.Pool.QueryRow(ctx, query,
		inputEntity.Email,
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *clientRepository) UpdateByID(ctx context.Context, entityId int, inputEntity *dbModel.Client) (*dbModel.Client, error) {
	tag := "clientRepository.UpdateByID"

	query := `		
		update public.client
		set 
			email = coalesce(nullif($2, ''), email)
		where record_id = $1
		returning 
			record_id,
			email,
			created_at,
			updated_at,
			guid
		`

	row := r.db.Pool.QueryRow(ctx, query,
		entityId,
		inputEntity.Email,
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *clientRepository) DeleteByID(ctx context.Context, entityId int) error {
	tag := "clientRepository.DeleteByID"

	query := `		
		delete from only public.client
		where record_id = $1
		`

	cmdTag, err := r.db.Pool.Exec(ctx, query, entityId)
	if err != nil || cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("%s. no rows deleted for id: %v, %w", tag, entityId, err)
	}
	return nil
}

func (r *clientRepository) FindByID(ctx context.Context, entityId int) (*dbModel.Client, error) {
	tag := "clientRepository.FindByID"

	query := `		
		select
			record_id,
			email,
			created_at,
			updated_at,
			guid
		from public.client
		where record_id = $1
		order by record_id
		`

	row := r.db.Pool.QueryRow(ctx, query, entityId)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *clientRepository) FindAll(ctx context.Context) ([]dbModel.Client, error) {
	tag := "clientRepository.FindAll"

	query := `		
		select
			record_id,
			email,
			created_at,
			updated_at,
			guid
		from public.client
		order by record_id
		`

	rows, err := r.db.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	var scannedEntities []dbModel.Client
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

func (r *clientRepository) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.Client, pageable.Page, error) {
	tag := "clientRepository.FindAllPageable"

	// retrieve total count
	queryCnt := `select count(record_id) from public.client`
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
			email,
			created_at,
			updated_at,
			guid
		from public.client
		order by record_id
		offset $1 limit $2
		`

	rows, err := r.db.Pool.Query(ctx, query, pq.GetOffset(), pq.GetLimit())
	if err != nil {
		return nil, pageable.Page{}, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	var scannedEntities []dbModel.Client
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
func scanFullRow(row pgx.Row) (*dbModel.Client, error) {
	var scannedEntity dbModel.Client
	err := row.Scan(
		&scannedEntity.RecordID,
		&scannedEntity.Email,
		&scannedEntity.CreatedAt,
		&scannedEntity.UpdatedAt,
		&scannedEntity.Guid,
	)
	if err != nil {
		return nil, err
	}
	return &scannedEntity, nil
}
