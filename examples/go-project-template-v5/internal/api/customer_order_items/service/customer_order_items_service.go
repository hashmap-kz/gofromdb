package service

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/customer_order_items/dto"
)

type CustomerOrderItemsService interface {
	Save(ctx context.Context, input *dto.CustomerOrderItemsCreateDto) (*dto.CustomerOrderItemsDto, error)
	UpdateByID(ctx context.Context, input *dto.CustomerOrderItemsUpdateDto, pkRecordID int) (*dto.CustomerOrderItemsDto, error)
	DeleteByID(ctx context.Context, pkRecordID int) error
	FindByID(ctx context.Context, pkRecordID int) (*dto.CustomerOrderItemsDto, error)
	FindAll(ctx context.Context) ([]dto.CustomerOrderItemsDto, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.CustomerOrderItemsDto, pageable.Page, error)
}
