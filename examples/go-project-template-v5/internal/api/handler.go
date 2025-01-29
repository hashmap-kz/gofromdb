package api

import (
	buyv1 "go-project-template-v5/internal/api/buy/handler/v1"
	buyitemv1 "go-project-template-v5/internal/api/buy_item/handler/v1"
	categoryv1 "go-project-template-v5/internal/api/category/handler/v1"
	clientv1 "go-project-template-v5/internal/api/client/handler/v1"
	productv1 "go-project-template-v5/internal/api/product/handler/v1"
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

func (h *Handler) Init(router *http.ServeMux) {
	// Buy routing
	buyHandler := buyv1.NewBuyHTTPHandler(h.Services.BuyService)
	router.HandleFunc("POST /api/v1/buys", buyHandler.Save)
	router.HandleFunc("GET /api/v1/buys", buyHandler.GetAll)
	router.HandleFunc("GET /api/v1/buys/pageable", buyHandler.GetAllPaginated)
	router.HandleFunc("GET /api/v1/buys/{id}", buyHandler.GetByID)
	router.HandleFunc("PUT /api/v1/buys/{id}", buyHandler.Update)
	router.HandleFunc("DELETE /api/v1/buys/{id}", buyHandler.Delete)

	// BuyItem routing
	buyItemHandler := buyitemv1.NewBuyItemHTTPHandler(h.Services.BuyItemService)
	router.HandleFunc("POST /api/v1/buy-items", buyItemHandler.Save)
	router.HandleFunc("GET /api/v1/buy-items", buyItemHandler.GetAll)
	router.HandleFunc("GET /api/v1/buy-items/pageable", buyItemHandler.GetAllPaginated)
	router.HandleFunc("GET /api/v1/buy-items/{id}", buyItemHandler.GetByID)
	router.HandleFunc("PUT /api/v1/buy-items/{id}", buyItemHandler.Update)
	router.HandleFunc("DELETE /api/v1/buy-items/{id}", buyItemHandler.Delete)

	// Category routing
	categoryHandler := categoryv1.NewCategoryHTTPHandler(h.Services.CategoryService)
	router.HandleFunc("POST /api/v1/categories", categoryHandler.Save)
	router.HandleFunc("GET /api/v1/categories", categoryHandler.GetAll)
	router.HandleFunc("GET /api/v1/categories/pageable", categoryHandler.GetAllPaginated)
	router.HandleFunc("GET /api/v1/categories/{id}", categoryHandler.GetByID)
	router.HandleFunc("PUT /api/v1/categories/{id}", categoryHandler.Update)
	router.HandleFunc("DELETE /api/v1/categories/{id}", categoryHandler.Delete)

	// Client routing
	clientHandler := clientv1.NewClientHTTPHandler(h.Services.ClientService)
	router.HandleFunc("POST /api/v1/clients", clientHandler.Save)
	router.HandleFunc("GET /api/v1/clients", clientHandler.GetAll)
	router.HandleFunc("GET /api/v1/clients/pageable", clientHandler.GetAllPaginated)
	router.HandleFunc("GET /api/v1/clients/{id}", clientHandler.GetByID)
	router.HandleFunc("PUT /api/v1/clients/{id}", clientHandler.Update)
	router.HandleFunc("DELETE /api/v1/clients/{id}", clientHandler.Delete)

	// Product routing
	productHandler := productv1.NewProductHTTPHandler(h.Services.ProductService)
	router.HandleFunc("POST /api/v1/products", productHandler.Save)
	router.HandleFunc("GET /api/v1/products", productHandler.GetAll)
	router.HandleFunc("GET /api/v1/products/pageable", productHandler.GetAllPaginated)
	router.HandleFunc("GET /api/v1/products/{id}", productHandler.GetByID)
	router.HandleFunc("PUT /api/v1/products/{id}", productHandler.Update)
	router.HandleFunc("DELETE /api/v1/products/{id}", productHandler.Delete)

}
