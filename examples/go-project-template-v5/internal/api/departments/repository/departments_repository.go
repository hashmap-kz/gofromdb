package repository

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	dbModel "go-project-template-v5/internal/api/departments/entity/postgres"
)

type DepartmentsRepository interface {
	Save(ctx context.Context, inputEntity *dbModel.Departments) (*dbModel.Departments, error)
	UpdateByID(ctx context.Context, inputEntity *dbModel.Departments, pkRecordID int) (*dbModel.Departments, error)
	DeleteByID(ctx context.Context, pkRecordID int) error
	FindByID(ctx context.Context, pkRecordID int) (*dbModel.Departments, error)
	FindAll(ctx context.Context) ([]dbModel.Departments, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.Departments, pageable.Page, error)
}
