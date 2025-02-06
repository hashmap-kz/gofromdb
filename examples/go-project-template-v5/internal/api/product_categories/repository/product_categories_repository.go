package repository

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	dbModel "go-project-template-v5/internal/api/product_categories/entity/postgres"
)

type ProductCategoriesRepository interface {
	Save(ctx context.Context, inputEntity *dbModel.ProductCategories) (*dbModel.ProductCategories, error)
	UpdateByID(ctx context.Context, inputEntity *dbModel.ProductCategories, pkRecordID int) (*dbModel.ProductCategories, error)
	DeleteByID(ctx context.Context, pkRecordID int) error
	FindByID(ctx context.Context, pkRecordID int) (*dbModel.ProductCategories, error)
	FindAll(ctx context.Context) ([]dbModel.ProductCategories, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.ProductCategories, pageable.Page, error)
}
