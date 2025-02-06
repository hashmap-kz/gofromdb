package service

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/products/dto"
)

type ProductsService interface {
	Save(ctx context.Context, input *dto.ProductsCreateDto) (*dto.ProductsDto, error)
	UpdateByID(ctx context.Context, input *dto.ProductsUpdateDto, pkRecordID int) (*dto.ProductsDto, error)
	DeleteByID(ctx context.Context, pkRecordID int) error
	FindByID(ctx context.Context, pkRecordID int) (*dto.ProductsDto, error)
	FindAll(ctx context.Context) ([]dto.ProductsDto, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.ProductsDto, pageable.Page, error)
}
