package repository

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	dbModel "go-project-template-v5/internal/api/product/entity/postgres"
)

type ProductRepository interface {
	Save(ctx context.Context, inputEntity *dbModel.Product) (*dbModel.Product, error)
	GetAll(ctx context.Context) ([]dbModel.Product, error)
	GetAllPaginated(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.Product, pageable.Page, error)
	Update(ctx context.Context, entityId int, inputEntity *dbModel.Product) (*dbModel.Product, error)
	Delete(ctx context.Context, entityId int) error
	GetByID(ctx context.Context, entityId int) (*dbModel.Product, error)
}
