package repository

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	dbModel "go-project-template-v5/internal/api/clients/entity/postgres"
)

type ClientsRepository interface {
	Save(ctx context.Context, inputEntity *dbModel.Clients) (*dbModel.Clients, error)
	UpdateByID(ctx context.Context, inputEntity *dbModel.Clients, pkRecordID int) (*dbModel.Clients, error)
	DeleteByID(ctx context.Context, pkRecordID int) error
	FindByID(ctx context.Context, pkRecordID int) (*dbModel.Clients, error)
	FindAll(ctx context.Context) ([]dbModel.Clients, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.Clients, pageable.Page, error)
}
