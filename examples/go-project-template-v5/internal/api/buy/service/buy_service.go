package service

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/buy/dto"
)

type BuyService interface {
	Save(ctx context.Context, input *dto.BuyCreateDto) (*dto.BuyDto, error)
	UpdateByID(ctx context.Context, pkRecordID int, input *dto.BuyUpdateDto) (*dto.BuyDto, error)
	DeleteByID(ctx context.Context, pkRecordID int) error
	FindByID(ctx context.Context, pkRecordID int) (*dto.BuyDto, error)
	FindAll(ctx context.Context) ([]dto.BuyDto, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.BuyDto, pageable.Page, error)
}
