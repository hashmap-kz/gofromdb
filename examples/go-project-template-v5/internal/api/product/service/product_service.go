package service

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/product/dto"
)

type ProductService interface {
	Save(ctx context.Context, input *dto.ProductCreateDto) (*dto.ProductDto, error)
	GetAll(ctx context.Context) ([]dto.ProductDto, error)
	GetAllPaginated(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.ProductDto, pageable.Page, error)
	Update(ctx context.Context, entityId int, input *dto.ProductUpdateDto) (*dto.ProductDto, error)
	Delete(ctx context.Context, id int) error
	GetByID(ctx context.Context, id int) (*dto.ProductDto, error)
}
