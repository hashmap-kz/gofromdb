package api

import (
	"context"

	categoriesServ "go-project-template-v5/internal/api/categories/service"
	categoriesImpl "go-project-template-v5/internal/api/categories/service/impl"
	currenciesServ "go-project-template-v5/internal/api/currencies/service"
	currenciesImpl "go-project-template-v5/internal/api/currencies/service/impl"
	customersServ "go-project-template-v5/internal/api/customers/service"
	customersImpl "go-project-template-v5/internal/api/customers/service/impl"
	productPricesServ "go-project-template-v5/internal/api/product_prices/service"
	productPricesImpl "go-project-template-v5/internal/api/product_prices/service/impl"
	productsServ "go-project-template-v5/internal/api/products/service"
	productsImpl "go-project-template-v5/internal/api/products/service/impl"
	purchaseItemsServ "go-project-template-v5/internal/api/purchase_items/service"
	purchaseItemsImpl "go-project-template-v5/internal/api/purchase_items/service/impl"
	purchaseStepsServ "go-project-template-v5/internal/api/purchase_steps/service"
	purchaseStepsImpl "go-project-template-v5/internal/api/purchase_steps/service/impl"
	purchasesServ "go-project-template-v5/internal/api/purchases/service"
	purchasesImpl "go-project-template-v5/internal/api/purchases/service/impl"
	stepsServ "go-project-template-v5/internal/api/steps/service"
	stepsImpl "go-project-template-v5/internal/api/steps/service/impl"
	usersServ "go-project-template-v5/internal/api/users/service"
	usersImpl "go-project-template-v5/internal/api/users/service/impl"
)

// Init all services

type Services struct {
	CategoriesService    categoriesServ.CategoriesService
	CurrenciesService    currenciesServ.CurrenciesService
	CustomersService     customersServ.CustomersService
	ProductPricesService productPricesServ.ProductPricesService
	ProductsService      productsServ.ProductsService
	PurchaseItemsService purchaseItemsServ.PurchaseItemsService
	PurchaseStepsService purchaseStepsServ.PurchaseStepsService
	PurchasesService     purchasesServ.PurchasesService
	StepsService         stepsServ.StepsService
	UsersService         usersServ.UsersService
}

type Deps struct {
	// TODO: other deps here
	Repos *Repositories
}

func NewServices(ctx context.Context, deps Deps) *Services {
	return &Services{
		CategoriesService:    categoriesImpl.NewCategoriesService(ctx, deps.Repos.CategoriesRepository),
		CurrenciesService:    currenciesImpl.NewCurrenciesService(ctx, deps.Repos.CurrenciesRepository),
		CustomersService:     customersImpl.NewCustomersService(ctx, deps.Repos.CustomersRepository),
		ProductPricesService: productPricesImpl.NewProductPricesService(ctx, deps.Repos.ProductPricesRepository),
		ProductsService:      productsImpl.NewProductsService(ctx, deps.Repos.ProductsRepository),
		PurchaseItemsService: purchaseItemsImpl.NewPurchaseItemsService(ctx, deps.Repos.PurchaseItemsRepository),
		PurchaseStepsService: purchaseStepsImpl.NewPurchaseStepsService(ctx, deps.Repos.PurchaseStepsRepository),
		PurchasesService:     purchasesImpl.NewPurchasesService(ctx, deps.Repos.PurchasesRepository),
		StepsService:         stepsImpl.NewStepsService(ctx, deps.Repos.StepsRepository),
		UsersService:         usersImpl.NewUsersService(ctx, deps.Repos.UsersRepository),
	}
}
