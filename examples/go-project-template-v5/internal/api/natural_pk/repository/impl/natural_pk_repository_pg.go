package impl

import (
	"context"
	"fmt"

	dbModel "go-project-template-v5/internal/api/natural_pk/entity/postgres"

	"go-project-template-v5/internal/api/natural_pk/repository"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/pkg/storage/postgres"

	"github.com/jackc/pgx/v5"
)

type naturalPkRepository struct {
	db *postgres.Postgres
}

var _ repository.NaturalPkRepository = &naturalPkRepository{}

func NewNaturalPkRepository(_ context.Context, db *postgres.Postgres) repository.NaturalPkRepository {
	return &naturalPkRepository{
		db: db,
	}
}

func (r *naturalPkRepository) Save(ctx context.Context, inputEntity *dbModel.NaturalPk) (*dbModel.NaturalPk, error) {
	tag := "naturalPkRepository.Save"

	query := `		
		insert into public.natural_pk (
			code,
			name
		)
		values ($1, $2)
		returning
			code,
			name
		`

	row := r.db.Pool.QueryRow(ctx, query,
		inputEntity.Code,
		inputEntity.Name,
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *naturalPkRepository) UpdateByID(ctx context.Context, inputEntity *dbModel.NaturalPk, pkCode string) (*dbModel.NaturalPk, error) {
	tag := "naturalPkRepository.UpdateByID"

	query := `		
		update public.natural_pk
		set
			name = coalesce(nullif($2, ''), name)
		where code = $1
		returning
			code,
			name
		`

	row := r.db.Pool.QueryRow(ctx, query,
		pkCode,
		inputEntity.Name,
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *naturalPkRepository) DeleteByID(ctx context.Context, pkCode string) error {
	tag := "naturalPkRepository.DeleteByID"

	query := `		
		delete from only public.natural_pk
		where code = $1
		`

	cmdTag, err := r.db.Pool.Exec(ctx, query, pkCode)
	if err != nil || cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("%s. no rows were deleted: %w", tag, err)
	}
	return nil
}

func (r *naturalPkRepository) FindByID(ctx context.Context, pkCode string) (*dbModel.NaturalPk, error) {
	tag := "naturalPkRepository.FindByID"

	query := `		
		select
			code,
			name
		from public.natural_pk
		where code = $1
		order by code
		`

	row := r.db.Pool.QueryRow(ctx, query, pkCode)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *naturalPkRepository) FindAll(ctx context.Context) ([]dbModel.NaturalPk, error) {
	tag := "naturalPkRepository.FindAll"

	query := `		
		select
			code,
			name
		from public.natural_pk
		order by code
		`

	rows, err := r.db.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	var scannedEntities []dbModel.NaturalPk
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

func (r *naturalPkRepository) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.NaturalPk, pageable.Page, error) {
	tag := "naturalPkRepository.FindAllPageable"

	// retrieve total count
	queryCnt := `select count(*) from public.natural_pk`
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
			code,
			name
		from public.natural_pk
		order by code
		offset $1 limit $2
		`

	rows, err := r.db.Pool.Query(ctx, query, pq.GetOffset(), pq.GetLimit())
	if err != nil {
		return nil, pageable.Page{}, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	var scannedEntities []dbModel.NaturalPk
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
func scanFullRow(row pgx.Row) (*dbModel.NaturalPk, error) {
	var scannedEntity dbModel.NaturalPk
	err := row.Scan(
		&scannedEntity.Code,
		&scannedEntity.Name,
	)
	if err != nil {
		return nil, err
	}
	return &scannedEntity, nil
}
