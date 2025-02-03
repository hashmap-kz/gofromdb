package service

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/product/dto"
)

type ProductService interface {
	Save(ctx context.Context, input *dto.ProductCreateDto) (*dto.ProductDto, error)
	UpdateByID(ctx context.Context, input *dto.ProductUpdateDto, pkRecordID int) (*dto.ProductDto, error)
	DeleteByID(ctx context.Context, pkRecordID int) error
	FindByID(ctx context.Context, pkRecordID int) (*dto.ProductDto, error)
	FindAll(ctx context.Context) ([]dto.ProductDto, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.ProductDto, pageable.Page, error)
}
