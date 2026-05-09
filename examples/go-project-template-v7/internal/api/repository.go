package api

import (
	"context"
	"go-project-template-v7/internal/api/categories"
	"go-project-template-v7/internal/api/order_items"
	"go-project-template-v7/internal/api/orders"
	"go-project-template-v7/internal/api/products"
	"go-project-template-v7/internal/api/users"
	"go-project-template-v7/pkg/storage/postgres"
)

// Init all repos

type Repositories struct {
	Categories categories.Repository
	OrderItems order_items.Repository
	Orders     orders.Repository
	Products   products.Repository
	Users      users.Repository
}

func NewRepositories(ctx context.Context, db *postgres.Postgres) *Repositories {
	return &Repositories{
		Categories: categories.NewRepository(ctx, db),
		OrderItems: order_items.NewRepository(ctx, db),
		Orders:     orders.NewRepository(ctx, db),
		Products:   products.NewRepository(ctx, db),
		Users:      users.NewRepository(ctx, db),
	}
}
