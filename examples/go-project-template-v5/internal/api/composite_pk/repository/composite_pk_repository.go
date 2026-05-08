package repository

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	dbModel "go-project-template-v5/internal/api/composite_pk/entity/postgres"
)

type CompositePkRepository interface {
	Save(ctx context.Context, inputEntity *dbModel.CompositePk) (*dbModel.CompositePk, error)
	UpdateByID(ctx context.Context, inputEntity *dbModel.CompositePk, pkTenantID int64, pkCode string) (*dbModel.CompositePk, error)
	DeleteByID(ctx context.Context, pkTenantID int64, pkCode string) error
	FindByID(ctx context.Context, pkTenantID int64, pkCode string) (*dbModel.CompositePk, error)
	FindAll(ctx context.Context) ([]dbModel.CompositePk, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.CompositePk, pageable.Page, error)
}
