package repository

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	dbModel "go-project-template-v5/internal/api/purchase_items/entity/postgres"
)

type PurchaseItemsRepository interface {
	Save(ctx context.Context, inputEntity *dbModel.PurchaseItems) (*dbModel.PurchaseItems, error)
	UpdateByID(ctx context.Context, inputEntity *dbModel.PurchaseItems, pkRecordID int) (*dbModel.PurchaseItems, error)
	DeleteByID(ctx context.Context, pkRecordID int) error
	FindByID(ctx context.Context, pkRecordID int) (*dbModel.PurchaseItems, error)
	FindAll(ctx context.Context) ([]dbModel.PurchaseItems, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.PurchaseItems, pageable.Page, error)
}
