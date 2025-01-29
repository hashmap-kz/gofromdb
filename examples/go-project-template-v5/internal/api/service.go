package api

import (
	"context"

	buyServ "go-project-template-v5/internal/api/buy/service"
	buyImpl "go-project-template-v5/internal/api/buy/service/impl"
	buyitemServ "go-project-template-v5/internal/api/buy_item/service"
	buyitemImpl "go-project-template-v5/internal/api/buy_item/service/impl"
	categoryServ "go-project-template-v5/internal/api/category/service"
	categoryImpl "go-project-template-v5/internal/api/category/service/impl"
	clientServ "go-project-template-v5/internal/api/client/service"
	clientImpl "go-project-template-v5/internal/api/client/service/impl"
	productServ "go-project-template-v5/internal/api/product/service"
	productImpl "go-project-template-v5/internal/api/product/service/impl"
)

// Init all services

type Services struct {
	ProductService  productServ.ProductService
	BuyService      buyServ.BuyService
	BuyItemService  buyitemServ.BuyItemService
	CategoryService categoryServ.CategoryService
	ClientService   clientServ.ClientService
}

type Deps struct {
	// TODO: other deps here
	Repos *Repositories
}

func NewServices(ctx context.Context, deps Deps) *Services {
	return &Services{
		ProductService:  productImpl.NewProductService(ctx, deps.Repos.ProductRepository),
		BuyService:      buyImpl.NewBuyService(ctx, deps.Repos.BuyRepository),
		BuyItemService:  buyitemImpl.NewBuyItemService(ctx, deps.Repos.BuyItemRepository),
		CategoryService: categoryImpl.NewCategoryService(ctx, deps.Repos.CategoryRepository),
		ClientService:   clientImpl.NewClientService(ctx, deps.Repos.ClientRepository),
	}
}
