package repository

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	dbModel "go-project-template-v5/internal/api/customer_orders/entity/postgres"
)

type CustomerOrdersRepository interface {
	Save(ctx context.Context, inputEntity *dbModel.CustomerOrders) (*dbModel.CustomerOrders, error)
	UpdateByID(ctx context.Context, inputEntity *dbModel.CustomerOrders, pkRecordID int) (*dbModel.CustomerOrders, error)
	DeleteByID(ctx context.Context, pkRecordID int) error
	FindByID(ctx context.Context, pkRecordID int) (*dbModel.CustomerOrders, error)
	FindAll(ctx context.Context) ([]dbModel.CustomerOrders, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.CustomerOrders, pageable.Page, error)
}
