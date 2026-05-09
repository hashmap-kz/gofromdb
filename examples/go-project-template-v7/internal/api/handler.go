package api

import (
	"net/http"

	"go-project-template-v5/internal/api/categories"
	"go-project-template-v5/internal/api/order_items"
	"go-project-template-v5/internal/api/orders"
	"go-project-template-v5/internal/api/products"
	"go-project-template-v5/internal/api/users"
)

type Handler struct {
	Services *Services
}

func NewHandler(services *Services) *Handler {
	return &Handler{
		Services: services,
	}
}

func (h *Handler) Init(router *http.ServeMux) {
	// Categories routes
	categoriesHandler := categories.NewHandler(h.Services.Categories)
	router.HandleFunc("POST /api/v1/categories", categoriesHandler.Save)
	router.HandleFunc("PUT /api/v1/categories/{record_id}", categoriesHandler.UpdateByID)
	router.HandleFunc("DELETE /api/v1/categories/{record_id}", categoriesHandler.DeleteByID)
	router.HandleFunc("GET /api/v1/categories/{record_id}", categoriesHandler.FindByID)
	router.HandleFunc("GET /api/v1/categories", categoriesHandler.FindAll)
	router.HandleFunc("GET /api/v1/categories/pageable", categoriesHandler.FindAllPageable)

	// OrderItems routes
	orderItemsHandler := order_items.NewHandler(h.Services.OrderItems)
	router.HandleFunc("POST /api/v1/order-items", orderItemsHandler.Save)
	router.HandleFunc("PUT /api/v1/order-items/{record_id}", orderItemsHandler.UpdateByID)
	router.HandleFunc("DELETE /api/v1/order-items/{record_id}", orderItemsHandler.DeleteByID)
	router.HandleFunc("GET /api/v1/order-items/{record_id}", orderItemsHandler.FindByID)
	router.HandleFunc("GET /api/v1/order-items", orderItemsHandler.FindAll)
	router.HandleFunc("GET /api/v1/order-items/pageable", orderItemsHandler.FindAllPageable)

	// Orders routes
	ordersHandler := orders.NewHandler(h.Services.Orders)
	router.HandleFunc("POST /api/v1/orders", ordersHandler.Save)
	router.HandleFunc("PUT /api/v1/orders/{record_id}", ordersHandler.UpdateByID)
	router.HandleFunc("DELETE /api/v1/orders/{record_id}", ordersHandler.DeleteByID)
	router.HandleFunc("GET /api/v1/orders/{record_id}", ordersHandler.FindByID)
	router.HandleFunc("GET /api/v1/orders", ordersHandler.FindAll)
	router.HandleFunc("GET /api/v1/orders/pageable", ordersHandler.FindAllPageable)

	// Products routes
	productsHandler := products.NewHandler(h.Services.Products)
	router.HandleFunc("POST /api/v1/products", productsHandler.Save)
	router.HandleFunc("PUT /api/v1/products/{record_id}", productsHandler.UpdateByID)
	router.HandleFunc("DELETE /api/v1/products/{record_id}", productsHandler.DeleteByID)
	router.HandleFunc("GET /api/v1/products/{record_id}", productsHandler.FindByID)
	router.HandleFunc("GET /api/v1/products", productsHandler.FindAll)
	router.HandleFunc("GET /api/v1/products/pageable", productsHandler.FindAllPageable)

	// Users routes
	usersHandler := users.NewHandler(h.Services.Users)
	router.HandleFunc("POST /api/v1/users", usersHandler.Save)
	router.HandleFunc("PUT /api/v1/users/{record_id}", usersHandler.UpdateByID)
	router.HandleFunc("DELETE /api/v1/users/{record_id}", usersHandler.DeleteByID)
	router.HandleFunc("GET /api/v1/users/{record_id}", usersHandler.FindByID)
	router.HandleFunc("GET /api/v1/users", usersHandler.FindAll)
	router.HandleFunc("GET /api/v1/users/pageable", usersHandler.FindAllPageable)
}
