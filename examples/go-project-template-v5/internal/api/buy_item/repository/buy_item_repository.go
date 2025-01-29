package repository

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	dbModel "go-project-template-v5/internal/api/buy_item/entity/postgres"
)

type BuyItemRepository interface {
	Save(ctx context.Context, input *dbModel.BuyItem) (*dbModel.BuyItem, error)
	GetAll(ctx context.Context) ([]dbModel.BuyItem, error)
	GetAllPaginated(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.BuyItem, pageable.Page, error)
	Update(ctx context.Context, input *dbModel.BuyItem) (*dbModel.BuyItem, error)
	Delete(ctx context.Context, id int) error
	GetByID(ctx context.Context, id int) (*dbModel.BuyItem, error)
}
