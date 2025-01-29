package api

import (
	"context"

	buyRepo "go-project-template-v5/internal/api/buy/repository"
	buyImpl "go-project-template-v5/internal/api/buy/repository/impl"
	buyitemRepo "go-project-template-v5/internal/api/buy_item/repository"
	buyitemImpl "go-project-template-v5/internal/api/buy_item/repository/impl"
	categoryRepo "go-project-template-v5/internal/api/category/repository"
	categoryImpl "go-project-template-v5/internal/api/category/repository/impl"
	clientRepo "go-project-template-v5/internal/api/client/repository"
	clientImpl "go-project-template-v5/internal/api/client/repository/impl"
	productRepo "go-project-template-v5/internal/api/product/repository"
	productImpl "go-project-template-v5/internal/api/product/repository/impl"

	"go-project-template-v5/pkg/storage/postgres"
)

// Init all repos

type Repositories struct {
	BuyItemRepository  buyitemRepo.BuyItemRepository
	CategoryRepository categoryRepo.CategoryRepository
	ClientRepository   clientRepo.ClientRepository
	ProductRepository  productRepo.ProductRepository
	BuyRepository      buyRepo.BuyRepository
}

func NewRepositories(ctx context.Context, db *postgres.Postgres) *Repositories {
	return &Repositories{
		BuyItemRepository:  buyitemImpl.NewBuyItemRepository(ctx, db),
		CategoryRepository: categoryImpl.NewCategoryRepository(ctx, db),
		ClientRepository:   clientImpl.NewClientRepository(ctx, db),
		ProductRepository:  productImpl.NewProductRepository(ctx, db),
		BuyRepository:      buyImpl.NewBuyRepository(ctx, db),
	}
}
