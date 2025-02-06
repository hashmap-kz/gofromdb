package repository

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	dbModel "go-project-template-v5/internal/api/purchase_workflow/entity/postgres"
)

type PurchaseWorkflowRepository interface {
	Save(ctx context.Context, inputEntity *dbModel.PurchaseWorkflow) (*dbModel.PurchaseWorkflow, error)
	UpdateByID(ctx context.Context, inputEntity *dbModel.PurchaseWorkflow, pkRecordID int) (*dbModel.PurchaseWorkflow, error)
	DeleteByID(ctx context.Context, pkRecordID int) error
	FindByID(ctx context.Context, pkRecordID int) (*dbModel.PurchaseWorkflow, error)
	FindAll(ctx context.Context) ([]dbModel.PurchaseWorkflow, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.PurchaseWorkflow, pageable.Page, error)
}
