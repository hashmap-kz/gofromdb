package api

import (
	"context"
	"go-project-template-v5/internal/api/categories"
	"go-project-template-v5/internal/api/users"
)

// Init all services

type Services struct {
	Categories categories.Service
	Users      users.Service
}

type Deps struct {
	// TODO: other deps here
	Repos *Repositories
}

func NewServices(ctx context.Context, deps Deps) *Services {
	return &Services{
		Categories: categories.NewService(ctx, deps.Repos.CategoriesRepository),
		Users:      users.NewService(ctx, deps.Repos.UsersRepository),
	}
}
