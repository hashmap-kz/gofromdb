package api

import (
	"net/http"

	currenciesv1 "go-project-template-v5/internal/api/currencies/handler/v1"
	customersv1 "go-project-template-v5/internal/api/customers/handler/v1"
	productCategoriesv1 "go-project-template-v5/internal/api/product_categories/handler/v1"
	productPricesv1 "go-project-template-v5/internal/api/product_prices/handler/v1"
	productsv1 "go-project-template-v5/internal/api/products/handler/v1"
	purchaseItemsv1 "go-project-template-v5/internal/api/purchase_items/handler/v1"
	purchaseStepsv1 "go-project-template-v5/internal/api/purchase_steps/handler/v1"
	purchaseWorkflowv1 "go-project-template-v5/internal/api/purchase_workflow/handler/v1"
	purchasesv1 "go-project-template-v5/internal/api/purchases/handler/v1"
	usersv1 "go-project-template-v5/internal/api/users/handler/v1"
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
	// Currencies routes
	currenciesHandler := currenciesv1.NewCurrenciesHTTPHandler(h.Services.CurrenciesService)
	router.HandleFunc("POST /api/v1/currencies", currenciesHandler.Save)
	router.HandleFunc("PUT /api/v1/currencies/{record_id}", currenciesHandler.UpdateByID)
	router.HandleFunc("DELETE /api/v1/currencies/{record_id}", currenciesHandler.DeleteByID)
	router.HandleFunc("GET /api/v1/currencies/{record_id}", currenciesHandler.FindByID)
	router.HandleFunc("GET /api/v1/currencies", currenciesHandler.FindAll)
	router.HandleFunc("GET /api/v1/currencies/pageable", currenciesHandler.FindAllPageable)

	// Customers routes
	customersHandler := customersv1.NewCustomersHTTPHandler(h.Services.CustomersService)
	router.HandleFunc("POST /api/v1/customers", customersHandler.Save)
	router.HandleFunc("PUT /api/v1/customers/{record_id}", customersHandler.UpdateByID)
	router.HandleFunc("DELETE /api/v1/customers/{record_id}", customersHandler.DeleteByID)
	router.HandleFunc("GET /api/v1/customers/{record_id}", customersHandler.FindByID)
	router.HandleFunc("GET /api/v1/customers", customersHandler.FindAll)
	router.HandleFunc("GET /api/v1/customers/pageable", customersHandler.FindAllPageable)

	// ProductCategories routes
	productCategoriesHandler := productCategoriesv1.NewProductCategoriesHTTPHandler(h.Services.ProductCategoriesService)
	router.HandleFunc("POST /api/v1/product-categories", productCategoriesHandler.Save)
	router.HandleFunc("PUT /api/v1/product-categories/{record_id}", productCategoriesHandler.UpdateByID)
	router.HandleFunc("DELETE /api/v1/product-categories/{record_id}", productCategoriesHandler.DeleteByID)
	router.HandleFunc("GET /api/v1/product-categories/{record_id}", productCategoriesHandler.FindByID)
	router.HandleFunc("GET /api/v1/product-categories", productCategoriesHandler.FindAll)
	router.HandleFunc("GET /api/v1/product-categories/pageable", productCategoriesHandler.FindAllPageable)

	// ProductPrices routes
	productPricesHandler := productPricesv1.NewProductPricesHTTPHandler(h.Services.ProductPricesService)
	router.HandleFunc("POST /api/v1/product-prices", productPricesHandler.Save)
	router.HandleFunc("PUT /api/v1/product-prices/{record_id}", productPricesHandler.UpdateByID)
	router.HandleFunc("DELETE /api/v1/product-prices/{record_id}", productPricesHandler.DeleteByID)
	router.HandleFunc("GET /api/v1/product-prices/{record_id}", productPricesHandler.FindByID)
	router.HandleFunc("GET /api/v1/product-prices", productPricesHandler.FindAll)
	router.HandleFunc("GET /api/v1/product-prices/pageable", productPricesHandler.FindAllPageable)

	// Products routes
	productsHandler := productsv1.NewProductsHTTPHandler(h.Services.ProductsService)
	router.HandleFunc("POST /api/v1/products", productsHandler.Save)
	router.HandleFunc("PUT /api/v1/products/{record_id}", productsHandler.UpdateByID)
	router.HandleFunc("DELETE /api/v1/products/{record_id}", productsHandler.DeleteByID)
	router.HandleFunc("GET /api/v1/products/{record_id}", productsHandler.FindByID)
	router.HandleFunc("GET /api/v1/products", productsHandler.FindAll)
	router.HandleFunc("GET /api/v1/products/pageable", productsHandler.FindAllPageable)

	// PurchaseItems routes
	purchaseItemsHandler := purchaseItemsv1.NewPurchaseItemsHTTPHandler(h.Services.PurchaseItemsService)
	router.HandleFunc("POST /api/v1/purchase-items", purchaseItemsHandler.Save)
	router.HandleFunc("PUT /api/v1/purchase-items/{record_id}", purchaseItemsHandler.UpdateByID)
	router.HandleFunc("DELETE /api/v1/purchase-items/{record_id}", purchaseItemsHandler.DeleteByID)
	router.HandleFunc("GET /api/v1/purchase-items/{record_id}", purchaseItemsHandler.FindByID)
	router.HandleFunc("GET /api/v1/purchase-items", purchaseItemsHandler.FindAll)
	router.HandleFunc("GET /api/v1/purchase-items/pageable", purchaseItemsHandler.FindAllPageable)

	// PurchaseSteps routes
	purchaseStepsHandler := purchaseStepsv1.NewPurchaseStepsHTTPHandler(h.Services.PurchaseStepsService)
	router.HandleFunc("POST /api/v1/purchase-steps", purchaseStepsHandler.Save)
	router.HandleFunc("PUT /api/v1/purchase-steps/{record_id}", purchaseStepsHandler.UpdateByID)
	router.HandleFunc("DELETE /api/v1/purchase-steps/{record_id}", purchaseStepsHandler.DeleteByID)
	router.HandleFunc("GET /api/v1/purchase-steps/{record_id}", purchaseStepsHandler.FindByID)
	router.HandleFunc("GET /api/v1/purchase-steps", purchaseStepsHandler.FindAll)
	router.HandleFunc("GET /api/v1/purchase-steps/pageable", purchaseStepsHandler.FindAllPageable)

	// PurchaseWorkflow routes
	purchaseWorkflowHandler := purchaseWorkflowv1.NewPurchaseWorkflowHTTPHandler(h.Services.PurchaseWorkflowService)
	router.HandleFunc("POST /api/v1/purchase-workflow", purchaseWorkflowHandler.Save)
	router.HandleFunc("PUT /api/v1/purchase-workflow/{record_id}", purchaseWorkflowHandler.UpdateByID)
	router.HandleFunc("DELETE /api/v1/purchase-workflow/{record_id}", purchaseWorkflowHandler.DeleteByID)
	router.HandleFunc("GET /api/v1/purchase-workflow/{record_id}", purchaseWorkflowHandler.FindByID)
	router.HandleFunc("GET /api/v1/purchase-workflow", purchaseWorkflowHandler.FindAll)
	router.HandleFunc("GET /api/v1/purchase-workflow/pageable", purchaseWorkflowHandler.FindAllPageable)

	// Purchases routes
	purchasesHandler := purchasesv1.NewPurchasesHTTPHandler(h.Services.PurchasesService)
	router.HandleFunc("POST /api/v1/purchases", purchasesHandler.Save)
	router.HandleFunc("PUT /api/v1/purchases/{record_id}", purchasesHandler.UpdateByID)
	router.HandleFunc("DELETE /api/v1/purchases/{record_id}", purchasesHandler.DeleteByID)
	router.HandleFunc("GET /api/v1/purchases/{record_id}", purchasesHandler.FindByID)
	router.HandleFunc("GET /api/v1/purchases", purchasesHandler.FindAll)
	router.HandleFunc("GET /api/v1/purchases/pageable", purchasesHandler.FindAllPageable)

	// Users routes
	usersHandler := usersv1.NewUsersHTTPHandler(h.Services.UsersService)
	router.HandleFunc("POST /api/v1/users", usersHandler.Save)
	router.HandleFunc("PUT /api/v1/users/{record_id}", usersHandler.UpdateByID)
	router.HandleFunc("DELETE /api/v1/users/{record_id}", usersHandler.DeleteByID)
	router.HandleFunc("GET /api/v1/users/{record_id}", usersHandler.FindByID)
	router.HandleFunc("GET /api/v1/users", usersHandler.FindAll)
	router.HandleFunc("GET /api/v1/users/pageable", usersHandler.FindAllPageable)
}
