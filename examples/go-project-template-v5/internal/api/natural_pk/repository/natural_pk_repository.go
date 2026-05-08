package repository

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	dbModel "go-project-template-v5/internal/api/natural_pk/entity/postgres"
)

type NaturalPkRepository interface {
	Save(ctx context.Context, inputEntity *dbModel.NaturalPk) (*dbModel.NaturalPk, error)
	UpdateByID(ctx context.Context, inputEntity *dbModel.NaturalPk, pkCode string) (*dbModel.NaturalPk, error)
	DeleteByID(ctx context.Context, pkCode string) error
	FindByID(ctx context.Context, pkCode string) (*dbModel.NaturalPk, error)
	FindAll(ctx context.Context) ([]dbModel.NaturalPk, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.NaturalPk, pageable.Page, error)
}
