package repository

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	dbModel "go-project-template-v5/internal/api/uuid_pk/entity/postgres"
)

type UUIDPkRepository interface {
	Save(ctx context.Context, inputEntity *dbModel.UUIDPk) (*dbModel.UUIDPk, error)
	UpdateByID(ctx context.Context, inputEntity *dbModel.UUIDPk, pkID string) (*dbModel.UUIDPk, error)
	DeleteByID(ctx context.Context, pkID string) error
	FindByID(ctx context.Context, pkID string) (*dbModel.UUIDPk, error)
	FindAll(ctx context.Context) ([]dbModel.UUIDPk, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.UUIDPk, pageable.Page, error)
}
