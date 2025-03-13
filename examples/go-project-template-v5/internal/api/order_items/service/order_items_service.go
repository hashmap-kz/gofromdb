package service

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/order_items/dto"
)

type OrderItemsService interface {
	Save(ctx context.Context, input *dto.OrderItemsCreateDto) (*dto.OrderItemsDto, error)
	UpdateByID(ctx context.Context, input *dto.OrderItemsUpdateDto, pkRecordID int) (*dto.OrderItemsDto, error)
	DeleteByID(ctx context.Context, pkRecordID int) error
	FindByID(ctx context.Context, pkRecordID int) (*dto.OrderItemsDto, error)
	FindAll(ctx context.Context) ([]dto.OrderItemsDto, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.OrderItemsDto, pageable.Page, error)
}
