package impl

import (
	"context"
	"fmt"

	dbModel "go-project-template-v5/internal/api/composite_pk/entity/postgres"

	"go-project-template-v5/internal/api/composite_pk/repository"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/pkg/storage/postgres"

	"github.com/jackc/pgx/v5"
)

type compositePkRepository struct {
	db *postgres.Postgres
}

var _ repository.CompositePkRepository = &compositePkRepository{}

func NewCompositePkRepository(_ context.Context, db *postgres.Postgres) repository.CompositePkRepository {
	return &compositePkRepository{
		db: db,
	}
}

func (r *compositePkRepository) Save(ctx context.Context, inputEntity *dbModel.CompositePk) (*dbModel.CompositePk, error) {
	tag := "compositePkRepository.Save"

	query := `		
		insert into public.composite_pk (
			tenant_id,
			code,
			name
		)
		values ($1, $2, $3)
		returning
			tenant_id,
			code,
			name
		`

	row := r.db.Pool.QueryRow(ctx, query,
		inputEntity.TenantID,
		inputEntity.Code,
		inputEntity.Name,
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *compositePkRepository) UpdateByID(ctx context.Context, inputEntity *dbModel.CompositePk, pkTenantID int64, pkCode string) (*dbModel.CompositePk, error) {
	tag := "compositePkRepository.UpdateByID"

	query := `		
		update public.composite_pk
		set 
			name = coalesce(nullif($3, ''), name)
		where tenant_id = $1 and code = $2
		returning 
			tenant_id,
			code,
			name
		`

	row := r.db.Pool.QueryRow(ctx, query,
		pkTenantID, pkCode,
		inputEntity.Name,
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *compositePkRepository) DeleteByID(ctx context.Context, pkTenantID int64, pkCode string) error {
	tag := "compositePkRepository.DeleteByID"

	query := `		
		delete from only public.composite_pk
		where tenant_id = $1 and code = $2
		`

	cmdTag, err := r.db.Pool.Exec(ctx, query, pkTenantID, pkCode)
	if err != nil || cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("%s. no rows were deleted: %w", tag, err)
	}
	return nil
}

func (r *compositePkRepository) FindByID(ctx context.Context, pkTenantID int64, pkCode string) (*dbModel.CompositePk, error) {
	tag := "compositePkRepository.FindByID"

	query := `		
		select
			tenant_id,
			code,
			name
		from public.composite_pk
		where tenant_id = $1 and code = $2
		order by tenant_id, code
		`

	row := r.db.Pool.QueryRow(ctx, query, pkTenantID, pkCode)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *compositePkRepository) FindAll(ctx context.Context) ([]dbModel.CompositePk, error) {
	tag := "compositePkRepository.FindAll"

	query := `		
		select
			tenant_id,
			code,
			name
		from public.composite_pk
		order by tenant_id, code
		`

	rows, err := r.db.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	var scannedEntities []dbModel.CompositePk
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

func (r *compositePkRepository) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.CompositePk, pageable.Page, error) {
	tag := "compositePkRepository.FindAllPageable"

	// retrieve total count
	queryCnt := `select count(*) from public.composite_pk`
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
			tenant_id,
			code,
			name
		from public.composite_pk
		order by tenant_id, code
		offset $1 limit $2
		`

	rows, err := r.db.Pool.Query(ctx, query, pq.GetOffset(), pq.GetLimit())
	if err != nil {
		return nil, pageable.Page{}, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	var scannedEntities []dbModel.CompositePk
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
func scanFullRow(row pgx.Row) (*dbModel.CompositePk, error) {
	var scannedEntity dbModel.CompositePk
	err := row.Scan(
		&scannedEntity.TenantID,
		&scannedEntity.Code,
		&scannedEntity.Name,
	)
	if err != nil {
		return nil, err
	}
	return &scannedEntity, nil
}
