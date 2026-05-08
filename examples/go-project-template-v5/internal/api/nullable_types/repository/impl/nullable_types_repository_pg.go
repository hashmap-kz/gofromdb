package impl

import (
	"context"
	"fmt"

	dbModel "go-project-template-v5/internal/api/nullable_types/entity/postgres"

	"go-project-template-v5/internal/api/nullable_types/repository"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/pkg/storage/postgres"

	"github.com/jackc/pgx/v5"
)

type nullableTypesRepository struct {
	db *postgres.Postgres
}

var _ repository.NullableTypesRepository = &nullableTypesRepository{}

func NewNullableTypesRepository(_ context.Context, db *postgres.Postgres) repository.NullableTypesRepository {
	return &nullableTypesRepository{
		db: db,
	}
}

func (r *nullableTypesRepository) Save(ctx context.Context, inputEntity *dbModel.NullableTypes) (*dbModel.NullableTypes, error) {
	tag := "nullableTypesRepository.Save"

	query := `		
		insert into public.nullable_types (
			name,
			amount,
			payload,
			tags,
			active
		)
		values ($1, $2, $3, $4, $5)
		returning
			id,
			name,
			amount,
			payload,
			tags,
			active,
			created_at
		`

	row := r.db.Pool.QueryRow(ctx, query,
		inputEntity.Name,
		inputEntity.Amount,
		inputEntity.Payload,
		inputEntity.Tags,
		inputEntity.Active,
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *nullableTypesRepository) UpdateByID(ctx context.Context, inputEntity *dbModel.NullableTypes, pkID int64) (*dbModel.NullableTypes, error) {
	tag := "nullableTypesRepository.UpdateByID"

	query := `		
		update public.nullable_types
		set 
			name    = coalesce(nullif($2, ''), name),
			amount  = coalesce(nullif($3, 0::numeric), amount),
			payload = coalesce(nullif($4, '{}'::jsonb), payload),
			tags    = coalesce(nullif($5, '{}'::text[]), tags),
			active  = coalesce(nullif($6, 'false'::bool), active)
		where id = $1
		returning 
			id,
			name,
			amount,
			payload,
			tags,
			active,
			created_at
		`

	row := r.db.Pool.QueryRow(ctx, query,
		pkID,
		inputEntity.Name,
		inputEntity.Amount,
		inputEntity.Payload,
		inputEntity.Tags,
		inputEntity.Active,
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *nullableTypesRepository) DeleteByID(ctx context.Context, pkID int64) error {
	tag := "nullableTypesRepository.DeleteByID"

	query := `		
		delete from only public.nullable_types
		where id = $1
		`

	cmdTag, err := r.db.Pool.Exec(ctx, query, pkID)
	if err != nil || cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("%s. no rows were deleted: %w", tag, err)
	}
	return nil
}

func (r *nullableTypesRepository) FindByID(ctx context.Context, pkID int64) (*dbModel.NullableTypes, error) {
	tag := "nullableTypesRepository.FindByID"

	query := `		
		select
			id,
			name,
			amount,
			payload,
			tags,
			active,
			created_at
		from public.nullable_types
		where id = $1
		order by id
		`

	row := r.db.Pool.QueryRow(ctx, query, pkID)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *nullableTypesRepository) FindAll(ctx context.Context) ([]dbModel.NullableTypes, error) {
	tag := "nullableTypesRepository.FindAll"

	query := `		
		select
			id,
			name,
			amount,
			payload,
			tags,
			active,
			created_at
		from public.nullable_types
		order by id
		`

	rows, err := r.db.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	var scannedEntities []dbModel.NullableTypes
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

func (r *nullableTypesRepository) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.NullableTypes, pageable.Page, error) {
	tag := "nullableTypesRepository.FindAllPageable"

	// retrieve total count
	queryCnt := `select count(*) from public.nullable_types`
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
			id,
			name,
			amount,
			payload,
			tags,
			active,
			created_at
		from public.nullable_types
		order by id
		offset $1 limit $2
		`

	rows, err := r.db.Pool.Query(ctx, query, pq.GetOffset(), pq.GetLimit())
	if err != nil {
		return nil, pageable.Page{}, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	var scannedEntities []dbModel.NullableTypes
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
func scanFullRow(row pgx.Row) (*dbModel.NullableTypes, error) {
	var scannedEntity dbModel.NullableTypes
	err := row.Scan(
		&scannedEntity.ID,
		&scannedEntity.Name,
		&scannedEntity.Amount,
		&scannedEntity.Payload,
		&scannedEntity.Tags,
		&scannedEntity.Active,
		&scannedEntity.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &scannedEntity, nil
}
