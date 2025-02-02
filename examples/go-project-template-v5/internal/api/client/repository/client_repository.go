package repository

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	dbModel "go-project-template-v5/internal/api/client/entity/postgres"
)

type ClientRepository interface {
	Save(ctx context.Context, inputEntity *dbModel.Client) (*dbModel.Client, error)
	UpdateByID(ctx context.Context, pkRecordID int, inputEntity *dbModel.Client) (*dbModel.Client, error)
	DeleteByID(ctx context.Context, pkRecordID int) error
	FindByID(ctx context.Context, pkRecordID int) (*dbModel.Client, error)
	FindAll(ctx context.Context) ([]dbModel.Client, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.Client, pageable.Page, error)
}
