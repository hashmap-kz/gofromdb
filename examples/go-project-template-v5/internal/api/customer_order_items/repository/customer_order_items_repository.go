package repository

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	dbModel "go-project-template-v5/internal/api/customer_order_items/entity/postgres"
)

type CustomerOrderItemsRepository interface {
	Save(ctx context.Context, inputEntity *dbModel.CustomerOrderItems) (*dbModel.CustomerOrderItems, error)
	UpdateByID(ctx context.Context, inputEntity *dbModel.CustomerOrderItems, pkRecordID int) (*dbModel.CustomerOrderItems, error)
	DeleteByID(ctx context.Context, pkRecordID int) error
	FindByID(ctx context.Context, pkRecordID int) (*dbModel.CustomerOrderItems, error)
	FindAll(ctx context.Context) ([]dbModel.CustomerOrderItems, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.CustomerOrderItems, pageable.Page, error)
}
