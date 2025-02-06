package service

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/product_prices/dto"
)

type ProductPricesService interface {
	Save(ctx context.Context, input *dto.ProductPricesCreateDto) (*dto.ProductPricesDto, error)
	UpdateByID(ctx context.Context, input *dto.ProductPricesUpdateDto, pkRecordID int) (*dto.ProductPricesDto, error)
	DeleteByID(ctx context.Context, pkRecordID int) error
	FindByID(ctx context.Context, pkRecordID int) (*dto.ProductPricesDto, error)
	FindAll(ctx context.Context) ([]dto.ProductPricesDto, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.ProductPricesDto, pageable.Page, error)
}
