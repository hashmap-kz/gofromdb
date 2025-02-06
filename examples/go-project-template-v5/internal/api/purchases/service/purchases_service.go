package service

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/purchases/dto"
)

type PurchasesService interface {
	Save(ctx context.Context, input *dto.PurchasesCreateDto) (*dto.PurchasesDto, error)
	UpdateByID(ctx context.Context, input *dto.PurchasesUpdateDto, pkRecordID int) (*dto.PurchasesDto, error)
	DeleteByID(ctx context.Context, pkRecordID int) error
	FindByID(ctx context.Context, pkRecordID int) (*dto.PurchasesDto, error)
	FindAll(ctx context.Context) ([]dto.PurchasesDto, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.PurchasesDto, pageable.Page, error)
}
