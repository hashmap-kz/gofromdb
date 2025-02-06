package repository

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	dbModel "go-project-template-v5/internal/api/products/entity/postgres"
)

type ProductsRepository interface {
	Save(ctx context.Context, inputEntity *dbModel.Products) (*dbModel.Products, error)
	UpdateByID(ctx context.Context, inputEntity *dbModel.Products, pkRecordID int) (*dbModel.Products, error)
	DeleteByID(ctx context.Context, pkRecordID int) error
	FindByID(ctx context.Context, pkRecordID int) (*dbModel.Products, error)
	FindAll(ctx context.Context) ([]dbModel.Products, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.Products, pageable.Page, error)
}
