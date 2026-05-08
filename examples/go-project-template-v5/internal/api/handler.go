package api

import (
	"net/http"

	categoriesv1 "go-project-template-v5/internal/api/categories/handler/v1"
	compositePkv1 "go-project-template-v5/internal/api/composite_pk/handler/v1"
	naturalPkv1 "go-project-template-v5/internal/api/natural_pk/handler/v1"
	noPkv1 "go-project-template-v5/internal/api/no_pk/handler/v1"
	nullableTypesv1 "go-project-template-v5/internal/api/nullable_types/handler/v1"
	orderItemsv1 "go-project-template-v5/internal/api/order_items/handler/v1"
	ordersv1 "go-project-template-v5/internal/api/orders/handler/v1"
	productsv1 "go-project-template-v5/internal/api/products/handler/v1"
	serialPkv1 "go-project-template-v5/internal/api/serial_pk/handler/v1"
	usersv1 "go-project-template-v5/internal/api/users/handler/v1"
	uUIDPkv1 "go-project-template-v5/internal/api/uuid_pk/handler/v1"
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

	// CompositePk routes
	compositePkHandler := compositePkv1.NewCompositePkHTTPHandler(h.Services.CompositePkService)
	router.HandleFunc("POST /api/v1/composite-pk", compositePkHandler.Save)
	router.HandleFunc("PUT /api/v1/composite-pk/{tenant_id}/{code}", compositePkHandler.UpdateByID)
	router.HandleFunc("DELETE /api/v1/composite-pk/{tenant_id}/{code}", compositePkHandler.DeleteByID)
	router.HandleFunc("GET /api/v1/composite-pk/{tenant_id}/{code}", compositePkHandler.FindByID)
	router.HandleFunc("GET /api/v1/composite-pk", compositePkHandler.FindAll)
	router.HandleFunc("GET /api/v1/composite-pk/pageable", compositePkHandler.FindAllPageable)

	// NaturalPk routes
	naturalPkHandler := naturalPkv1.NewNaturalPkHTTPHandler(h.Services.NaturalPkService)
	router.HandleFunc("POST /api/v1/natural-pk", naturalPkHandler.Save)
	router.HandleFunc("PUT /api/v1/natural-pk/{code}", naturalPkHandler.UpdateByID)
	router.HandleFunc("DELETE /api/v1/natural-pk/{code}", naturalPkHandler.DeleteByID)
	router.HandleFunc("GET /api/v1/natural-pk/{code}", naturalPkHandler.FindByID)
	router.HandleFunc("GET /api/v1/natural-pk", naturalPkHandler.FindAll)
	router.HandleFunc("GET /api/v1/natural-pk/pageable", naturalPkHandler.FindAllPageable)

	// NoPk routes
	noPkHandler := noPkv1.NewNoPkHTTPHandler(h.Services.NoPkService)
	router.HandleFunc("POST /api/v1/no-pk", noPkHandler.Save)
	router.HandleFunc("GET /api/v1/no-pk", noPkHandler.FindAll)
	router.HandleFunc("GET /api/v1/no-pk/pageable", noPkHandler.FindAllPageable)

	// NullableTypes routes
	nullableTypesHandler := nullableTypesv1.NewNullableTypesHTTPHandler(h.Services.NullableTypesService)
	router.HandleFunc("POST /api/v1/nullable-types", nullableTypesHandler.Save)
	router.HandleFunc("PUT /api/v1/nullable-types/{id}", nullableTypesHandler.UpdateByID)
	router.HandleFunc("DELETE /api/v1/nullable-types/{id}", nullableTypesHandler.DeleteByID)
	router.HandleFunc("GET /api/v1/nullable-types/{id}", nullableTypesHandler.FindByID)
	router.HandleFunc("GET /api/v1/nullable-types", nullableTypesHandler.FindAll)
	router.HandleFunc("GET /api/v1/nullable-types/pageable", nullableTypesHandler.FindAllPageable)

	// OrderItems routes
	orderItemsHandler := orderItemsv1.NewOrderItemsHTTPHandler(h.Services.OrderItemsService)
	router.HandleFunc("POST /api/v1/order-items", orderItemsHandler.Save)
	router.HandleFunc("PUT /api/v1/order-items/{record_id}", orderItemsHandler.UpdateByID)
	router.HandleFunc("DELETE /api/v1/order-items/{record_id}", orderItemsHandler.DeleteByID)
	router.HandleFunc("GET /api/v1/order-items/{record_id}", orderItemsHandler.FindByID)
	router.HandleFunc("GET /api/v1/order-items", orderItemsHandler.FindAll)
	router.HandleFunc("GET /api/v1/order-items/pageable", orderItemsHandler.FindAllPageable)

	// Orders routes
	ordersHandler := ordersv1.NewOrdersHTTPHandler(h.Services.OrdersService)
	router.HandleFunc("POST /api/v1/orders", ordersHandler.Save)
	router.HandleFunc("PUT /api/v1/orders/{record_id}", ordersHandler.UpdateByID)
	router.HandleFunc("DELETE /api/v1/orders/{record_id}", ordersHandler.DeleteByID)
	router.HandleFunc("GET /api/v1/orders/{record_id}", ordersHandler.FindByID)
	router.HandleFunc("GET /api/v1/orders", ordersHandler.FindAll)
	router.HandleFunc("GET /api/v1/orders/pageable", ordersHandler.FindAllPageable)

	// Products routes
	productsHandler := productsv1.NewProductsHTTPHandler(h.Services.ProductsService)
	router.HandleFunc("POST /api/v1/products", productsHandler.Save)
	router.HandleFunc("PUT /api/v1/products/{record_id}", productsHandler.UpdateByID)
	router.HandleFunc("DELETE /api/v1/products/{record_id}", productsHandler.DeleteByID)
	router.HandleFunc("GET /api/v1/products/{record_id}", productsHandler.FindByID)
	router.HandleFunc("GET /api/v1/products", productsHandler.FindAll)
	router.HandleFunc("GET /api/v1/products/pageable", productsHandler.FindAllPageable)

	// SerialPk routes
	serialPkHandler := serialPkv1.NewSerialPkHTTPHandler(h.Services.SerialPkService)
	router.HandleFunc("POST /api/v1/serial-pk", serialPkHandler.Save)
	router.HandleFunc("PUT /api/v1/serial-pk/{id}", serialPkHandler.UpdateByID)
	router.HandleFunc("DELETE /api/v1/serial-pk/{id}", serialPkHandler.DeleteByID)
	router.HandleFunc("GET /api/v1/serial-pk/{id}", serialPkHandler.FindByID)
	router.HandleFunc("GET /api/v1/serial-pk", serialPkHandler.FindAll)
	router.HandleFunc("GET /api/v1/serial-pk/pageable", serialPkHandler.FindAllPageable)

	// UUIDPk routes
	uUIDPkHandler := uUIDPkv1.NewUUIDPkHTTPHandler(h.Services.UUIDPkService)
	router.HandleFunc("POST /api/v1/uuid-pk", uUIDPkHandler.Save)
	router.HandleFunc("PUT /api/v1/uuid-pk/{id}", uUIDPkHandler.UpdateByID)
	router.HandleFunc("DELETE /api/v1/uuid-pk/{id}", uUIDPkHandler.DeleteByID)
	router.HandleFunc("GET /api/v1/uuid-pk/{id}", uUIDPkHandler.FindByID)
	router.HandleFunc("GET /api/v1/uuid-pk", uUIDPkHandler.FindAll)
	router.HandleFunc("GET /api/v1/uuid-pk/pageable", uUIDPkHandler.FindAllPageable)

	// Users routes
	usersHandler := usersv1.NewUsersHTTPHandler(h.Services.UsersService)
	router.HandleFunc("POST /api/v1/users", usersHandler.Save)
	router.HandleFunc("PUT /api/v1/users/{record_id}", usersHandler.UpdateByID)
	router.HandleFunc("DELETE /api/v1/users/{record_id}", usersHandler.DeleteByID)
	router.HandleFunc("GET /api/v1/users/{record_id}", usersHandler.FindByID)
	router.HandleFunc("GET /api/v1/users", usersHandler.FindAll)
	router.HandleFunc("GET /api/v1/users/pageable", usersHandler.FindAllPageable)
}
