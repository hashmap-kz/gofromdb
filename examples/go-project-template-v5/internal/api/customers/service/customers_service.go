package service

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/customers/dto"
)

type CustomersService interface {
	Save(ctx context.Context, input *dto.CustomersCreateDto) (*dto.CustomersDto, error)
	UpdateByID(ctx context.Context, input *dto.CustomersUpdateDto, pkRecordID int) (*dto.CustomersDto, error)
	DeleteByID(ctx context.Context, pkRecordID int) error
	FindByID(ctx context.Context, pkRecordID int) (*dto.CustomersDto, error)
	FindAll(ctx context.Context) ([]dto.CustomersDto, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.CustomersDto, pageable.Page, error)
}
