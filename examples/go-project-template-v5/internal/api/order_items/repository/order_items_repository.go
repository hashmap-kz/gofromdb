package repository

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	dbModel "go-project-template-v5/internal/api/order_items/entity/postgres"
)

type OrderItemsRepository interface {
	Save(ctx context.Context, inputEntity *dbModel.OrderItems) (*dbModel.OrderItems, error)
	UpdateByID(ctx context.Context, inputEntity *dbModel.OrderItems, pkRecordID int) (*dbModel.OrderItems, error)
	DeleteByID(ctx context.Context, pkRecordID int) error
	FindByID(ctx context.Context, pkRecordID int) (*dbModel.OrderItems, error)
	FindAll(ctx context.Context) ([]dbModel.OrderItems, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.OrderItems, pageable.Page, error)
}
