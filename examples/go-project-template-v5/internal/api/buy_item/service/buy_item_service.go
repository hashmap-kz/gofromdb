package service

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/buy_item/dto"
)

type BuyItemService interface {
	Save(ctx context.Context, input *dto.BuyItemCreateDto) (*dto.BuyItemDto, error)
	UpdateByID(ctx context.Context, input *dto.BuyItemUpdateDto, pkRecordID int) (*dto.BuyItemDto, error)
	DeleteByID(ctx context.Context, pkRecordID int) error
	FindByID(ctx context.Context, pkRecordID int) (*dto.BuyItemDto, error)
	FindAll(ctx context.Context) ([]dto.BuyItemDto, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.BuyItemDto, pageable.Page, error)
}
