package service

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/steps/dto"
)

type StepsService interface {
	Save(ctx context.Context, input *dto.StepsCreateDto) (*dto.StepsDto, error)
	UpdateByID(ctx context.Context, input *dto.StepsUpdateDto, pkRecordID int) (*dto.StepsDto, error)
	DeleteByID(ctx context.Context, pkRecordID int) error
	FindByID(ctx context.Context, pkRecordID int) (*dto.StepsDto, error)
	FindAll(ctx context.Context) ([]dto.StepsDto, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.StepsDto, pageable.Page, error)
}
