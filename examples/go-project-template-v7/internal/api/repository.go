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
	"go-project-template-v7/pkg/storage/postgres"
)

// Init all repos

type Repositories struct {
	Authors              authors.Repository
	BookAuthors          book_authors.Repository
	BookTranslations     book_translations.Repository
	Books                books.Repository
	BookstoreSalesOrders bookstore_sales_orders.Repository
	Categories           categories.Repository
	Customers            customers.Repository
	DiscountCodes        discount_codes.Repository
	ImportBatches        import_batches.Repository
	ImportErrors         import_errors.Repository
	OrderItems           order_items.Repository
	OrderLines           order_lines.Repository
	Products             products.Repository
	PublicOrders         public_orders.Repository
	Publishers           publishers.Repository
	StockEvents          stock_events.Repository
	StockLevels          stock_levels.Repository
	Users                users.Repository
	Warehouses           warehouses.Repository
}

func NewRepositories(ctx context.Context, db *postgres.Postgres) *Repositories {
	return &Repositories{
		Authors:              authors.NewRepository(ctx, db),
		BookAuthors:          book_authors.NewRepository(ctx, db),
		BookTranslations:     book_translations.NewRepository(ctx, db),
		Books:                books.NewRepository(ctx, db),
		BookstoreSalesOrders: bookstore_sales_orders.NewRepository(ctx, db),
		Categories:           categories.NewRepository(ctx, db),
		Customers:            customers.NewRepository(ctx, db),
		DiscountCodes:        discount_codes.NewRepository(ctx, db),
		ImportBatches:        import_batches.NewRepository(ctx, db),
		ImportErrors:         import_errors.NewRepository(ctx, db),
		OrderItems:           order_items.NewRepository(ctx, db),
		OrderLines:           order_lines.NewRepository(ctx, db),
		Products:             products.NewRepository(ctx, db),
		PublicOrders:         public_orders.NewRepository(ctx, db),
		Publishers:           publishers.NewRepository(ctx, db),
		StockEvents:          stock_events.NewRepository(ctx, db),
		StockLevels:          stock_levels.NewRepository(ctx, db),
		Users:                users.NewRepository(ctx, db),
		Warehouses:           warehouses.NewRepository(ctx, db),
	}
}
