package repository

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	dbModel "go-project-template-v5/internal/api/buy_item/entity/postgres"
)

type BuyItemRepository interface {
	Save(ctx context.Context, inputEntity *dbModel.BuyItem) (*dbModel.BuyItem, error)
	UpdateByID(ctx context.Context, inputEntity *dbModel.BuyItem, pkRecordID int) (*dbModel.BuyItem, error)
	DeleteByID(ctx context.Context, pkRecordID int) error
	FindByID(ctx context.Context, pkRecordID int) (*dbModel.BuyItem, error)
	FindAll(ctx context.Context) ([]dbModel.BuyItem, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.BuyItem, pageable.Page, error)
}
