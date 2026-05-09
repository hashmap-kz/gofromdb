package api

import (
	"context"
	"go-project-template-v7/internal/api/bookstore_catalog/authors"
	"go-project-template-v7/internal/api/bookstore_catalog/book_authors"
	"go-project-template-v7/internal/api/bookstore_catalog/book_translations"
	"go-project-template-v7/internal/api/bookstore_catalog/books"
	"go-project-template-v7/internal/api/bookstore_catalog/publishers"
	"go-project-template-v7/internal/api/bookstore_import/import_batches"
	"go-project-template-v7/internal/api/bookstore_import/import_errors"
	"go-project-template-v7/internal/api/bookstore_inventory/stock_events"
	"go-project-template-v7/internal/api/bookstore_inventory/stock_levels"
	"go-project-template-v7/internal/api/bookstore_inventory/warehouses"
	"go-project-template-v7/internal/api/bookstore_sales/customers"
	"go-project-template-v7/internal/api/bookstore_sales/discount_codes"
	"go-project-template-v7/internal/api/bookstore_sales/order_lines"
	bookstore_sales_orders "go-project-template-v7/internal/api/bookstore_sales/orders"
	"go-project-template-v7/internal/api/public/categories"
	"go-project-template-v7/internal/api/public/order_items"
	public_orders "go-project-template-v7/internal/api/public/orders"
	"go-project-template-v7/internal/api/public/products"
	"go-project-template-v7/internal/api/public/users"
)

// Init all services

type Services struct {
	Authors              authors.Service
	BookAuthors          book_authors.Service
	BookTranslations     book_translations.Service
	Books                books.Service
	BookstoreSalesOrders bookstore_sales_orders.Service
	Categories           categories.Service
	Customers            customers.Service
	DiscountCodes        discount_codes.Service
	ImportBatches        import_batches.Service
	ImportErrors         import_errors.Service
	OrderItems           order_items.Service
	OrderLines           order_lines.Service
	Products             products.Service
	PublicOrders         public_orders.Service
	Publishers           publishers.Service
	StockEvents          stock_events.Service
	StockLevels          stock_levels.Service
	Users                users.Service
	Warehouses           warehouses.Service
}

type Deps struct {
	// TODO: other deps here
	Repos *Repositories
}

func NewServices(ctx context.Context, deps Deps) *Services {
	return &Services{
		Authors:              authors.NewService(ctx, deps.Repos.Authors),
		BookAuthors:          book_authors.NewService(ctx, deps.Repos.BookAuthors),
		BookTranslations:     book_translations.NewService(ctx, deps.Repos.BookTranslations),
		Books:                books.NewService(ctx, deps.Repos.Books),
		BookstoreSalesOrders: bookstore_sales_orders.NewService(ctx, deps.Repos.BookstoreSalesOrders),
		Categories:           categories.NewService(ctx, deps.Repos.Categories),
		Customers:            customers.NewService(ctx, deps.Repos.Customers),
		DiscountCodes:        discount_codes.NewService(ctx, deps.Repos.DiscountCodes),
		ImportBatches:        import_batches.NewService(ctx, deps.Repos.ImportBatches),
		ImportErrors:         import_errors.NewService(ctx, deps.Repos.ImportErrors),
		OrderItems:           order_items.NewService(ctx, deps.Repos.OrderItems),
		OrderLines:           order_lines.NewService(ctx, deps.Repos.OrderLines),
		Products:             products.NewService(ctx, deps.Repos.Products),
		PublicOrders:         public_orders.NewService(ctx, deps.Repos.PublicOrders),
		Publishers:           publishers.NewService(ctx, deps.Repos.Publishers),
		StockEvents:          stock_events.NewService(ctx, deps.Repos.StockEvents),
		StockLevels:          stock_levels.NewService(ctx, deps.Repos.StockLevels),
		Users:                users.NewService(ctx, deps.Repos.Users),
		Warehouses:           warehouses.NewService(ctx, deps.Repos.Warehouses),
	}
}
