package service

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/category/dto"
)

type CategoryService interface {
	Save(ctx context.Context, input *dto.CategoryCreateDto) (*dto.CategoryDto, error)
	UpdateByID(ctx context.Context, input *dto.CategoryUpdateDto, pkRecordID int) (*dto.CategoryDto, error)
	DeleteByID(ctx context.Context, pkRecordID int) error
	FindByID(ctx context.Context, pkRecordID int) (*dto.CategoryDto, error)
	FindAll(ctx context.Context) ([]dto.CategoryDto, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.CategoryDto, pageable.Page, error)
}
