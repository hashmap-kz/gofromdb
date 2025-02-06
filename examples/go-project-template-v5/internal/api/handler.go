package api

import (
	"net/http"

	categoriesv1 "go-project-template-v5/internal/api/categories/handler/v1"
	clientsv1 "go-project-template-v5/internal/api/clients/handler/v1"
	customerOrderItemsv1 "go-project-template-v5/internal/api/customer_order_items/handler/v1"
	customerOrdersv1 "go-project-template-v5/internal/api/customer_orders/handler/v1"
	productsv1 "go-project-template-v5/internal/api/products/handler/v1"
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
	categoriesHandler := categoriesv1.NewCategoriesHTTPHandler(h.Services.CategoriesService)
	router.HandleFunc("POST /api/v1/categories", categoriesHandler.Save)
	router.HandleFunc("PUT /api/v1/categories/{record_id}", categoriesHandler.UpdateByID)
	router.HandleFunc("DELETE /api/v1/categories/{record_id}", categoriesHandler.DeleteByID)
	router.HandleFunc("GET /api/v1/categories/{record_id}", categoriesHandler.FindByID)
	router.HandleFunc("GET /api/v1/categories", categoriesHandler.FindAll)
	router.HandleFunc("GET /api/v1/categories/pageable", categoriesHandler.FindAllPageable)

	// Clients routes
	clientsHandler := clientsv1.NewClientsHTTPHandler(h.Services.ClientsService)
	router.HandleFunc("POST /api/v1/clients", clientsHandler.Save)
	router.HandleFunc("PUT /api/v1/clients/{record_id}", clientsHandler.UpdateByID)
	router.HandleFunc("DELETE /api/v1/clients/{record_id}", clientsHandler.DeleteByID)
	router.HandleFunc("GET /api/v1/clients/{record_id}", clientsHandler.FindByID)
	router.HandleFunc("GET /api/v1/clients", clientsHandler.FindAll)
	router.HandleFunc("GET /api/v1/clients/pageable", clientsHandler.FindAllPageable)

	// CustomerOrderItems routes
	customerOrderItemsHandler := customerOrderItemsv1.NewCustomerOrderItemsHTTPHandler(h.Services.CustomerOrderItemsService)
	router.HandleFunc("POST /api/v1/customer-order-items", customerOrderItemsHandler.Save)
	router.HandleFunc("PUT /api/v1/customer-order-items/{record_id}", customerOrderItemsHandler.UpdateByID)
	router.HandleFunc("DELETE /api/v1/customer-order-items/{record_id}", customerOrderItemsHandler.DeleteByID)
	router.HandleFunc("GET /api/v1/customer-order-items/{record_id}", customerOrderItemsHandler.FindByID)
	router.HandleFunc("GET /api/v1/customer-order-items", customerOrderItemsHandler.FindAll)
	router.HandleFunc("GET /api/v1/customer-order-items/pageable", customerOrderItemsHandler.FindAllPageable)

	// CustomerOrders routes
	customerOrdersHandler := customerOrdersv1.NewCustomerOrdersHTTPHandler(h.Services.CustomerOrdersService)
	router.HandleFunc("POST /api/v1/customer-orders", customerOrdersHandler.Save)
	router.HandleFunc("PUT /api/v1/customer-orders/{record_id}", customerOrdersHandler.UpdateByID)
	router.HandleFunc("DELETE /api/v1/customer-orders/{record_id}", customerOrdersHandler.DeleteByID)
	router.HandleFunc("GET /api/v1/customer-orders/{record_id}", customerOrdersHandler.FindByID)
	router.HandleFunc("GET /api/v1/customer-orders", customerOrdersHandler.FindAll)
	router.HandleFunc("GET /api/v1/customer-orders/pageable", customerOrdersHandler.FindAllPageable)

	// Products routes
	productsHandler := productsv1.NewProductsHTTPHandler(h.Services.ProductsService)
	router.HandleFunc("POST /api/v1/products", productsHandler.Save)
	router.HandleFunc("PUT /api/v1/products/{record_id}", productsHandler.UpdateByID)
	router.HandleFunc("DELETE /api/v1/products/{record_id}", productsHandler.DeleteByID)
	router.HandleFunc("GET /api/v1/products/{record_id}", productsHandler.FindByID)
	router.HandleFunc("GET /api/v1/products", productsHandler.FindAll)
	router.HandleFunc("GET /api/v1/products/pageable", productsHandler.FindAllPageable)
}
