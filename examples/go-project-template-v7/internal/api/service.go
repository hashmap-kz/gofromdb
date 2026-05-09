package api

import (
	"context"
	"go-project-template-v7/internal/api/categories"
	"go-project-template-v7/internal/api/order_items"
	"go-project-template-v7/internal/api/orders"
	"go-project-template-v7/internal/api/products"
	"go-project-template-v7/internal/api/users"
)

// Init all services

type Services struct {
	Categories categories.Service
	OrderItems order_items.Service
	Orders     orders.Service
	Products   products.Service
	Users      users.Service
}

type Deps struct {
	// TODO: other deps here
	Repos *Repositories
}

func NewServices(ctx context.Context, deps Deps) *Services {
	return &Services{
		Categories: categories.NewService(ctx, deps.Repos.Categories),
		OrderItems: order_items.NewService(ctx, deps.Repos.OrderItems),
		Orders:     orders.NewService(ctx, deps.Repos.Orders),
		Products:   products.NewService(ctx, deps.Repos.Products),
		Users:      users.NewService(ctx, deps.Repos.Users),
	}
}
