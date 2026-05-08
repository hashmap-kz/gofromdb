package service

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/no_pk/dto"
)

type NoPkService interface {
	Save(ctx context.Context, input *dto.NoPkCreateDto) (*dto.NoPkDto, error)
	FindAll(ctx context.Context) ([]dto.NoPkDto, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.NoPkDto, pageable.Page, error)
}
