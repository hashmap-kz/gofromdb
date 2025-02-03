package repository

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	dbModel "go-project-template-v5/internal/api/product/entity/postgres"
)

type ProductRepository interface {
	Save(ctx context.Context, inputEntity *dbModel.Product) (*dbModel.Product, error)
	UpdateByID(ctx context.Context, inputEntity *dbModel.Product, pkRecordID int) (*dbModel.Product, error)
	DeleteByID(ctx context.Context, pkRecordID int) error
	FindByID(ctx context.Context, pkRecordID int) (*dbModel.Product, error)
	FindAll(ctx context.Context) ([]dbModel.Product, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.Product, pageable.Page, error)
}
