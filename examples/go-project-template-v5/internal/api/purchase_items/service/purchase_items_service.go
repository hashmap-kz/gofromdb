package service

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/purchase_items/dto"
)

type PurchaseItemsService interface {
	Save(ctx context.Context, input *dto.PurchaseItemsCreateDto) (*dto.PurchaseItemsDto, error)
	UpdateByID(ctx context.Context, input *dto.PurchaseItemsUpdateDto, pkRecordID int) (*dto.PurchaseItemsDto, error)
	DeleteByID(ctx context.Context, pkRecordID int) error
	FindByID(ctx context.Context, pkRecordID int) (*dto.PurchaseItemsDto, error)
	FindAll(ctx context.Context) ([]dto.PurchaseItemsDto, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.PurchaseItemsDto, pageable.Page, error)
}
