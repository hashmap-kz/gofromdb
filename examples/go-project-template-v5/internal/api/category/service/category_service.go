package service

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/category/dto"
)

type CategoryService interface {
	Save(ctx context.Context, input *dto.CategoryCreateDto) (*dto.CategoryDto, error)
	GetAll(ctx context.Context) ([]dto.CategoryDto, error)
	GetAllPaginated(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.CategoryDto, pageable.Page, error)
	Update(ctx context.Context, input *dto.CategoryUpdateDto) (*dto.CategoryDto, error)
	Delete(ctx context.Context, id int) error
	GetByID(ctx context.Context, id int) (*dto.CategoryDto, error)
}
