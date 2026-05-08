package api

import (
	"context"

	categoriesRepo "go-project-template-v5/internal/api/categories/repository"
	categoriesImpl "go-project-template-v5/internal/api/categories/repository/impl"
	compositePkRepo "go-project-template-v5/internal/api/composite_pk/repository"
	compositePkImpl "go-project-template-v5/internal/api/composite_pk/repository/impl"
	naturalPkRepo "go-project-template-v5/internal/api/natural_pk/repository"
	naturalPkImpl "go-project-template-v5/internal/api/natural_pk/repository/impl"
	noPkRepo "go-project-template-v5/internal/api/no_pk/repository"
	noPkImpl "go-project-template-v5/internal/api/no_pk/repository/impl"
	nullableTypesRepo "go-project-template-v5/internal/api/nullable_types/repository"
	nullableTypesImpl "go-project-template-v5/internal/api/nullable_types/repository/impl"
	orderItemsRepo "go-project-template-v5/internal/api/order_items/repository"
	orderItemsImpl "go-project-template-v5/internal/api/order_items/repository/impl"
	ordersRepo "go-project-template-v5/internal/api/orders/repository"
	ordersImpl "go-project-template-v5/internal/api/orders/repository/impl"
	productsRepo "go-project-template-v5/internal/api/products/repository"
	productsImpl "go-project-template-v5/internal/api/products/repository/impl"
	serialPkRepo "go-project-template-v5/internal/api/serial_pk/repository"
	serialPkImpl "go-project-template-v5/internal/api/serial_pk/repository/impl"
	usersRepo "go-project-template-v5/internal/api/users/repository"
	usersImpl "go-project-template-v5/internal/api/users/repository/impl"
	uUIDPkRepo "go-project-template-v5/internal/api/uuid_pk/repository"
	uUIDPkImpl "go-project-template-v5/internal/api/uuid_pk/repository/impl"

	"go-project-template-v5/pkg/storage/postgres"
)

// Init all repos

type Repositories struct {
	CategoriesRepository    categoriesRepo.CategoriesRepository
	CompositePkRepository   compositePkRepo.CompositePkRepository
	NaturalPkRepository     naturalPkRepo.NaturalPkRepository
	NoPkRepository          noPkRepo.NoPkRepository
	NullableTypesRepository nullableTypesRepo.NullableTypesRepository
	OrderItemsRepository    orderItemsRepo.OrderItemsRepository
	OrdersRepository        ordersRepo.OrdersRepository
	ProductsRepository      productsRepo.ProductsRepository
	SerialPkRepository      serialPkRepo.SerialPkRepository
	UUIDPkRepository        uUIDPkRepo.UUIDPkRepository
	UsersRepository         usersRepo.UsersRepository
}

func NewRepositories(ctx context.Context, db *postgres.Postgres) *Repositories {
	return &Repositories{
		CategoriesRepository:    categoriesImpl.NewCategoriesRepository(ctx, db),
		CompositePkRepository:   compositePkImpl.NewCompositePkRepository(ctx, db),
		NaturalPkRepository:     naturalPkImpl.NewNaturalPkRepository(ctx, db),
		NoPkRepository:          noPkImpl.NewNoPkRepository(ctx, db),
		NullableTypesRepository: nullableTypesImpl.NewNullableTypesRepository(ctx, db),
		OrderItemsRepository:    orderItemsImpl.NewOrderItemsRepository(ctx, db),
		OrdersRepository:        ordersImpl.NewOrdersRepository(ctx, db),
		ProductsRepository:      productsImpl.NewProductsRepository(ctx, db),
		SerialPkRepository:      serialPkImpl.NewSerialPkRepository(ctx, db),
		UUIDPkRepository:        uUIDPkImpl.NewUUIDPkRepository(ctx, db),
		UsersRepository:         usersImpl.NewUsersRepository(ctx, db),
	}
}
