package service

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/buy/dto"
)

type BuyService interface {
	Save(ctx context.Context, input *dto.BuyCreateDto) (*dto.BuyDto, error)
	GetAll(ctx context.Context) ([]dto.BuyDto, error)
	GetAllPaginated(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.BuyDto, pageable.Page, error)
	Update(ctx context.Context, input *dto.BuyUpdateDto) (*dto.BuyDto, error)
	Delete(ctx context.Context, id int) error
	GetByID(ctx context.Context, id int) (*dto.BuyDto, error)
}
