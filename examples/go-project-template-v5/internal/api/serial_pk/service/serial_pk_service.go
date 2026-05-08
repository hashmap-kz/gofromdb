package service

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/serial_pk/dto"
)

type SerialPkService interface {
	Save(ctx context.Context, input *dto.SerialPkCreateDto) (*dto.SerialPkDto, error)
	UpdateByID(ctx context.Context, input *dto.SerialPkUpdateDto, pkID int64) (*dto.SerialPkDto, error)
	DeleteByID(ctx context.Context, pkID int64) error
	FindByID(ctx context.Context, pkID int64) (*dto.SerialPkDto, error)
	FindAll(ctx context.Context) ([]dto.SerialPkDto, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.SerialPkDto, pageable.Page, error)
}
