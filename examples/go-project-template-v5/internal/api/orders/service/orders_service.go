package service

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/orders/dto"
)

type OrdersService interface {
	Save(ctx context.Context, input *dto.OrdersCreateDto) (*dto.OrdersDto, error)
	UpdateByID(ctx context.Context, input *dto.OrdersUpdateDto, pkRecordID int) (*dto.OrdersDto, error)
	DeleteByID(ctx context.Context, pkRecordID int) error
	FindByID(ctx context.Context, pkRecordID int) (*dto.OrdersDto, error)
	FindAll(ctx context.Context) ([]dto.OrdersDto, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.OrdersDto, pageable.Page, error)
}
