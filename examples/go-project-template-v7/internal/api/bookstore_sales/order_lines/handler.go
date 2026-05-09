package order_lines

import (
	"fmt"
	"go-project-template-v7/pkg/httputils"
	"go-project-template-v7/pkg/pageable"
	"go-project-template-v7/pkg/validator"
	"net/http"
)

type Handler struct {
	svc Service
}

func NewHandler(svc Service) *Handler {
	return &Handler{
		svc: svc,
	}
}

// Save
//
// @Summary Create new item
// @Description Create new item handler
// @Tags order-lines
// @Accept json
// @Produce json
// @Param request body orderLinesCreateRequest true "Create input"
// @Success 201 {object} orderLinesResponse
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid request payload)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error"
// @Router /api/v1/order-lines [post]
func (h *Handler) Save(w http.ResponseWriter, r *http.Request) {
	// read RequestBody
	req := &orderLinesCreateRequest{}
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
	resp, err := h.svc.Save(r.Context(), createInput)
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
// @Tags order-lines
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Param request body orderLinesUpdateRequest true "Update input"
// @Success 201 {object} orderLinesResponse
// @Failure 400 {object} httputils.ErrorResponse "Bad Request"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error"
// @Router /api/v1/order-lines/{order_id}/{line_no} [put]
func (h *Handler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	pkOrderID, err := httputils.PathValueI64(r, "order_id")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}
	pkLineNo, err := httputils.PathValueI32(r, "line_no")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	req := &orderLinesUpdateRequest{}
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
	resp, err := h.svc.UpdateByID(r.Context(), updateInput, pkOrderID, pkLineNo)
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
// @Tags order-lines
// @Accept json
// @Produce json
// @Param id path int true "OrderLines ID"
// @Success 204 "No Content (Successfully deleted)"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid ID format)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Deletion failed)"
// @Router /api/v1/order-lines/{order_id}/{line_no} [delete]
func (h *Handler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	pkOrderID, err := httputils.PathValueI64(r, "order_id")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}
	pkLineNo, err := httputils.PathValueI32(r, "line_no")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	err = h.svc.DeleteByID(r.Context(), pkOrderID, pkLineNo)
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
// @Tags order-lines
// @Produce json
// @Param id path int true "Item ID"
// @Success 200 {object} orderLinesResponse "Single item"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid ID format)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Deletion failed)"
// @Router /api/v1/order-lines/{order_id}/{line_no} [get]
func (h *Handler) FindByID(w http.ResponseWriter, r *http.Request) {
	pkOrderID, err := httputils.PathValueI64(r, "order_id")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}
	pkLineNo, err := httputils.PathValueI32(r, "line_no")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	resp, err := h.svc.FindByID(r.Context(), pkOrderID, pkLineNo)
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
// @Tags order-lines
// @Accept json
// @Produce  json
// @Success 200 {object} orderLinesResponseList "List of all items"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Service failure)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Data processing failure)"
// @Router /api/v1/order-lines [get]
func (h *Handler) FindAll(w http.ResponseWriter, r *http.Request) {
	// call service
	resp, err := h.svc.FindAll(r.Context())
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
	httputils.WriteJSON(w, http.StatusOK, orderLinesResponseList{
		Data: dtosToPayloads,
	})
}

// FindAllPageable
//
// @Summary Get paginated list
// @Description Retrieves a paginated list using pagination parameters.
// @Tags order-lines
// @Accept json
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param size query int false "Number of items per page (default: 10)"
// @Param sort query string false "Sort order, e.g., 'name,asc'"
// @Success 200 {object} orderLinesResponseList "Paginated list of OrderLines"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid pagination parameters or service failure)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Data processing failure)"
// @Router /api/v1/order-lines/pageable [get]
func (h *Handler) FindAllPageable(w http.ResponseWriter, r *http.Request) {
	pq, err := pageable.GetPaginationFromCtx(r)
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// call service
	resp, page, err := h.svc.FindAllPageable(r.Context(), pq)
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
	httputils.WriteJSON(w, http.StatusOK, orderLinesResponseList{
		Data: dtosToPayloads,
		Page: &page,
	})
}

// mappers

func mapCreateRequestToCreateInputDto(inputRequest *orderLinesCreateRequest) (*CreateDto, error) {
	if inputRequest == nil {
		return nil, fmt.Errorf("unexpected nil input for mapping between orderLinesCreateRequest->CreateDto")
	}
	return &CreateDto{
		OrderID:        inputRequest.OrderID,
		LineNo:         inputRequest.LineNo,
		BookID:         inputRequest.BookID,
		Quantity:       inputRequest.Quantity,
		UnitPrice:      inputRequest.UnitPrice,
		DiscountAmount: inputRequest.DiscountAmount,
		Note:           inputRequest.Note,
	}, nil
}

func mapUpdateRequestToUpdateInputDto(inputRequest *orderLinesUpdateRequest) (*UpdateDto, error) {
	if inputRequest == nil {
		return nil, fmt.Errorf("unexpected nil input for mapping between orderLinesUpdateRequest->UpdateDto")
	}
	return &UpdateDto{
		BookID:         inputRequest.BookID,
		Quantity:       inputRequest.Quantity,
		UnitPrice:      inputRequest.UnitPrice,
		DiscountAmount: inputRequest.DiscountAmount,
		Note:           inputRequest.Note,
	}, nil
}

func mapDtosToPayloads(inputDtos []Dto) ([]orderLinesResponse, error) {
	outputResponses := make([]orderLinesResponse, 0, len(inputDtos))
	for i := range inputDtos { // Iterate using index to avoid copying (gocritic:rangeValCopy)
		toPayload, err := mapDtoToPayload(&inputDtos[i])
		if err != nil {
			return nil, err
		}
		outputResponses = append(outputResponses, toPayload)
	}
	return outputResponses, nil
}

func mapDtoToPayload(inputDto *Dto) (orderLinesResponse, error) {
	if inputDto == nil {
		return orderLinesResponse{}, fmt.Errorf("unexpected nil input for mapping between Dto->orderLinesResponse")
	}
	return orderLinesResponse{
		OrderID:        inputDto.OrderID,
		LineNo:         inputDto.LineNo,
		BookID:         inputDto.BookID,
		Quantity:       inputDto.Quantity,
		UnitPrice:      inputDto.UnitPrice,
		DiscountAmount: inputDto.DiscountAmount,
		Note:           inputDto.Note,
	}, nil
}
