package impl

import (
	"context"
	"fmt"

	dbModel "go-project-template-v5/internal/api/purchase_workflow/entity/postgres"

	"go-project-template-v5/internal/api/purchase_workflow/repository"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/pkg/storage/postgres"

	"github.com/jackc/pgx/v5"
)

type purchaseWorkflowRepository struct {
	db *postgres.Postgres
}

var _ repository.PurchaseWorkflowRepository = &purchaseWorkflowRepository{}

func NewPurchaseWorkflowRepository(_ context.Context, db *postgres.Postgres) repository.PurchaseWorkflowRepository {
	return &purchaseWorkflowRepository{
		db: db,
	}
}

func (r *purchaseWorkflowRepository) Save(ctx context.Context, inputEntity *dbModel.PurchaseWorkflow) (*dbModel.PurchaseWorkflow, error) {
	tag := "purchaseWorkflowRepository.Save"

	query := `		
		insert into public.purchase_workflow (
			valid_period,
			buy_id,
			purchase_step_id
		)
		values ($1, $2, $3)
		returning
			record_id,
			valid_period,
			buy_id,
			purchase_step_id,
			created_at,
			updated_at,
			guid
		`

	row := r.db.Pool.QueryRow(ctx, query,
		inputEntity.ValidPeriod,
		inputEntity.BuyID,
		inputEntity.PurchaseStepID,
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *purchaseWorkflowRepository) UpdateByID(ctx context.Context, inputEntity *dbModel.PurchaseWorkflow, pkRecordID int) (*dbModel.PurchaseWorkflow, error) {
	tag := "purchaseWorkflowRepository.UpdateByID"

	query := `		
		update public.purchase_workflow
		set 
			valid_period     = coalesce(nullif($2, 'empty'::daterange), valid_period),
			buy_id           = coalesce(nullif($3, 0::int4), buy_id),
			purchase_step_id = coalesce(nullif($4, 0::int4), purchase_step_id)
		where record_id = $1
		returning 
			record_id,
			valid_period,
			buy_id,
			purchase_step_id,
			created_at,
			updated_at,
			guid
		`

	row := r.db.Pool.QueryRow(ctx, query,
		pkRecordID,
		inputEntity.ValidPeriod,
		inputEntity.BuyID,
		inputEntity.PurchaseStepID,
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *purchaseWorkflowRepository) DeleteByID(ctx context.Context, pkRecordID int) error {
	tag := "purchaseWorkflowRepository.DeleteByID"

	query := `		
		delete from only public.purchase_workflow
		where record_id = $1
		`

	cmdTag, err := r.db.Pool.Exec(ctx, query, pkRecordID)
	if err != nil || cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("%s. no rows were deleted: %w", tag, err)
	}
	return nil
}

func (r *purchaseWorkflowRepository) FindByID(ctx context.Context, pkRecordID int) (*dbModel.PurchaseWorkflow, error) {
	tag := "purchaseWorkflowRepository.FindByID"

	query := `		
		select
			record_id,
			valid_period,
			buy_id,
			purchase_step_id,
			created_at,
			updated_at,
			guid
		from public.purchase_workflow
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

func (r *purchaseWorkflowRepository) FindAll(ctx context.Context) ([]dbModel.PurchaseWorkflow, error) {
	tag := "purchaseWorkflowRepository.FindAll"

	query := `		
		select
			record_id,
			valid_period,
			buy_id,
			purchase_step_id,
			created_at,
			updated_at,
			guid
		from public.purchase_workflow
		order by record_id
		`

	rows, err := r.db.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	var scannedEntities []dbModel.PurchaseWorkflow
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

func (r *purchaseWorkflowRepository) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.PurchaseWorkflow, pageable.Page, error) {
	tag := "purchaseWorkflowRepository.FindAllPageable"

	// retrieve total count
	queryCnt := `select count(*) from public.purchase_workflow`
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
			valid_period,
			buy_id,
			purchase_step_id,
			created_at,
			updated_at,
			guid
		from public.purchase_workflow
		order by record_id
		offset $1 limit $2
		`

	rows, err := r.db.Pool.Query(ctx, query, pq.GetOffset(), pq.GetLimit())
	if err != nil {
		return nil, pageable.Page{}, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	var scannedEntities []dbModel.PurchaseWorkflow
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
func scanFullRow(row pgx.Row) (*dbModel.PurchaseWorkflow, error) {
	var scannedEntity dbModel.PurchaseWorkflow
	err := row.Scan(
		&scannedEntity.RecordID,
		&scannedEntity.ValidPeriod,
		&scannedEntity.BuyID,
		&scannedEntity.PurchaseStepID,
		&scannedEntity.CreatedAt,
		&scannedEntity.UpdatedAt,
		&scannedEntity.Guid,
	)
	if err != nil {
		return nil, err
	}
	return &scannedEntity, nil
}
