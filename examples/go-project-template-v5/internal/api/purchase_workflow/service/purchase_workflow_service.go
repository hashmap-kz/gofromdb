package service

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/purchase_workflow/dto"
)

type PurchaseWorkflowService interface {
	Save(ctx context.Context, input *dto.PurchaseWorkflowCreateDto) (*dto.PurchaseWorkflowDto, error)
	UpdateByID(ctx context.Context, input *dto.PurchaseWorkflowUpdateDto, pkRecordID int) (*dto.PurchaseWorkflowDto, error)
	DeleteByID(ctx context.Context, pkRecordID int) error
	FindByID(ctx context.Context, pkRecordID int) (*dto.PurchaseWorkflowDto, error)
	FindAll(ctx context.Context) ([]dto.PurchaseWorkflowDto, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.PurchaseWorkflowDto, pageable.Page, error)
}
