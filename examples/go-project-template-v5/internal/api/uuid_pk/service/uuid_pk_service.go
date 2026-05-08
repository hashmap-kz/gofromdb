package service

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/uuid_pk/dto"
)

type UUIDPkService interface {
	Save(ctx context.Context, input *dto.UUIDPkCreateDto) (*dto.UUIDPkDto, error)
	UpdateByID(ctx context.Context, input *dto.UUIDPkUpdateDto, pkID string) (*dto.UUIDPkDto, error)
	DeleteByID(ctx context.Context, pkID string) error
	FindByID(ctx context.Context, pkID string) (*dto.UUIDPkDto, error)
	FindAll(ctx context.Context) ([]dto.UUIDPkDto, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.UUIDPkDto, pageable.Page, error)
}
