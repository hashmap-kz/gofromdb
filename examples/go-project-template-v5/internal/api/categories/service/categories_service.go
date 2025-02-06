package service

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/categories/dto"
)

type CategoriesService interface {
	Save(ctx context.Context, input *dto.CategoriesCreateDto) (*dto.CategoriesDto, error)
	UpdateByID(ctx context.Context, input *dto.CategoriesUpdateDto, pkRecordID int) (*dto.CategoriesDto, error)
	DeleteByID(ctx context.Context, pkRecordID int) error
	FindByID(ctx context.Context, pkRecordID int) (*dto.CategoriesDto, error)
	FindAll(ctx context.Context) ([]dto.CategoriesDto, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.CategoriesDto, pageable.Page, error)
}
