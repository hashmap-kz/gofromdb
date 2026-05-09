package api

import (
	"context"

	"go-project-template-v5/internal/api/categories"
	"go-project-template-v5/internal/api/order_items"
	"go-project-template-v5/internal/api/orders"
	"go-project-template-v5/internal/api/products"
	"go-project-template-v5/internal/api/users"

	"go-project-template-v5/pkg/storage/postgres"
)

// Init all repos

type Repositories struct {
	CategoriesRepository categories.Repository
	OrderItemsRepository order_items.Repository
	OrdersRepository     orders.Repository
	ProductsRepository   products.Repository
	UsersRepository      users.Repository
}

func NewRepositories(ctx context.Context, db *postgres.Postgres) *Repositories {
	return &Repositories{
		CategoriesRepository: categories.NewRepository(ctx, db),
		OrderItemsRepository: order_items.NewRepository(ctx, db),
		OrdersRepository:     orders.NewRepository(ctx, db),
		ProductsRepository:   products.NewRepository(ctx, db),
		UsersRepository:      users.NewRepository(ctx, db),
	}
}
