package repository

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	dbModel "go-project-template-v5/internal/api/buy/entity/postgres"
)

type BuyRepository interface {
	Save(ctx context.Context, inputEntity *dbModel.Buy) (*dbModel.Buy, error)
	UpdateByID(ctx context.Context, entityId int, inputEntity *dbModel.Buy) (*dbModel.Buy, error)
	DeleteByID(ctx context.Context, entityId int) error
	FindByID(ctx context.Context, entityId int) (*dbModel.Buy, error)
	FindAll(ctx context.Context) ([]dbModel.Buy, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.Buy, pageable.Page, error)
}
