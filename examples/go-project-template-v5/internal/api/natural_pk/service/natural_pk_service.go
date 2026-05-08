package service

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/natural_pk/dto"
)

type NaturalPkService interface {
	Save(ctx context.Context, input *dto.NaturalPkCreateDto) (*dto.NaturalPkDto, error)
	UpdateByID(ctx context.Context, input *dto.NaturalPkUpdateDto, pkCode string) (*dto.NaturalPkDto, error)
	DeleteByID(ctx context.Context, pkCode string) error
	FindByID(ctx context.Context, pkCode string) (*dto.NaturalPkDto, error)
	FindAll(ctx context.Context) ([]dto.NaturalPkDto, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.NaturalPkDto, pageable.Page, error)
}
