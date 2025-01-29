package api

import (
	"net/http"

	clientv1 "go-project-template-v5/internal/api/client/handler/v1"
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
	// TODO: other handlers/routes here
	h.initClientRoutes(router)
}

func (h *Handler) initClientRoutes(router *http.ServeMux) {
	clientHandler := clientv1.NewClientHTTPHandler(h.Services.ClientService)
	router.HandleFunc("POST /api/v1/clients", clientHandler.Save)
	router.HandleFunc("GET /api/v1/clients", clientHandler.GetAll)
	router.HandleFunc("GET /api/v1/clients/pageable", clientHandler.GetAllPaginated)
	router.HandleFunc("GET /api/v1/clients/{id}", clientHandler.GetByID)
	router.HandleFunc("PUT /api/v1/clients/{id}", clientHandler.Update)
	router.HandleFunc("DELETE /api/v1/clients/{id}", clientHandler.Delete)
}
