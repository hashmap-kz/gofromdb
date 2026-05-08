package repository

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	dbModel "go-project-template-v5/internal/api/serial_pk/entity/postgres"
)

type SerialPkRepository interface {
	Save(ctx context.Context, inputEntity *dbModel.SerialPk) (*dbModel.SerialPk, error)
	UpdateByID(ctx context.Context, inputEntity *dbModel.SerialPk, pkID int64) (*dbModel.SerialPk, error)
	DeleteByID(ctx context.Context, pkID int64) error
	FindByID(ctx context.Context, pkID int64) (*dbModel.SerialPk, error)
	FindAll(ctx context.Context) ([]dbModel.SerialPk, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.SerialPk, pageable.Page, error)
}
