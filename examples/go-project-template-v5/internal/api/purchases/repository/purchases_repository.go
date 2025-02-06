package repository

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	dbModel "go-project-template-v5/internal/api/purchases/entity/postgres"
)

type PurchasesRepository interface {
	Save(ctx context.Context, inputEntity *dbModel.Purchases) (*dbModel.Purchases, error)
	UpdateByID(ctx context.Context, inputEntity *dbModel.Purchases, pkRecordID int) (*dbModel.Purchases, error)
	DeleteByID(ctx context.Context, pkRecordID int) error
	FindByID(ctx context.Context, pkRecordID int) (*dbModel.Purchases, error)
	FindAll(ctx context.Context) ([]dbModel.Purchases, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.Purchases, pageable.Page, error)
}
