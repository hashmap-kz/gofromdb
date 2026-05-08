package impl

import (
	"context"
	"fmt"

	dbModel "go-project-template-v5/internal/api/serial_pk/entity/postgres"

	"go-project-template-v5/internal/api/serial_pk/repository"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/pkg/storage/postgres"

	"github.com/jackc/pgx/v5"
)

type serialPkRepository struct {
	db *postgres.Postgres
}

var _ repository.SerialPkRepository = &serialPkRepository{}

func NewSerialPkRepository(_ context.Context, db *postgres.Postgres) repository.SerialPkRepository {
	return &serialPkRepository{
		db: db,
	}
}

func (r *serialPkRepository) Save(ctx context.Context, inputEntity *dbModel.SerialPk) (*dbModel.SerialPk, error) {
	tag := "serialPkRepository.Save"

	query := `		
		insert into public.serial_pk (
			name
		)
		values ($1)
		returning
			id,
			name
		`

	row := r.db.Pool.QueryRow(ctx, query,
		inputEntity.Name,
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *serialPkRepository) UpdateByID(ctx context.Context, inputEntity *dbModel.SerialPk, pkID int64) (*dbModel.SerialPk, error) {
	tag := "serialPkRepository.UpdateByID"

	query := `		
		update public.serial_pk
		set 
			name = coalesce(nullif($2, ''), name)
		where id = $1
		returning 
			id,
			name
		`

	row := r.db.Pool.QueryRow(ctx, query,
		pkID,
		inputEntity.Name,
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *serialPkRepository) DeleteByID(ctx context.Context, pkID int64) error {
	tag := "serialPkRepository.DeleteByID"

	query := `		
		delete from only public.serial_pk
		where id = $1
		`

	cmdTag, err := r.db.Pool.Exec(ctx, query, pkID)
	if err != nil || cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("%s. no rows were deleted: %w", tag, err)
	}
	return nil
}

func (r *serialPkRepository) FindByID(ctx context.Context, pkID int64) (*dbModel.SerialPk, error) {
	tag := "serialPkRepository.FindByID"

	query := `		
		select
			id,
			name
		from public.serial_pk
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

func (r *serialPkRepository) FindAll(ctx context.Context) ([]dbModel.SerialPk, error) {
	tag := "serialPkRepository.FindAll"

	query := `		
		select
			id,
			name
		from public.serial_pk
		order by id
		`

	rows, err := r.db.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	var scannedEntities []dbModel.SerialPk
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

func (r *serialPkRepository) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.SerialPk, pageable.Page, error) {
	tag := "serialPkRepository.FindAllPageable"

	// retrieve total count
	queryCnt := `select count(*) from public.serial_pk`
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
			name
		from public.serial_pk
		order by id
		offset $1 limit $2
		`

	rows, err := r.db.Pool.Query(ctx, query, pq.GetOffset(), pq.GetLimit())
	if err != nil {
		return nil, pageable.Page{}, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	var scannedEntities []dbModel.SerialPk
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
func scanFullRow(row pgx.Row) (*dbModel.SerialPk, error) {
	var scannedEntity dbModel.SerialPk
	err := row.Scan(
		&scannedEntity.ID,
		&scannedEntity.Name,
	)
	if err != nil {
		return nil, err
	}
	return &scannedEntity, nil
}
