package v1

import (
	"fmt"
	"net/http"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/product/dto"
	"go-project-template-v5/internal/api/product/service"

	"go-project-template-v5/pkg/httputils"
	"go-project-template-v5/pkg/validator"
)

type ProductHTTPHandler struct {
	productService service.ProductService
}

func NewProductHTTPHandler(productService service.ProductService) *ProductHTTPHandler {
	return &ProductHTTPHandler{
		productService: productService,
	}
}

// Save
//
// @Summary Create new Product
// @Description Create Product handler
// @Tags Product
// @Accept json
// @Produce json
// @Param request body productCreateRequest true "Create input"
// @Success 201 {object} productResponse
// @Failure 400 {object} httputils.ErrorResponse "Bad Request"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error"
// @Router /api/v1/products [post]
func (h *ProductHTTPHandler) Save(w http.ResponseWriter, r *http.Request) {
	// read RequestBody
	req := &productCreateRequest{}
	if err := httputils.ReadJSON(r, &req); err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// check RequestBody
	if err := validator.ValidateStruct(req); err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// call service
	resp, err := h.productService.Save(r.Context(), &dto.ProductCreateDto{
		CategoryID:  req.CategoryID,
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		httputils.WriteJSON(w, http.StatusInternalServerError, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// convert service-model to handler-payload
	dtoToPayload, err := mapDtoToPayload(resp)
	if err != nil {
		httputils.WriteJSON(w, http.StatusInternalServerError, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// 200 OK
	httputils.WriteJSON(w, http.StatusOK, dtoToPayload)
}

// GetAll retrieves all Product.
//
// @Summary Get all Product
// @Description Retrieves a list of all Product without pagination.
// @Tags Product
// @Accept json
// @Produce  json
// @Success 200 {object} productResponseList "List of all BuyItems"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Service failure)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Data processing failure)"
// @Router /api/v1/products [get]
func (h *ProductHTTPHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	// call service
	resp, err := h.productService.GetAll(r.Context())
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// convert service-model to handler-payload
	dtosToPayloads, err := mapDtosToPayloads(resp)
	if err != nil {
		httputils.WriteJSON(w, http.StatusInternalServerError, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// 200 OK
	httputils.WriteJSON(w, http.StatusOK, productResponseList{
		Data: dtosToPayloads,
	})
}

// GetAllPaginated retrieves a paginated list of Product.
//
// @Summary Get paginated list of Product
// @Description Retrieves a paginated list of Product with pagination parameters.
// @Tags Product
// @Accept json
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param size query int false "Number of items per page (default: 10)"
// @Param sort query string false "Sort order, e.g., 'name,asc'"
// @Success 200 {object} productResponseList "Paginated list of Product"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid pagination parameters or service failure)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Data processing failure)"
// @Router /api/v1/products/pageable [get]
func (h *ProductHTTPHandler) GetAllPaginated(w http.ResponseWriter, r *http.Request) {
	pq, err := pageable.GetPaginationFromCtx(r)
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// call service
	resp, page, err := h.productService.GetAllPaginated(r.Context(), pq)
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// convert service-model to handler-payload
	dtosToPayloads, err := mapDtosToPayloads(resp)
	if err != nil {
		httputils.WriteJSON(w, http.StatusInternalServerError, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// 200 OK
	httputils.WriteJSON(w, http.StatusOK, productResponseList{
		Data: dtosToPayloads,
		Page: &page,
	})
}

// Update
//
// @Summary Update existing Product
// @Description Update Product handler
// @Tags Product
// @Accept json
// @Produce json
// @Param request body productUpdateRequest true "Update input"
// @Success 201 {object} productResponse
// @Failure 400 {object} httputils.ErrorResponse "Bad Request"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error"
// @Router /api/v1/products [put]
func (h *ProductHTTPHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := httputils.PathValueI64(r, "id")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	req := &productUpdateRequest{}
	if err := httputils.ReadJSON(r, &req); err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// TODO: types - int(id)
	resp, err := h.productService.Update(r.Context(), int(id), &dto.ProductUpdateDto{
		CategoryID:  req.CategoryID,
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		httputils.WriteJSON(w, http.StatusInternalServerError, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// convert service-model to handler-payload
	dtoToPayload, err := mapDtoToPayload(resp)
	if err != nil {
		httputils.WriteJSON(w, http.StatusInternalServerError, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// 201 OK
	httputils.WriteJSON(w, http.StatusCreated, dtoToPayload)
}

// Delete handles the deletion of a Product.
//
// @Summary Delete a Product
// @Description Deletes a Product by its ID
// @Tags Product
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 204 "No Content (Successfully deleted)"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid ID format)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Deletion failed)"
// @Router /api/v1/products [delete]
func (h *ProductHTTPHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := httputils.PathValueI64(r, "id")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	err = h.productService.Delete(r.Context(), int(id))
	if err != nil {
		httputils.WriteJSON(w, http.StatusInternalServerError, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// 204 OK
	w.WriteHeader(http.StatusNoContent)
}

func (h *ProductHTTPHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := httputils.PathValueI64(r, "id")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	resp, err := h.productService.GetByID(r.Context(), int(id))
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	dtoToPayload, err := mapDtoToPayload(resp)
	if err != nil {
		httputils.WriteJSON(w, http.StatusInternalServerError, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	httputils.WriteJSON(w, http.StatusOK, dtoToPayload)
}

// mappers

func mapDtosToPayloads(inputDtos []dto.ProductDto) ([]productResponse, error) {
	var outputResponses []productResponse
	for _, inputDto := range inputDtos {
		toPayload, err := mapDtoToPayload(&inputDto)
		if err != nil {
			return nil, err
		}
		outputResponses = append(outputResponses, toPayload)
	}
	return outputResponses, nil
}

func mapDtoToPayload(inputDto *dto.ProductDto) (productResponse, error) {
	if inputDto == nil {
		return productResponse{}, fmt.Errorf("unexpected nil input for mapping between ProductDto->productResponse")
	}
	return productResponse{
		RecordID:    inputDto.RecordID,
		CategoryID:  inputDto.CategoryID,
		Name:        inputDto.Name,
		Description: inputDto.Description,
		CreatedAt:   inputDto.CreatedAt,
		UpdatedAt:   inputDto.UpdatedAt,
		Guid:        inputDto.Guid,
	}, nil
}
