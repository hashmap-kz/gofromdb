package v1

import (
	"fmt"
	"net/http"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/order_items/dto"
	"go-project-template-v5/internal/api/order_items/service"

	"go-project-template-v5/pkg/httputils"
	"go-project-template-v5/pkg/validator"
)

type OrderItemsHTTPHandler struct {
	orderItemsService service.OrderItemsService
}

func NewOrderItemsHTTPHandler(orderItemsService service.OrderItemsService) *OrderItemsHTTPHandler {
	return &OrderItemsHTTPHandler{
		orderItemsService: orderItemsService,
	}
}

// Save
//
// @Summary Create new item
// @Description Create new item handler
// @Tags order-items
// @Accept json
// @Produce json
// @Param request body orderItemsCreateRequest true "Create input"
// @Success 201 {object} orderItemsResponse
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid request payload)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error"
// @Router /api/v1/order-items [post]
func (h *OrderItemsHTTPHandler) Save(w http.ResponseWriter, r *http.Request) {
	// read RequestBody
	req := &orderItemsCreateRequest{}
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
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// call service
	resp, err := h.orderItemsService.Save(r.Context(), createInput)
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
// @Tags order-items
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Param request body orderItemsUpdateRequest true "Update input"
// @Success 201 {object} orderItemsResponse
// @Failure 400 {object} httputils.ErrorResponse "Bad Request"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error"
// @Router /api/v1/order-items/{record_id} [put]
func (h *OrderItemsHTTPHandler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	pkRecordID, err := httputils.PathValueI32(r, "record_id")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	req := &orderItemsUpdateRequest{}
	if err := httputils.ReadJSON(r, &req); err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}
	if err := validator.ValidateStruct(req); err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// convert handler-request-payload into service-dto
	updateInput, err := mapUpdateRequestToUpdateInputDto(req)
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// call service
	resp, err := h.orderItemsService.UpdateByID(r.Context(), updateInput, pkRecordID)
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
// @Tags order-items
// @Accept json
// @Produce json
// @Param id path int true "OrderItems ID"
// @Success 204 "No Content (Successfully deleted)"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid ID format)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Deletion failed)"
// @Router /api/v1/order-items/{record_id} [delete]
func (h *OrderItemsHTTPHandler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	pkRecordID, err := httputils.PathValueI32(r, "record_id")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	err = h.orderItemsService.DeleteByID(r.Context(), pkRecordID)
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
// @Tags order-items
// @Produce json
// @Param id path int true "Item ID"
// @Success 200 {object} orderItemsResponse "Single item"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid ID format)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Deletion failed)"
// @Router /api/v1/order-items/{record_id} [get]
func (h *OrderItemsHTTPHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	pkRecordID, err := httputils.PathValueI32(r, "record_id")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	resp, err := h.orderItemsService.FindByID(r.Context(), pkRecordID)
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
// @Tags order-items
// @Accept json
// @Produce  json
// @Success 200 {object} orderItemsResponseList "List of all items"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Service failure)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Data processing failure)"
// @Router /api/v1/order-items [get]
func (h *OrderItemsHTTPHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	// call service
	resp, err := h.orderItemsService.FindAll(r.Context())
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
	httputils.WriteJSON(w, http.StatusOK, orderItemsResponseList{
		Data: dtosToPayloads,
	})
}

// FindAllPageable
//
// @Summary Get paginated list
// @Description Retrieves a paginated list using pagination parameters.
// @Tags order-items
// @Accept json
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param size query int false "Number of items per page (default: 10)"
// @Param sort query string false "Sort order, e.g., 'name,asc'"
// @Success 200 {object} orderItemsResponseList "Paginated list of OrderItems"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid pagination parameters or service failure)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Data processing failure)"
// @Router /api/v1/order-items/pageable [get]
func (h *OrderItemsHTTPHandler) FindAllPageable(w http.ResponseWriter, r *http.Request) {
	pq, err := pageable.GetPaginationFromCtx(r)
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// call service
	resp, page, err := h.orderItemsService.FindAllPageable(r.Context(), pq)
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
	httputils.WriteJSON(w, http.StatusOK, orderItemsResponseList{
		Data: dtosToPayloads,
		Page: &page,
	})
}

// mappers

func mapCreateRequestToCreateInputDto(inputRequest *orderItemsCreateRequest) (*dto.OrderItemsCreateDto, error) {
	if inputRequest == nil {
		return nil, fmt.Errorf("unexpected nil input for mapping between orderItemsCreateRequest->OrderItemsCreateDto")
	}
	return &dto.OrderItemsCreateDto{
		OrderID:   inputRequest.OrderID,
		ProductID: inputRequest.ProductID,
		Quantity:  inputRequest.Quantity,
		Price:     inputRequest.Price,
	}, nil
}

func mapUpdateRequestToUpdateInputDto(inputRequest *orderItemsUpdateRequest) (*dto.OrderItemsUpdateDto, error) {
	if inputRequest == nil {
		return nil, fmt.Errorf("unexpected nil input for mapping between orderItemsUpdateRequest->OrderItemsUpdateDto")
	}
	return &dto.OrderItemsUpdateDto{
		OrderID:   inputRequest.OrderID,
		ProductID: inputRequest.ProductID,
		Quantity:  inputRequest.Quantity,
		Price:     inputRequest.Price,
	}, nil
}

func mapDtosToPayloads(inputDtos []dto.OrderItemsDto) ([]orderItemsResponse, error) {
	outputResponses := make([]orderItemsResponse, 0, len(inputDtos))
	for i := range inputDtos { // Iterate using index to avoid copying (gocritic:rangeValCopy)
		toPayload, err := mapDtoToPayload(&inputDtos[i])
		if err != nil {
			return nil, err
		}
		outputResponses = append(outputResponses, toPayload)
	}
	return outputResponses, nil
}

func mapDtoToPayload(inputDto *dto.OrderItemsDto) (orderItemsResponse, error) {
	if inputDto == nil {
		return orderItemsResponse{}, fmt.Errorf("unexpected nil input for mapping between OrderItemsDto->orderItemsResponse")
	}
	return orderItemsResponse{
		RecordID:  inputDto.RecordID,
		OrderID:   inputDto.OrderID,
		ProductID: inputDto.ProductID,
		Quantity:  inputDto.Quantity,
		Price:     inputDto.Price,
		CreatedAt: inputDto.CreatedAt,
		UpdatedAt: inputDto.UpdatedAt,
		GUID:      inputDto.GUID,
	}, nil
}
