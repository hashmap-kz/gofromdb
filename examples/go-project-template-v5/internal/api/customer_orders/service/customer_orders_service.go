package service

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/customer_orders/dto"
)

type CustomerOrdersService interface {
	Save(ctx context.Context, input *dto.CustomerOrdersCreateDto) (*dto.CustomerOrdersDto, error)
	UpdateByID(ctx context.Context, input *dto.CustomerOrdersUpdateDto, pkRecordID int) (*dto.CustomerOrdersDto, error)
	DeleteByID(ctx context.Context, pkRecordID int) error
	FindByID(ctx context.Context, pkRecordID int) (*dto.CustomerOrdersDto, error)
	FindAll(ctx context.Context) ([]dto.CustomerOrdersDto, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.CustomerOrdersDto, pageable.Page, error)
}
