package repository

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	dbModel "go-project-template-v5/internal/api/customers/entity/postgres"
)

type CustomersRepository interface {
	Save(ctx context.Context, inputEntity *dbModel.Customers) (*dbModel.Customers, error)
	UpdateByID(ctx context.Context, inputEntity *dbModel.Customers, pkRecordID int) (*dbModel.Customers, error)
	DeleteByID(ctx context.Context, pkRecordID int) error
	FindByID(ctx context.Context, pkRecordID int) (*dbModel.Customers, error)
	FindAll(ctx context.Context) ([]dbModel.Customers, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.Customers, pageable.Page, error)
}
