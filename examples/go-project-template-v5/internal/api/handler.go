package api

import (
	"net/http"

	buyv1 "go-project-template-v5/internal/api/buy/handler/v1"
	buyitemv1 "go-project-template-v5/internal/api/buy_item/handler/v1"
	categoryv1 "go-project-template-v5/internal/api/category/handler/v1"
	clientv1 "go-project-template-v5/internal/api/client/handler/v1"
	productv1 "go-project-template-v5/internal/api/product/handler/v1"
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
	// Buy routes
	buyHandler := buyv1.NewBuyHTTPHandler(h.Services.BuyService)
	router.HandleFunc("POST /api/v1/buys", buyHandler.Save)
	router.HandleFunc("PUT /api/v1/buys/{id}", buyHandler.UpdateByID)
	router.HandleFunc("DELETE /api/v1/buys/{id}", buyHandler.DeleteByID)
	router.HandleFunc("GET /api/v1/buys/{id}", buyHandler.FindByID)
	router.HandleFunc("GET /api/v1/buys", buyHandler.FindAll)
	router.HandleFunc("GET /api/v1/buys/pageable", buyHandler.FindAllPageable)

	// BuyItem routes
	buyItemHandler := buyitemv1.NewBuyItemHTTPHandler(h.Services.BuyItemService)
	router.HandleFunc("POST /api/v1/buy-items", buyItemHandler.Save)
	router.HandleFunc("PUT /api/v1/buy-items/{id}", buyItemHandler.UpdateByID)
	router.HandleFunc("DELETE /api/v1/buy-items/{id}", buyItemHandler.DeleteByID)
	router.HandleFunc("GET /api/v1/buy-items/{id}", buyItemHandler.FindByID)
	router.HandleFunc("GET /api/v1/buy-items", buyItemHandler.FindAll)
	router.HandleFunc("GET /api/v1/buy-items/pageable", buyItemHandler.FindAllPageable)

	// Category routes
	categoryHandler := categoryv1.NewCategoryHTTPHandler(h.Services.CategoryService)
	router.HandleFunc("POST /api/v1/categories", categoryHandler.Save)
	router.HandleFunc("PUT /api/v1/categories/{id}", categoryHandler.UpdateByID)
	router.HandleFunc("DELETE /api/v1/categories/{id}", categoryHandler.DeleteByID)
	router.HandleFunc("GET /api/v1/categories/{id}", categoryHandler.FindByID)
	router.HandleFunc("GET /api/v1/categories", categoryHandler.FindAll)
	router.HandleFunc("GET /api/v1/categories/pageable", categoryHandler.FindAllPageable)

	// Client routes
	clientHandler := clientv1.NewClientHTTPHandler(h.Services.ClientService)
	router.HandleFunc("POST /api/v1/clients", clientHandler.Save)
	router.HandleFunc("PUT /api/v1/clients/{id}", clientHandler.UpdateByID)
	router.HandleFunc("DELETE /api/v1/clients/{id}", clientHandler.DeleteByID)
	router.HandleFunc("GET /api/v1/clients/{id}", clientHandler.FindByID)
	router.HandleFunc("GET /api/v1/clients", clientHandler.FindAll)
	router.HandleFunc("GET /api/v1/clients/pageable", clientHandler.FindAllPageable)

	// Product routes
	productHandler := productv1.NewProductHTTPHandler(h.Services.ProductService)
	router.HandleFunc("POST /api/v1/products", productHandler.Save)
	router.HandleFunc("PUT /api/v1/products/{id}", productHandler.UpdateByID)
	router.HandleFunc("DELETE /api/v1/products/{id}", productHandler.DeleteByID)
	router.HandleFunc("GET /api/v1/products/{id}", productHandler.FindByID)
	router.HandleFunc("GET /api/v1/products", productHandler.FindAll)
	router.HandleFunc("GET /api/v1/products/pageable", productHandler.FindAllPageable)
}
