package service

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/purchase_steps/dto"
)

type PurchaseStepsService interface {
	Save(ctx context.Context, input *dto.PurchaseStepsCreateDto) (*dto.PurchaseStepsDto, error)
	UpdateByID(ctx context.Context, input *dto.PurchaseStepsUpdateDto, pkRecordID int) (*dto.PurchaseStepsDto, error)
	DeleteByID(ctx context.Context, pkRecordID int) error
	FindByID(ctx context.Context, pkRecordID int) (*dto.PurchaseStepsDto, error)
	FindAll(ctx context.Context) ([]dto.PurchaseStepsDto, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.PurchaseStepsDto, pageable.Page, error)
}
