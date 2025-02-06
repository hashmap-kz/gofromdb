package repository

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	dbModel "go-project-template-v5/internal/api/currencies/entity/postgres"
)

type CurrenciesRepository interface {
	Save(ctx context.Context, inputEntity *dbModel.Currencies) (*dbModel.Currencies, error)
	UpdateByID(ctx context.Context, inputEntity *dbModel.Currencies, pkRecordID int) (*dbModel.Currencies, error)
	DeleteByID(ctx context.Context, pkRecordID int) error
	FindByID(ctx context.Context, pkRecordID int) (*dbModel.Currencies, error)
	FindAll(ctx context.Context) ([]dbModel.Currencies, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.Currencies, pageable.Page, error)
}
