package api

import (
	"context"

	currenciesRepo "go-project-template-v5/internal/api/currencies/repository"
	currenciesImpl "go-project-template-v5/internal/api/currencies/repository/impl"
	customersRepo "go-project-template-v5/internal/api/customers/repository"
	customersImpl "go-project-template-v5/internal/api/customers/repository/impl"
	productCategoriesRepo "go-project-template-v5/internal/api/product_categories/repository"
	productCategoriesImpl "go-project-template-v5/internal/api/product_categories/repository/impl"
	productPricesRepo "go-project-template-v5/internal/api/product_prices/repository"
	productPricesImpl "go-project-template-v5/internal/api/product_prices/repository/impl"
	productsRepo "go-project-template-v5/internal/api/products/repository"
	productsImpl "go-project-template-v5/internal/api/products/repository/impl"
	purchaseItemsRepo "go-project-template-v5/internal/api/purchase_items/repository"
	purchaseItemsImpl "go-project-template-v5/internal/api/purchase_items/repository/impl"
	purchaseStepsRepo "go-project-template-v5/internal/api/purchase_steps/repository"
	purchaseStepsImpl "go-project-template-v5/internal/api/purchase_steps/repository/impl"
	purchaseWorkflowRepo "go-project-template-v5/internal/api/purchase_workflow/repository"
	purchaseWorkflowImpl "go-project-template-v5/internal/api/purchase_workflow/repository/impl"
	purchasesRepo "go-project-template-v5/internal/api/purchases/repository"
	purchasesImpl "go-project-template-v5/internal/api/purchases/repository/impl"
	usersRepo "go-project-template-v5/internal/api/users/repository"
	usersImpl "go-project-template-v5/internal/api/users/repository/impl"

	"go-project-template-v5/pkg/storage/postgres"
)

// Init all repos

type Repositories struct {
	CurrenciesRepository        currenciesRepo.CurrenciesRepository
	CustomersRepository         customersRepo.CustomersRepository
	ProductCategoriesRepository productCategoriesRepo.ProductCategoriesRepository
	ProductPricesRepository     productPricesRepo.ProductPricesRepository
	ProductsRepository          productsRepo.ProductsRepository
	PurchaseItemsRepository     purchaseItemsRepo.PurchaseItemsRepository
	PurchaseStepsRepository     purchaseStepsRepo.PurchaseStepsRepository
	PurchaseWorkflowRepository  purchaseWorkflowRepo.PurchaseWorkflowRepository
	PurchasesRepository         purchasesRepo.PurchasesRepository
	UsersRepository             usersRepo.UsersRepository
}

func NewRepositories(ctx context.Context, db *postgres.Postgres) *Repositories {
	return &Repositories{
		CurrenciesRepository:        currenciesImpl.NewCurrenciesRepository(ctx, db),
		CustomersRepository:         customersImpl.NewCustomersRepository(ctx, db),
		ProductCategoriesRepository: productCategoriesImpl.NewProductCategoriesRepository(ctx, db),
		ProductPricesRepository:     productPricesImpl.NewProductPricesRepository(ctx, db),
		ProductsRepository:          productsImpl.NewProductsRepository(ctx, db),
		PurchaseItemsRepository:     purchaseItemsImpl.NewPurchaseItemsRepository(ctx, db),
		PurchaseStepsRepository:     purchaseStepsImpl.NewPurchaseStepsRepository(ctx, db),
		PurchaseWorkflowRepository:  purchaseWorkflowImpl.NewPurchaseWorkflowRepository(ctx, db),
		PurchasesRepository:         purchasesImpl.NewPurchasesRepository(ctx, db),
		UsersRepository:             usersImpl.NewUsersRepository(ctx, db),
	}
}
