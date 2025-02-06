package api

import (
	"context"

	categoriesRepo "go-project-template-v5/internal/api/categories/repository"
	categoriesImpl "go-project-template-v5/internal/api/categories/repository/impl"
	clientsRepo "go-project-template-v5/internal/api/clients/repository"
	clientsImpl "go-project-template-v5/internal/api/clients/repository/impl"
	customerOrderItemsRepo "go-project-template-v5/internal/api/customer_order_items/repository"
	customerOrderItemsImpl "go-project-template-v5/internal/api/customer_order_items/repository/impl"
	customerOrdersRepo "go-project-template-v5/internal/api/customer_orders/repository"
	customerOrdersImpl "go-project-template-v5/internal/api/customer_orders/repository/impl"
	productsRepo "go-project-template-v5/internal/api/products/repository"
	productsImpl "go-project-template-v5/internal/api/products/repository/impl"

	"go-project-template-v5/pkg/storage/postgres"
)

// Init all repos

type Repositories struct {
	CategoriesRepository         categoriesRepo.CategoriesRepository
	ClientsRepository            clientsRepo.ClientsRepository
	CustomerOrderItemsRepository customerOrderItemsRepo.CustomerOrderItemsRepository
	CustomerOrdersRepository     customerOrdersRepo.CustomerOrdersRepository
	ProductsRepository           productsRepo.ProductsRepository
}

func NewRepositories(ctx context.Context, db *postgres.Postgres) *Repositories {
	return &Repositories{
		CategoriesRepository:         categoriesImpl.NewCategoriesRepository(ctx, db),
		ClientsRepository:            clientsImpl.NewClientsRepository(ctx, db),
		CustomerOrderItemsRepository: customerOrderItemsImpl.NewCustomerOrderItemsRepository(ctx, db),
		CustomerOrdersRepository:     customerOrdersImpl.NewCustomerOrdersRepository(ctx, db),
		ProductsRepository:           productsImpl.NewProductsRepository(ctx, db),
	}
}
