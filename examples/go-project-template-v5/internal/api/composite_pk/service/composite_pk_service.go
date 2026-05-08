package service

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/composite_pk/dto"
)

type CompositePkService interface {
	Save(ctx context.Context, input *dto.CompositePkCreateDto) (*dto.CompositePkDto, error)
	UpdateByID(ctx context.Context, input *dto.CompositePkUpdateDto, pkTenantID int64, pkCode string) (*dto.CompositePkDto, error)
	DeleteByID(ctx context.Context, pkTenantID int64, pkCode string) error
	FindByID(ctx context.Context, pkTenantID int64, pkCode string) (*dto.CompositePkDto, error)
	FindAll(ctx context.Context) ([]dto.CompositePkDto, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.CompositePkDto, pageable.Page, error)
}
