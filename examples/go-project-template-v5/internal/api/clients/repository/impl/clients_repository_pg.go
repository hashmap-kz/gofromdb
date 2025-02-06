package impl

import (
	"context"
	"fmt"

	dbModel "go-project-template-v5/internal/api/clients/entity/postgres"

	"go-project-template-v5/internal/api/clients/repository"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/pkg/storage/postgres"

	"github.com/jackc/pgx/v5"
)

type clientsRepository struct {
	db *postgres.Postgres
}

var _ repository.ClientsRepository = &clientsRepository{}

func NewClientsRepository(_ context.Context, db *postgres.Postgres) repository.ClientsRepository {
	return &clientsRepository{
		db: db,
	}
}

func (r *clientsRepository) Save(ctx context.Context, inputEntity *dbModel.Clients) (*dbModel.Clients, error) {
	tag := "clientsRepository.Save"

	query := `		
		insert into public.clients (
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

func (r *clientsRepository) UpdateByID(ctx context.Context, inputEntity *dbModel.Clients, pkRecordID int) (*dbModel.Clients, error) {
	tag := "clientsRepository.UpdateByID"

	query := `		
		update public.clients
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
		pkRecordID,
		inputEntity.Email,
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *clientsRepository) DeleteByID(ctx context.Context, pkRecordID int) error {
	tag := "clientsRepository.DeleteByID"

	query := `		
		delete from only public.clients
		where record_id = $1
		`

	cmdTag, err := r.db.Pool.Exec(ctx, query, pkRecordID)
	if err != nil || cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("%s. no rows were deleted: %w", tag, err)
	}
	return nil
}

func (r *clientsRepository) FindByID(ctx context.Context, pkRecordID int) (*dbModel.Clients, error) {
	tag := "clientsRepository.FindByID"

	query := `		
		select
			record_id,
			email,
			created_at,
			updated_at,
			guid
		from public.clients
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

func (r *clientsRepository) FindAll(ctx context.Context) ([]dbModel.Clients, error) {
	tag := "clientsRepository.FindAll"

	query := `		
		select
			record_id,
			email,
			created_at,
			updated_at,
			guid
		from public.clients
		order by record_id
		`

	rows, err := r.db.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	var scannedEntities []dbModel.Clients
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

func (r *clientsRepository) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.Clients, pageable.Page, error) {
	tag := "clientsRepository.FindAllPageable"

	// retrieve total count
	queryCnt := `select count(*) from public.clients`
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
		from public.clients
		order by record_id
		offset $1 limit $2
		`

	rows, err := r.db.Pool.Query(ctx, query, pq.GetOffset(), pq.GetLimit())
	if err != nil {
		return nil, pageable.Page{}, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	var scannedEntities []dbModel.Clients
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
func scanFullRow(row pgx.Row) (*dbModel.Clients, error) {
	var scannedEntity dbModel.Clients
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
