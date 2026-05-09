package api

import (
	"context"
	"go-project-template-v5/internal/api/categories"
	"go-project-template-v5/internal/api/users"
	"go-project-template-v5/pkg/storage/postgres"
)

// Init all repos

type Repositories struct {
	CategoriesRepository categories.Repository
	UsersRepository      users.Repository
}

func NewRepositories(ctx context.Context, db *postgres.Postgres) *Repositories {
	return &Repositories{
		CategoriesRepository: categories.NewRepository(ctx, db),
		UsersRepository:      users.NewRepository(ctx, db),
	}
}
