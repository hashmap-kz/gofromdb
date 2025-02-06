package impl

import (
	"context"
	"fmt"

	dbModel "go-project-template-v5/internal/api/users/entity/postgres"

	"go-project-template-v5/internal/api/users/repository"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/pkg/storage/postgres"

	"github.com/jackc/pgx/v5"
)

type usersRepository struct {
	db *postgres.Postgres
}

var _ repository.UsersRepository = &usersRepository{}

func NewUsersRepository(_ context.Context, db *postgres.Postgres) repository.UsersRepository {
	return &usersRepository{
		db: db,
	}
}

func (r *usersRepository) Save(ctx context.Context, inputEntity *dbModel.Users) (*dbModel.Users, error) {
	tag := "usersRepository.Save"

	query := `		
		insert into public.users (
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

func (r *usersRepository) UpdateByID(ctx context.Context, inputEntity *dbModel.Users, pkRecordID int) (*dbModel.Users, error) {
	tag := "usersRepository.UpdateByID"

	query := `		
		update public.users
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

func (r *usersRepository) DeleteByID(ctx context.Context, pkRecordID int) error {
	tag := "usersRepository.DeleteByID"

	query := `		
		delete from only public.users
		where record_id = $1
		`

	cmdTag, err := r.db.Pool.Exec(ctx, query, pkRecordID)
	if err != nil || cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("%s. no rows were deleted: %w", tag, err)
	}
	return nil
}

func (r *usersRepository) FindByID(ctx context.Context, pkRecordID int) (*dbModel.Users, error) {
	tag := "usersRepository.FindByID"

	query := `		
		select
			record_id,
			email,
			created_at,
			updated_at,
			guid
		from public.users
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

func (r *usersRepository) FindAll(ctx context.Context) ([]dbModel.Users, error) {
	tag := "usersRepository.FindAll"

	query := `		
		select
			record_id,
			email,
			created_at,
			updated_at,
			guid
		from public.users
		order by record_id
		`

	rows, err := r.db.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	var scannedEntities []dbModel.Users
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

func (r *usersRepository) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.Users, pageable.Page, error) {
	tag := "usersRepository.FindAllPageable"

	// retrieve total count
	queryCnt := `select count(*) from public.users`
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
		from public.users
		order by record_id
		offset $1 limit $2
		`

	rows, err := r.db.Pool.Query(ctx, query, pq.GetOffset(), pq.GetLimit())
	if err != nil {
		return nil, pageable.Page{}, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	var scannedEntities []dbModel.Users
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
func scanFullRow(row pgx.Row) (*dbModel.Users, error) {
	var scannedEntity dbModel.Users
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
