package api

import (
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
	"net/http"
)

type Handler struct {
	Services *Services
}

func NewHandler(services *Services) *Handler {
	return &Handler{
		Services: services,
	}
}

func (h *Handler) Mount(router *http.ServeMux) {
	// Authors routes
	authorsHandler := authors.NewHandler(h.Services.Authors)
	router.HandleFunc("POST /api/v1/authors", authorsHandler.Save)
	router.HandleFunc("PUT /api/v1/authors/{author_id}", authorsHandler.UpdateByID)
	router.HandleFunc("DELETE /api/v1/authors/{author_id}", authorsHandler.DeleteByID)
	router.HandleFunc("GET /api/v1/authors/{author_id}", authorsHandler.FindByID)
	router.HandleFunc("GET /api/v1/authors", authorsHandler.FindAll)
	router.HandleFunc("GET /api/v1/authors/pageable", authorsHandler.FindAllPageable)

	// BookAuthors routes
	bookAuthorsHandler := book_authors.NewHandler(h.Services.BookAuthors)
	router.HandleFunc("POST /api/v1/book-authors", bookAuthorsHandler.Save)
	router.HandleFunc("PUT /api/v1/book-authors/{book_id}/{author_id}", bookAuthorsHandler.UpdateByID)
	router.HandleFunc("DELETE /api/v1/book-authors/{book_id}/{author_id}", bookAuthorsHandler.DeleteByID)
	router.HandleFunc("GET /api/v1/book-authors/{book_id}/{author_id}", bookAuthorsHandler.FindByID)
	router.HandleFunc("GET /api/v1/book-authors", bookAuthorsHandler.FindAll)
	router.HandleFunc("GET /api/v1/book-authors/pageable", bookAuthorsHandler.FindAllPageable)

	// BookTranslations routes
	bookTranslationsHandler := book_translations.NewHandler(h.Services.BookTranslations)
	router.HandleFunc("POST /api/v1/book-translations", bookTranslationsHandler.Save)
	router.HandleFunc("PUT /api/v1/book-translations/{book_id}/{language_code}", bookTranslationsHandler.UpdateByID)
	router.HandleFunc("DELETE /api/v1/book-translations/{book_id}/{language_code}", bookTranslationsHandler.DeleteByID)
	router.HandleFunc("GET /api/v1/book-translations/{book_id}/{language_code}", bookTranslationsHandler.FindByID)
	router.HandleFunc("GET /api/v1/book-translations", bookTranslationsHandler.FindAll)
	router.HandleFunc("GET /api/v1/book-translations/pageable", bookTranslationsHandler.FindAllPageable)

	// Books routes
	booksHandler := books.NewHandler(h.Services.Books)
	router.HandleFunc("POST /api/v1/books", booksHandler.Save)
	router.HandleFunc("PUT /api/v1/books/{book_id}", booksHandler.UpdateByID)
	router.HandleFunc("DELETE /api/v1/books/{book_id}", booksHandler.DeleteByID)
	router.HandleFunc("GET /api/v1/books/{book_id}", booksHandler.FindByID)
	router.HandleFunc("GET /api/v1/books", booksHandler.FindAll)
	router.HandleFunc("GET /api/v1/books/pageable", booksHandler.FindAllPageable)

	// BookstoreSalesOrders routes
	bookstoreSalesOrdersHandler := bookstore_sales_orders.NewHandler(h.Services.BookstoreSalesOrders)
	router.HandleFunc("POST /api/v1/bookstore-sales-orders", bookstoreSalesOrdersHandler.Save)
	router.HandleFunc("PUT /api/v1/bookstore-sales-orders/{order_id}", bookstoreSalesOrdersHandler.UpdateByID)
	router.HandleFunc("DELETE /api/v1/bookstore-sales-orders/{order_id}", bookstoreSalesOrdersHandler.DeleteByID)
	router.HandleFunc("GET /api/v1/bookstore-sales-orders/{order_id}", bookstoreSalesOrdersHandler.FindByID)
	router.HandleFunc("GET /api/v1/bookstore-sales-orders", bookstoreSalesOrdersHandler.FindAll)
	router.HandleFunc("GET /api/v1/bookstore-sales-orders/pageable", bookstoreSalesOrdersHandler.FindAllPageable)

	// Categories routes
	categoriesHandler := categories.NewHandler(h.Services.Categories)
	router.HandleFunc("POST /api/v1/categories", categoriesHandler.Save)
	router.HandleFunc("PUT /api/v1/categories/{record_id}", categoriesHandler.UpdateByID)
	router.HandleFunc("DELETE /api/v1/categories/{record_id}", categoriesHandler.DeleteByID)
	router.HandleFunc("GET /api/v1/categories/{record_id}", categoriesHandler.FindByID)
	router.HandleFunc("GET /api/v1/categories", categoriesHandler.FindAll)
	router.HandleFunc("GET /api/v1/categories/pageable", categoriesHandler.FindAllPageable)

	// Customers routes
	customersHandler := customers.NewHandler(h.Services.Customers)
	router.HandleFunc("POST /api/v1/customers", customersHandler.Save)
	router.HandleFunc("PUT /api/v1/customers/{customer_id}", customersHandler.UpdateByID)
	router.HandleFunc("DELETE /api/v1/customers/{customer_id}", customersHandler.DeleteByID)
	router.HandleFunc("GET /api/v1/customers/{customer_id}", customersHandler.FindByID)
	router.HandleFunc("GET /api/v1/customers", customersHandler.FindAll)
	router.HandleFunc("GET /api/v1/customers/pageable", customersHandler.FindAllPageable)

	// DiscountCodes routes
	discountCodesHandler := discount_codes.NewHandler(h.Services.DiscountCodes)
	router.HandleFunc("POST /api/v1/discount-codes", discountCodesHandler.Save)
	router.HandleFunc("PUT /api/v1/discount-codes/{code}", discountCodesHandler.UpdateByID)
	router.HandleFunc("DELETE /api/v1/discount-codes/{code}", discountCodesHandler.DeleteByID)
	router.HandleFunc("GET /api/v1/discount-codes/{code}", discountCodesHandler.FindByID)
	router.HandleFunc("GET /api/v1/discount-codes", discountCodesHandler.FindAll)
	router.HandleFunc("GET /api/v1/discount-codes/pageable", discountCodesHandler.FindAllPageable)

	// ImportBatches routes
	importBatchesHandler := import_batches.NewHandler(h.Services.ImportBatches)
	router.HandleFunc("POST /api/v1/import-batches", importBatchesHandler.Save)
	router.HandleFunc("PUT /api/v1/import-batches/{source_name}/{batch_no}", importBatchesHandler.UpdateByID)
	router.HandleFunc("DELETE /api/v1/import-batches/{source_name}/{batch_no}", importBatchesHandler.DeleteByID)
	router.HandleFunc("GET /api/v1/import-batches/{source_name}/{batch_no}", importBatchesHandler.FindByID)
	router.HandleFunc("GET /api/v1/import-batches", importBatchesHandler.FindAll)
	router.HandleFunc("GET /api/v1/import-batches/pageable", importBatchesHandler.FindAllPageable)

	// ImportErrors routes
	importErrorsHandler := import_errors.NewHandler(h.Services.ImportErrors)
	router.HandleFunc("POST /api/v1/import-errors", importErrorsHandler.Save)
	router.HandleFunc("GET /api/v1/import-errors", importErrorsHandler.FindAll)
	router.HandleFunc("GET /api/v1/import-errors/pageable", importErrorsHandler.FindAllPageable)

	// OrderItems routes
	orderItemsHandler := order_items.NewHandler(h.Services.OrderItems)
	router.HandleFunc("POST /api/v1/order-items", orderItemsHandler.Save)
	router.HandleFunc("PUT /api/v1/order-items/{record_id}", orderItemsHandler.UpdateByID)
	router.HandleFunc("DELETE /api/v1/order-items/{record_id}", orderItemsHandler.DeleteByID)
	router.HandleFunc("GET /api/v1/order-items/{record_id}", orderItemsHandler.FindByID)
	router.HandleFunc("GET /api/v1/order-items", orderItemsHandler.FindAll)
	router.HandleFunc("GET /api/v1/order-items/pageable", orderItemsHandler.FindAllPageable)

	// OrderLines routes
	orderLinesHandler := order_lines.NewHandler(h.Services.OrderLines)
	router.HandleFunc("POST /api/v1/order-lines", orderLinesHandler.Save)
	router.HandleFunc("PUT /api/v1/order-lines/{order_id}/{line_no}", orderLinesHandler.UpdateByID)
	router.HandleFunc("DELETE /api/v1/order-lines/{order_id}/{line_no}", orderLinesHandler.DeleteByID)
	router.HandleFunc("GET /api/v1/order-lines/{order_id}/{line_no}", orderLinesHandler.FindByID)
	router.HandleFunc("GET /api/v1/order-lines", orderLinesHandler.FindAll)
	router.HandleFunc("GET /api/v1/order-lines/pageable", orderLinesHandler.FindAllPageable)

	// Products routes
	productsHandler := products.NewHandler(h.Services.Products)
	router.HandleFunc("POST /api/v1/products", productsHandler.Save)
	router.HandleFunc("PUT /api/v1/products/{record_id}", productsHandler.UpdateByID)
	router.HandleFunc("DELETE /api/v1/products/{record_id}", productsHandler.DeleteByID)
	router.HandleFunc("GET /api/v1/products/{record_id}", productsHandler.FindByID)
	router.HandleFunc("GET /api/v1/products", productsHandler.FindAll)
	router.HandleFunc("GET /api/v1/products/pageable", productsHandler.FindAllPageable)

	// PublicOrders routes
	publicOrdersHandler := public_orders.NewHandler(h.Services.PublicOrders)
	router.HandleFunc("POST /api/v1/public-orders", publicOrdersHandler.Save)
	router.HandleFunc("PUT /api/v1/public-orders/{record_id}", publicOrdersHandler.UpdateByID)
	router.HandleFunc("DELETE /api/v1/public-orders/{record_id}", publicOrdersHandler.DeleteByID)
	router.HandleFunc("GET /api/v1/public-orders/{record_id}", publicOrdersHandler.FindByID)
	router.HandleFunc("GET /api/v1/public-orders", publicOrdersHandler.FindAll)
	router.HandleFunc("GET /api/v1/public-orders/pageable", publicOrdersHandler.FindAllPageable)

	// Publishers routes
	publishersHandler := publishers.NewHandler(h.Services.Publishers)
	router.HandleFunc("POST /api/v1/publishers", publishersHandler.Save)
	router.HandleFunc("PUT /api/v1/publishers/{code}", publishersHandler.UpdateByID)
	router.HandleFunc("DELETE /api/v1/publishers/{code}", publishersHandler.DeleteByID)
	router.HandleFunc("GET /api/v1/publishers/{code}", publishersHandler.FindByID)
	router.HandleFunc("GET /api/v1/publishers", publishersHandler.FindAll)
	router.HandleFunc("GET /api/v1/publishers/pageable", publishersHandler.FindAllPageable)

	// StockEvents routes
	stockEventsHandler := stock_events.NewHandler(h.Services.StockEvents)
	router.HandleFunc("POST /api/v1/stock-events", stockEventsHandler.Save)
	router.HandleFunc("GET /api/v1/stock-events", stockEventsHandler.FindAll)
	router.HandleFunc("GET /api/v1/stock-events/pageable", stockEventsHandler.FindAllPageable)

	// StockLevels routes
	stockLevelsHandler := stock_levels.NewHandler(h.Services.StockLevels)
	router.HandleFunc("POST /api/v1/stock-levels", stockLevelsHandler.Save)
	router.HandleFunc("PUT /api/v1/stock-levels/{warehouse_code}/{book_id}", stockLevelsHandler.UpdateByID)
	router.HandleFunc("DELETE /api/v1/stock-levels/{warehouse_code}/{book_id}", stockLevelsHandler.DeleteByID)
	router.HandleFunc("GET /api/v1/stock-levels/{warehouse_code}/{book_id}", stockLevelsHandler.FindByID)
	router.HandleFunc("GET /api/v1/stock-levels", stockLevelsHandler.FindAll)
	router.HandleFunc("GET /api/v1/stock-levels/pageable", stockLevelsHandler.FindAllPageable)

	// Users routes
	usersHandler := users.NewHandler(h.Services.Users)
	router.HandleFunc("POST /api/v1/users", usersHandler.Save)
	router.HandleFunc("PUT /api/v1/users/{record_id}", usersHandler.UpdateByID)
	router.HandleFunc("DELETE /api/v1/users/{record_id}", usersHandler.DeleteByID)
	router.HandleFunc("GET /api/v1/users/{record_id}", usersHandler.FindByID)
	router.HandleFunc("GET /api/v1/users", usersHandler.FindAll)
	router.HandleFunc("GET /api/v1/users/pageable", usersHandler.FindAllPageable)

	// Warehouses routes
	warehousesHandler := warehouses.NewHandler(h.Services.Warehouses)
	router.HandleFunc("POST /api/v1/warehouses", warehousesHandler.Save)
	router.HandleFunc("PUT /api/v1/warehouses/{code}", warehousesHandler.UpdateByID)
	router.HandleFunc("DELETE /api/v1/warehouses/{code}", warehousesHandler.DeleteByID)
	router.HandleFunc("GET /api/v1/warehouses/{code}", warehousesHandler.FindByID)
	router.HandleFunc("GET /api/v1/warehouses", warehousesHandler.FindAll)
	router.HandleFunc("GET /api/v1/warehouses/pageable", warehousesHandler.FindAllPageable)
}
