package repository

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	dbModel "go-project-template-v5/internal/api/category/entity/postgres"
)

type CategoryRepository interface {
	Save(ctx context.Context, inputEntity *dbModel.Category) (*dbModel.Category, error)
	GetAll(ctx context.Context) ([]dbModel.Category, error)
	GetAllPaginated(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.Category, pageable.Page, error)
	Update(ctx context.Context, entityId int, inputEntity *dbModel.Category) (*dbModel.Category, error)
	Delete(ctx context.Context, entityId int) error
	GetByID(ctx context.Context, entityId int) (*dbModel.Category, error)
}
