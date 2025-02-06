package repository

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	dbModel "go-project-template-v5/internal/api/categories/entity/postgres"
)

type CategoriesRepository interface {
	Save(ctx context.Context, inputEntity *dbModel.Categories) (*dbModel.Categories, error)
	UpdateByID(ctx context.Context, inputEntity *dbModel.Categories, pkRecordID int) (*dbModel.Categories, error)
	DeleteByID(ctx context.Context, pkRecordID int) error
	FindByID(ctx context.Context, pkRecordID int) (*dbModel.Categories, error)
	FindAll(ctx context.Context) ([]dbModel.Categories, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.Categories, pageable.Page, error)
}
