package api

import (
	"go-project-template-v5/internal/api/categories"
	"go-project-template-v5/internal/api/users"
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
	// Categories routes
	categoriesHandler := categories.NewHandler(h.Services.Categories)
	router.HandleFunc("POST /api/v1/categories", categoriesHandler.Save)
	router.HandleFunc("PUT /api/v1/categories/{record_id}", categoriesHandler.UpdateByID)
	router.HandleFunc("DELETE /api/v1/categories/{record_id}", categoriesHandler.DeleteByID)
	router.HandleFunc("GET /api/v1/categories/{record_id}", categoriesHandler.FindByID)
	router.HandleFunc("GET /api/v1/categories", categoriesHandler.FindAll)
	router.HandleFunc("GET /api/v1/categories/pageable", categoriesHandler.FindAllPageable)

	// Users routes
	usersHandler := users.NewHandler(h.Services.Users)
	router.HandleFunc("POST /api/v1/users", usersHandler.Save)
	router.HandleFunc("PUT /api/v1/users/{record_id}", usersHandler.UpdateByID)
	router.HandleFunc("DELETE /api/v1/users/{record_id}", usersHandler.DeleteByID)
	router.HandleFunc("GET /api/v1/users/{record_id}", usersHandler.FindByID)
	router.HandleFunc("GET /api/v1/users", usersHandler.FindAll)
	router.HandleFunc("GET /api/v1/users/pageable", usersHandler.FindAllPageable)
}
