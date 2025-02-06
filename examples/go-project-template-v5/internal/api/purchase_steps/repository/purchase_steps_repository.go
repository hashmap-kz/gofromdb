package repository

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	dbModel "go-project-template-v5/internal/api/purchase_steps/entity/postgres"
)

type PurchaseStepsRepository interface {
	Save(ctx context.Context, inputEntity *dbModel.PurchaseSteps) (*dbModel.PurchaseSteps, error)
	UpdateByID(ctx context.Context, inputEntity *dbModel.PurchaseSteps, pkRecordID int) (*dbModel.PurchaseSteps, error)
	DeleteByID(ctx context.Context, pkRecordID int) error
	FindByID(ctx context.Context, pkRecordID int) (*dbModel.PurchaseSteps, error)
	FindAll(ctx context.Context) ([]dbModel.PurchaseSteps, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.PurchaseSteps, pageable.Page, error)
}
