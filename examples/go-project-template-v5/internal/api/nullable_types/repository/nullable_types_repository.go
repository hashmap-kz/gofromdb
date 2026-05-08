package repository

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	dbModel "go-project-template-v5/internal/api/nullable_types/entity/postgres"
)

type NullableTypesRepository interface {
	Save(ctx context.Context, inputEntity *dbModel.NullableTypes) (*dbModel.NullableTypes, error)
	UpdateByID(ctx context.Context, inputEntity *dbModel.NullableTypes, pkID int64) (*dbModel.NullableTypes, error)
	DeleteByID(ctx context.Context, pkID int64) error
	FindByID(ctx context.Context, pkID int64) (*dbModel.NullableTypes, error)
	FindAll(ctx context.Context) ([]dbModel.NullableTypes, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.NullableTypes, pageable.Page, error)
}
