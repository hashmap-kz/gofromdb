package repository

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	dbModel "go-project-template-v5/internal/api/orders/entity/postgres"
)

type OrdersRepository interface {
	Save(ctx context.Context, inputEntity *dbModel.Orders) (*dbModel.Orders, error)
	UpdateByID(ctx context.Context, inputEntity *dbModel.Orders, pkRecordID int) (*dbModel.Orders, error)
	DeleteByID(ctx context.Context, pkRecordID int) error
	FindByID(ctx context.Context, pkRecordID int) (*dbModel.Orders, error)
	FindAll(ctx context.Context) ([]dbModel.Orders, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.Orders, pageable.Page, error)
}
