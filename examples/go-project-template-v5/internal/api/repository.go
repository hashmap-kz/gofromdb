package api

import (
	"context"

	categoriesRepo "go-project-template-v5/internal/api/categories/repository"
	categoriesImpl "go-project-template-v5/internal/api/categories/repository/impl"
	currenciesRepo "go-project-template-v5/internal/api/currencies/repository"
	currenciesImpl "go-project-template-v5/internal/api/currencies/repository/impl"
	customersRepo "go-project-template-v5/internal/api/customers/repository"
	customersImpl "go-project-template-v5/internal/api/customers/repository/impl"
	departmentsRepo "go-project-template-v5/internal/api/departments/repository"
	departmentsImpl "go-project-template-v5/internal/api/departments/repository/impl"
	jobTitlesRepo "go-project-template-v5/internal/api/job_titles/repository"
	jobTitlesImpl "go-project-template-v5/internal/api/job_titles/repository/impl"
	productPricesRepo "go-project-template-v5/internal/api/product_prices/repository"
	productPricesImpl "go-project-template-v5/internal/api/product_prices/repository/impl"
	productsRepo "go-project-template-v5/internal/api/products/repository"
	productsImpl "go-project-template-v5/internal/api/products/repository/impl"
	purchaseItemsRepo "go-project-template-v5/internal/api/purchase_items/repository"
	purchaseItemsImpl "go-project-template-v5/internal/api/purchase_items/repository/impl"
	purchaseStepsRepo "go-project-template-v5/internal/api/purchase_steps/repository"
	purchaseStepsImpl "go-project-template-v5/internal/api/purchase_steps/repository/impl"
	purchasesRepo "go-project-template-v5/internal/api/purchases/repository"
	purchasesImpl "go-project-template-v5/internal/api/purchases/repository/impl"
	stepsRepo "go-project-template-v5/internal/api/steps/repository"
	stepsImpl "go-project-template-v5/internal/api/steps/repository/impl"
	usersRepo "go-project-template-v5/internal/api/users/repository"
	usersImpl "go-project-template-v5/internal/api/users/repository/impl"

	"go-project-template-v5/pkg/storage/postgres"
)

// Init all repos

type Repositories struct {
	CategoriesRepository    categoriesRepo.CategoriesRepository
	CurrenciesRepository    currenciesRepo.CurrenciesRepository
	CustomersRepository     customersRepo.CustomersRepository
	DepartmentsRepository   departmentsRepo.DepartmentsRepository
	JobTitlesRepository     jobTitlesRepo.JobTitlesRepository
	ProductPricesRepository productPricesRepo.ProductPricesRepository
	ProductsRepository      productsRepo.ProductsRepository
	PurchaseItemsRepository purchaseItemsRepo.PurchaseItemsRepository
	PurchaseStepsRepository purchaseStepsRepo.PurchaseStepsRepository
	PurchasesRepository     purchasesRepo.PurchasesRepository
	StepsRepository         stepsRepo.StepsRepository
	UsersRepository         usersRepo.UsersRepository
}

func NewRepositories(ctx context.Context, db *postgres.Postgres) *Repositories {
	return &Repositories{
		CategoriesRepository:    categoriesImpl.NewCategoriesRepository(ctx, db),
		CurrenciesRepository:    currenciesImpl.NewCurrenciesRepository(ctx, db),
		CustomersRepository:     customersImpl.NewCustomersRepository(ctx, db),
		DepartmentsRepository:   departmentsImpl.NewDepartmentsRepository(ctx, db),
		JobTitlesRepository:     jobTitlesImpl.NewJobTitlesRepository(ctx, db),
		ProductPricesRepository: productPricesImpl.NewProductPricesRepository(ctx, db),
		ProductsRepository:      productsImpl.NewProductsRepository(ctx, db),
		PurchaseItemsRepository: purchaseItemsImpl.NewPurchaseItemsRepository(ctx, db),
		PurchaseStepsRepository: purchaseStepsImpl.NewPurchaseStepsRepository(ctx, db),
		PurchasesRepository:     purchasesImpl.NewPurchasesRepository(ctx, db),
		StepsRepository:         stepsImpl.NewStepsRepository(ctx, db),
		UsersRepository:         usersImpl.NewUsersRepository(ctx, db),
	}
}
