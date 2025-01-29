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
	// Product routing
	productHandler := productv1.NewProductHTTPHandler(h.Services.ProductService)
	router.HandleFunc("POST /api/v1/product", productHandler.Save)
	router.HandleFunc("GET /api/v1/product", productHandler.GetAll)
	router.HandleFunc("GET /api/v1/product/pageable", productHandler.GetAllPaginated)
	router.HandleFunc("GET /api/v1/product/{id}", productHandler.GetByID)
	router.HandleFunc("PUT /api/v1/product/{id}", productHandler.Update)
	router.HandleFunc("DELETE /api/v1/product/{id}", productHandler.Delete)

	// Buy routing
	buyHandler := buyv1.NewBuyHTTPHandler(h.Services.BuyService)
	router.HandleFunc("POST /api/v1/buy", buyHandler.Save)
	router.HandleFunc("GET /api/v1/buy", buyHandler.GetAll)
	router.HandleFunc("GET /api/v1/buy/pageable", buyHandler.GetAllPaginated)
	router.HandleFunc("GET /api/v1/buy/{id}", buyHandler.GetByID)
	router.HandleFunc("PUT /api/v1/buy/{id}", buyHandler.Update)
	router.HandleFunc("DELETE /api/v1/buy/{id}", buyHandler.Delete)

	// BuyItem routing
	buyItemHandler := buyitemv1.NewBuyItemHTTPHandler(h.Services.BuyItemService)
	router.HandleFunc("POST /api/v1/buy_item", buyItemHandler.Save)
	router.HandleFunc("GET /api/v1/buy_item", buyItemHandler.GetAll)
	router.HandleFunc("GET /api/v1/buy_item/pageable", buyItemHandler.GetAllPaginated)
	router.HandleFunc("GET /api/v1/buy_item/{id}", buyItemHandler.GetByID)
	router.HandleFunc("PUT /api/v1/buy_item/{id}", buyItemHandler.Update)
	router.HandleFunc("DELETE /api/v1/buy_item/{id}", buyItemHandler.Delete)

	// Category routing
	categoryHandler := categoryv1.NewCategoryHTTPHandler(h.Services.CategoryService)
	router.HandleFunc("POST /api/v1/category", categoryHandler.Save)
	router.HandleFunc("GET /api/v1/category", categoryHandler.GetAll)
	router.HandleFunc("GET /api/v1/category/pageable", categoryHandler.GetAllPaginated)
	router.HandleFunc("GET /api/v1/category/{id}", categoryHandler.GetByID)
	router.HandleFunc("PUT /api/v1/category/{id}", categoryHandler.Update)
	router.HandleFunc("DELETE /api/v1/category/{id}", categoryHandler.Delete)

	// Client routing
	clientHandler := clientv1.NewClientHTTPHandler(h.Services.ClientService)
	router.HandleFunc("POST /api/v1/client", clientHandler.Save)
	router.HandleFunc("GET /api/v1/client", clientHandler.GetAll)
	router.HandleFunc("GET /api/v1/client/pageable", clientHandler.GetAllPaginated)
	router.HandleFunc("GET /api/v1/client/{id}", clientHandler.GetByID)
	router.HandleFunc("PUT /api/v1/client/{id}", clientHandler.Update)
	router.HandleFunc("DELETE /api/v1/client/{id}", clientHandler.Delete)
}
