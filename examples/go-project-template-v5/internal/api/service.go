package api

import (
	"context"

	categoriesServ "go-project-template-v5/internal/api/categories/service"
	categoriesImpl "go-project-template-v5/internal/api/categories/service/impl"
	compositePkServ "go-project-template-v5/internal/api/composite_pk/service"
	compositePkImpl "go-project-template-v5/internal/api/composite_pk/service/impl"
	naturalPkServ "go-project-template-v5/internal/api/natural_pk/service"
	naturalPkImpl "go-project-template-v5/internal/api/natural_pk/service/impl"
	noPkServ "go-project-template-v5/internal/api/no_pk/service"
	noPkImpl "go-project-template-v5/internal/api/no_pk/service/impl"
	nullableTypesServ "go-project-template-v5/internal/api/nullable_types/service"
	nullableTypesImpl "go-project-template-v5/internal/api/nullable_types/service/impl"
	orderItemsServ "go-project-template-v5/internal/api/order_items/service"
	orderItemsImpl "go-project-template-v5/internal/api/order_items/service/impl"
	ordersServ "go-project-template-v5/internal/api/orders/service"
	ordersImpl "go-project-template-v5/internal/api/orders/service/impl"
	productsServ "go-project-template-v5/internal/api/products/service"
	productsImpl "go-project-template-v5/internal/api/products/service/impl"
	serialPkServ "go-project-template-v5/internal/api/serial_pk/service"
	serialPkImpl "go-project-template-v5/internal/api/serial_pk/service/impl"
	usersServ "go-project-template-v5/internal/api/users/service"
	usersImpl "go-project-template-v5/internal/api/users/service/impl"
	uUIDPkServ "go-project-template-v5/internal/api/uuid_pk/service"
	uUIDPkImpl "go-project-template-v5/internal/api/uuid_pk/service/impl"
)

// Init all services

type Services struct {
	CategoriesService    categoriesServ.CategoriesService
	CompositePkService   compositePkServ.CompositePkService
	NaturalPkService     naturalPkServ.NaturalPkService
	NoPkService          noPkServ.NoPkService
	NullableTypesService nullableTypesServ.NullableTypesService
	OrderItemsService    orderItemsServ.OrderItemsService
	OrdersService        ordersServ.OrdersService
	ProductsService      productsServ.ProductsService
	SerialPkService      serialPkServ.SerialPkService
	UUIDPkService        uUIDPkServ.UUIDPkService
	UsersService         usersServ.UsersService
}

type Deps struct {
	// TODO: other deps here
	Repos *Repositories
}

func NewServices(ctx context.Context, deps Deps) *Services {
	return &Services{
		CategoriesService:    categoriesImpl.NewCategoriesService(ctx, deps.Repos.CategoriesRepository),
		CompositePkService:   compositePkImpl.NewCompositePkService(ctx, deps.Repos.CompositePkRepository),
		NaturalPkService:     naturalPkImpl.NewNaturalPkService(ctx, deps.Repos.NaturalPkRepository),
		NoPkService:          noPkImpl.NewNoPkService(ctx, deps.Repos.NoPkRepository),
		NullableTypesService: nullableTypesImpl.NewNullableTypesService(ctx, deps.Repos.NullableTypesRepository),
		OrderItemsService:    orderItemsImpl.NewOrderItemsService(ctx, deps.Repos.OrderItemsRepository),
		OrdersService:        ordersImpl.NewOrdersService(ctx, deps.Repos.OrdersRepository),
		ProductsService:      productsImpl.NewProductsService(ctx, deps.Repos.ProductsRepository),
		SerialPkService:      serialPkImpl.NewSerialPkService(ctx, deps.Repos.SerialPkRepository),
		UUIDPkService:        uUIDPkImpl.NewUUIDPkService(ctx, deps.Repos.UUIDPkRepository),
		UsersService:         usersImpl.NewUsersService(ctx, deps.Repos.UsersRepository),
	}
}
