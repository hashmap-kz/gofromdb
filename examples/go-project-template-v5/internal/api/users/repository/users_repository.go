package repository

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	dbModel "go-project-template-v5/internal/api/users/entity/postgres"
)

type UsersRepository interface {
	Save(ctx context.Context, inputEntity *dbModel.Users) (*dbModel.Users, error)
	UpdateByID(ctx context.Context, inputEntity *dbModel.Users, pkRecordID int) (*dbModel.Users, error)
	DeleteByID(ctx context.Context, pkRecordID int) error
	FindByID(ctx context.Context, pkRecordID int) (*dbModel.Users, error)
	FindAll(ctx context.Context) ([]dbModel.Users, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.Users, pageable.Page, error)
}
