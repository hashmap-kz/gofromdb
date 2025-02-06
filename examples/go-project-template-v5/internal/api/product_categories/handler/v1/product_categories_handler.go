package v1

import (
	"fmt"
	"net/http"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/product_categories/dto"
	"go-project-template-v5/internal/api/product_categories/service"

	"go-project-template-v5/pkg/httputils"
	"go-project-template-v5/pkg/validator"
)

type ProductCategoriesHTTPHandler struct {
	productCategoriesService service.ProductCategoriesService
}

func NewProductCategoriesHTTPHandler(productCategoriesService service.ProductCategoriesService) *ProductCategoriesHTTPHandler {
	return &ProductCategoriesHTTPHandler{
		productCategoriesService: productCategoriesService,
	}
}

// Save
//
// @Summary Create new item
// @Description Create new item handler
// @Tags product-categories
// @Accept json
// @Produce json
// @Param request body productCategoriesCreateRequest true "Create input"
// @Success 201 {object} productCategoriesResponse
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid request payload)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error"
// @Router /api/v1/product-categories [post]
func (h *ProductCategoriesHTTPHandler) Save(w http.ResponseWriter, r *http.Request) {
	// read RequestBody
	req := &productCategoriesCreateRequest{}
	if err := httputils.ReadJSON(r, &req); err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// check RequestBody
	if err := validator.ValidateStruct(req); err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// convert handler-request-payload into service-dto
	createInput, err := mapCreateRequestToCreateInputDto(req)
	if err := validator.ValidateStruct(req); err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// call service
	resp, err := h.productCategoriesService.Save(r.Context(), createInput)
	if err != nil {
		httputils.WriteJSON(w, http.StatusInternalServerError, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// convert service-dto into handler-response-payload
	dtoToPayload, err := mapDtoToPayload(resp)
	if err != nil {
		httputils.WriteJSON(w, http.StatusInternalServerError, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// 200 OK
	httputils.WriteJSON(w, http.StatusOK, dtoToPayload)
}

// UpdateByID
//
// @Summary Update existing item
// @Description Updates an item by its ID
// @Tags product-categories
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Param request body productCategoriesUpdateRequest true "Update input"
// @Success 201 {object} productCategoriesResponse
// @Failure 400 {object} httputils.ErrorResponse "Bad Request"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error"
// @Router /api/v1/product-categories/{record_id} [put]
func (h *ProductCategoriesHTTPHandler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	pkRecordID, err := httputils.PathValueI32(r, "record_id")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	req := &productCategoriesUpdateRequest{}
	if err := httputils.ReadJSON(r, &req); err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// convert handler-request-payload into service-dto
	updateInput, err := mapUpdateRequestToUpdateInputDto(req)
	if err := validator.ValidateStruct(req); err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// call service
	resp, err := h.productCategoriesService.UpdateByID(r.Context(), updateInput, pkRecordID)
	if err != nil {
		httputils.WriteJSON(w, http.StatusInternalServerError, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// convert service-dto into handler-response-payload
	dtoToPayload, err := mapDtoToPayload(resp)
	if err != nil {
		httputils.WriteJSON(w, http.StatusInternalServerError, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// 201 OK
	httputils.WriteJSON(w, http.StatusCreated, dtoToPayload)
}

// DeleteByID
//
// @Summary Delete existing item
// @Description Deletes an item by its ID
// @Tags product-categories
// @Accept json
// @Produce json
// @Param id path int true "ProductCategories ID"
// @Success 204 "No Content (Successfully deleted)"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid ID format)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Deletion failed)"
// @Router /api/v1/product-categories/{record_id} [delete]
func (h *ProductCategoriesHTTPHandler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	pkRecordID, err := httputils.PathValueI32(r, "record_id")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	err = h.productCategoriesService.DeleteByID(r.Context(), pkRecordID)
	if err != nil {
		httputils.WriteJSON(w, http.StatusInternalServerError, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// 204 OK
	w.WriteHeader(http.StatusNoContent)
}

// FindByID retrieves a purchase by its ID.
//
// @Summary Get item by ID
// @Description Retrieves the details based on the provided ID in the request path.
// @Tags product-categories
// @Produce json
// @Param id path int true "Item ID"
// @Success 200 {object} productCategoriesResponse "Single item"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid ID format)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Deletion failed)"
// @Router /api/v1/product-categories/{record_id} [get]
func (h *ProductCategoriesHTTPHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	pkRecordID, err := httputils.PathValueI32(r, "record_id")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	resp, err := h.productCategoriesService.FindByID(r.Context(), pkRecordID)
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

// FindAll
//
// @Summary Get all
// @Description Retrieves a list without pagination.
// @Tags product-categories
// @Accept json
// @Produce  json
// @Success 200 {object} productCategoriesResponseList "List of all items"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Service failure)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Data processing failure)"
// @Router /api/v1/product-categories [get]
func (h *ProductCategoriesHTTPHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	// call service
	resp, err := h.productCategoriesService.FindAll(r.Context())
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
	httputils.WriteJSON(w, http.StatusOK, productCategoriesResponseList{
		Data: dtosToPayloads,
	})
}

// FindAllPageable
//
// @Summary Get paginated list
// @Description Retrieves a paginated list using pagination parameters.
// @Tags product-categories
// @Accept json
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param size query int false "Number of items per page (default: 10)"
// @Param sort query string false "Sort order, e.g., 'name,asc'"
// @Success 200 {object} productCategoriesResponseList "Paginated list of ProductCategories"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid pagination parameters or service failure)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Data processing failure)"
// @Router /api/v1/product-categories/pageable [get]
func (h *ProductCategoriesHTTPHandler) FindAllPageable(w http.ResponseWriter, r *http.Request) {
	pq, err := pageable.GetPaginationFromCtx(r)
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// call service
	resp, page, err := h.productCategoriesService.FindAllPageable(r.Context(), pq)
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
	httputils.WriteJSON(w, http.StatusOK, productCategoriesResponseList{
		Data: dtosToPayloads,
		Page: &page,
	})
}

// mappers

func mapCreateRequestToCreateInputDto(inputRequest *productCategoriesCreateRequest) (*dto.ProductCategoriesCreateDto, error) {
	if inputRequest == nil {
		return nil, fmt.Errorf("unexpected nil input for mapping between productCategoriesCreateRequest->ProductCategoriesCreateDto")
	}
	return &dto.ProductCategoriesCreateDto{
		Name:        inputRequest.Name,
		ParentID:    inputRequest.ParentID,
		ValidPeriod: inputRequest.ValidPeriod,
	}, nil
}

func mapUpdateRequestToUpdateInputDto(inputRequest *productCategoriesUpdateRequest) (*dto.ProductCategoriesUpdateDto, error) {
	if inputRequest == nil {
		return nil, fmt.Errorf("unexpected nil input for mapping between productCategoriesUpdateRequest->ProductCategoriesUpdateDto")
	}
	return &dto.ProductCategoriesUpdateDto{
		Name:        inputRequest.Name,
		ParentID:    inputRequest.ParentID,
		ValidPeriod: inputRequest.ValidPeriod,
	}, nil
}

func mapDtosToPayloads(inputDtos []dto.ProductCategoriesDto) ([]productCategoriesResponse, error) {
	var outputResponses []productCategoriesResponse
	for _, inputDto := range inputDtos {
		toPayload, err := mapDtoToPayload(&inputDto)
		if err != nil {
			return nil, err
		}
		outputResponses = append(outputResponses, toPayload)
	}
	return outputResponses, nil
}

func mapDtoToPayload(inputDto *dto.ProductCategoriesDto) (productCategoriesResponse, error) {
	if inputDto == nil {
		return productCategoriesResponse{}, fmt.Errorf("unexpected nil input for mapping between ProductCategoriesDto->productCategoriesResponse")
	}
	return productCategoriesResponse{
		RecordID:    inputDto.RecordID,
		Name:        inputDto.Name,
		ParentID:    inputDto.ParentID,
		ValidPeriod: inputDto.ValidPeriod,
		IsCurrent:   inputDto.IsCurrent,
		CreatedAt:   inputDto.CreatedAt,
		UpdatedAt:   inputDto.UpdatedAt,
		Guid:        inputDto.Guid,
	}, nil
}
