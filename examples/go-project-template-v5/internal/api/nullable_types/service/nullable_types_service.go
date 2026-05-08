package service

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/nullable_types/dto"
)

type NullableTypesService interface {
	Save(ctx context.Context, input *dto.NullableTypesCreateDto) (*dto.NullableTypesDto, error)
	UpdateByID(ctx context.Context, input *dto.NullableTypesUpdateDto, pkID int64) (*dto.NullableTypesDto, error)
	DeleteByID(ctx context.Context, pkID int64) error
	FindByID(ctx context.Context, pkID int64) (*dto.NullableTypesDto, error)
	FindAll(ctx context.Context) ([]dto.NullableTypesDto, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.NullableTypesDto, pageable.Page, error)
}
