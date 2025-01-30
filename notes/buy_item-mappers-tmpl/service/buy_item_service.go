package service

import (
	"context"
	"go-project-template-v5/internal/api/buy_item/dto"
	"go-project-template-v5/pkg/pageable"
)

type BuyItemService interface {
	Save(ctx context.Context, input *dto.BuyItemCreateDto) (*dto.BuyItemDto, error)
	GetAll(ctx context.Context) ([]dto.BuyItemDto, error)
	GetAllPaginated(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.BuyItemDto, pageable.Page, error)
	Update(ctx context.Context, entityId int, input *dto.BuyItemUpdateDto) (*dto.BuyItemDto, error)
	Delete(ctx context.Context, id int) error
	GetByID(ctx context.Context, id int) (*dto.BuyItemDto, error)
}
