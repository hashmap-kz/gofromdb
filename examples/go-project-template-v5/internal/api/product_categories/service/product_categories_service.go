package service

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/product_categories/dto"
)

type ProductCategoriesService interface {
	Save(ctx context.Context, input *dto.ProductCategoriesCreateDto) (*dto.ProductCategoriesDto, error)
	UpdateByID(ctx context.Context, input *dto.ProductCategoriesUpdateDto, pkRecordID int) (*dto.ProductCategoriesDto, error)
	DeleteByID(ctx context.Context, pkRecordID int) error
	FindByID(ctx context.Context, pkRecordID int) (*dto.ProductCategoriesDto, error)
	FindAll(ctx context.Context) ([]dto.ProductCategoriesDto, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.ProductCategoriesDto, pageable.Page, error)
}
