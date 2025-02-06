package api

import (
	"context"

	categoriesServ "go-project-template-v5/internal/api/categories/service"
	categoriesImpl "go-project-template-v5/internal/api/categories/service/impl"
	clientsServ "go-project-template-v5/internal/api/clients/service"
	clientsImpl "go-project-template-v5/internal/api/clients/service/impl"
	customerOrderItemsServ "go-project-template-v5/internal/api/customer_order_items/service"
	customerOrderItemsImpl "go-project-template-v5/internal/api/customer_order_items/service/impl"
	customerOrdersServ "go-project-template-v5/internal/api/customer_orders/service"
	customerOrdersImpl "go-project-template-v5/internal/api/customer_orders/service/impl"
	productsServ "go-project-template-v5/internal/api/products/service"
	productsImpl "go-project-template-v5/internal/api/products/service/impl"
)

// Init all services

type Services struct {
	CategoriesService         categoriesServ.CategoriesService
	ClientsService            clientsServ.ClientsService
	CustomerOrderItemsService customerOrderItemsServ.CustomerOrderItemsService
	CustomerOrdersService     customerOrdersServ.CustomerOrdersService
	ProductsService           productsServ.ProductsService
}

type Deps struct {
	// TODO: other deps here
	Repos *Repositories
}

func NewServices(ctx context.Context, deps Deps) *Services {
	return &Services{
		CategoriesService:         categoriesImpl.NewCategoriesService(ctx, deps.Repos.CategoriesRepository),
		ClientsService:            clientsImpl.NewClientsService(ctx, deps.Repos.ClientsRepository),
		CustomerOrderItemsService: customerOrderItemsImpl.NewCustomerOrderItemsService(ctx, deps.Repos.CustomerOrderItemsRepository),
		CustomerOrdersService:     customerOrdersImpl.NewCustomerOrdersService(ctx, deps.Repos.CustomerOrdersRepository),
		ProductsService:           productsImpl.NewProductsService(ctx, deps.Repos.ProductsRepository),
	}
}
