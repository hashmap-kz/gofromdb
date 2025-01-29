package http

import (
	"context"
	"net/http"

	"go-project-template-v5/pkg/httputils"

	httpSwagger "github.com/swaggo/http-swagger"
	// for swagger
	_ "go-project-template-v5/docs"
)

func InitRouter(_ context.Context) *http.ServeMux {
	router := http.NewServeMux()

	// docs

	router.Handle("/swagger-ui/", httpSwagger.WrapHandler)
	router.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./docs/swagger.json")
	})

	// internal

	router.HandleFunc("GET /healthz", func(w http.ResponseWriter, _ *http.Request) {
		httputils.WriteJSON(w, http.StatusOK, map[string]string{
			"status": "UP",
		})
	})

	return router
}
