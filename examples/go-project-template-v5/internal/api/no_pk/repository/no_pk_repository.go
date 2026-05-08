package repository

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	dbModel "go-project-template-v5/internal/api/no_pk/entity/postgres"
)

type NoPkRepository interface {
	Save(ctx context.Context, inputEntity *dbModel.NoPk) (*dbModel.NoPk, error)
	FindAll(ctx context.Context) ([]dbModel.NoPk, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.NoPk, pageable.Page, error)
}
