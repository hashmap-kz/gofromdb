package repository

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	dbModel "go-project-template-v5/internal/api/category/entity/postgres"
)

type CategoryRepository interface {
	Save(ctx context.Context, inputEntity *dbModel.Category) (*dbModel.Category, error)
	UpdateByID(ctx context.Context, entityId int, inputEntity *dbModel.Category) (*dbModel.Category, error)
	DeleteByID(ctx context.Context, entityId int) error
	FindByID(ctx context.Context, entityId int) (*dbModel.Category, error)
	FindAll(ctx context.Context) ([]dbModel.Category, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.Category, pageable.Page, error)
}
