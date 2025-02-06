package repository

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	dbModel "go-project-template-v5/internal/api/steps/entity/postgres"
)

type StepsRepository interface {
	Save(ctx context.Context, inputEntity *dbModel.Steps) (*dbModel.Steps, error)
	UpdateByID(ctx context.Context, inputEntity *dbModel.Steps, pkRecordID int) (*dbModel.Steps, error)
	DeleteByID(ctx context.Context, pkRecordID int) error
	FindByID(ctx context.Context, pkRecordID int) (*dbModel.Steps, error)
	FindAll(ctx context.Context) ([]dbModel.Steps, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.Steps, pageable.Page, error)
}
