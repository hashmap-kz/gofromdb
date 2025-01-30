package repository

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	dbModel "go-project-template-v5/internal/api/buy/entity/postgres"
)

type BuyRepository interface {
	Save(ctx context.Context, inputEntity *dbModel.Buy) (*dbModel.Buy, error)
	GetAll(ctx context.Context) ([]dbModel.Buy, error)
	GetAllPaginated(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.Buy, pageable.Page, error)
	Update(ctx context.Context, entityId int, inputEntity *dbModel.Buy) (*dbModel.Buy, error)
	Delete(ctx context.Context, entityId int) error
	GetByID(ctx context.Context, entityId int) (*dbModel.Buy, error)
}
