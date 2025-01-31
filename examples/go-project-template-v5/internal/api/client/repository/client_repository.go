package repository

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	dbModel "go-project-template-v5/internal/api/client/entity/postgres"
)

type ClientRepository interface {
	Save(ctx context.Context, inputEntity *dbModel.Client) (*dbModel.Client, error)
	GetAll(ctx context.Context) ([]dbModel.Client, error)
	GetAllPaginated(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.Client, pageable.Page, error)
	Update(ctx context.Context, entityId int, inputEntity *dbModel.Client) (*dbModel.Client, error)
	Delete(ctx context.Context, entityId int) error
	GetByID(ctx context.Context, entityId int) (*dbModel.Client, error)
}
