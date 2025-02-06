package repository

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	dbModel "go-project-template-v5/internal/api/product_prices/entity/postgres"
)

type ProductPricesRepository interface {
	Save(ctx context.Context, inputEntity *dbModel.ProductPrices) (*dbModel.ProductPrices, error)
	UpdateByID(ctx context.Context, inputEntity *dbModel.ProductPrices, pkRecordID int) (*dbModel.ProductPrices, error)
	DeleteByID(ctx context.Context, pkRecordID int) error
	FindByID(ctx context.Context, pkRecordID int) (*dbModel.ProductPrices, error)
	FindAll(ctx context.Context) ([]dbModel.ProductPrices, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.ProductPrices, pageable.Page, error)
}
