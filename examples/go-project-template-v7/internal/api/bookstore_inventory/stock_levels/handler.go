package stock_levels

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
// @Tags stock-levels
// @Accept json
// @Produce json
// @Param request body stockLevelsCreateRequest true "Create input"
// @Success 201 {object} stockLevelsResponse
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid request payload)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error"
// @Router /api/v1/stock-levels [post]
func (h *Handler) Save(w http.ResponseWriter, r *http.Request) {
	// read RequestBody
	req := &stockLevelsCreateRequest{}
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
// @Tags stock-levels
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Param request body stockLevelsUpdateRequest true "Update input"
// @Success 201 {object} stockLevelsResponse
// @Failure 400 {object} httputils.ErrorResponse "Bad Request"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error"
// @Router /api/v1/stock-levels/{warehouse_code}/{book_id} [put]
func (h *Handler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	pkWarehouseCode, err := httputils.PathValueString(r, "warehouse_code")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}
	pkBookID, err := httputils.PathValueI64(r, "book_id")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	req := &stockLevelsUpdateRequest{}
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
	resp, err := h.svc.UpdateByID(r.Context(), updateInput, pkWarehouseCode, pkBookID)
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
// @Tags stock-levels
// @Accept json
// @Produce json
// @Param id path int true "StockLevels ID"
// @Success 204 "No Content (Successfully deleted)"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid ID format)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Deletion failed)"
// @Router /api/v1/stock-levels/{warehouse_code}/{book_id} [delete]
func (h *Handler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	pkWarehouseCode, err := httputils.PathValueString(r, "warehouse_code")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}
	pkBookID, err := httputils.PathValueI64(r, "book_id")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	err = h.svc.DeleteByID(r.Context(), pkWarehouseCode, pkBookID)
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
// @Tags stock-levels
// @Produce json
// @Param id path int true "Item ID"
// @Success 200 {object} stockLevelsResponse "Single item"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid ID format)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Deletion failed)"
// @Router /api/v1/stock-levels/{warehouse_code}/{book_id} [get]
func (h *Handler) FindByID(w http.ResponseWriter, r *http.Request) {
	pkWarehouseCode, err := httputils.PathValueString(r, "warehouse_code")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}
	pkBookID, err := httputils.PathValueI64(r, "book_id")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	resp, err := h.svc.FindByID(r.Context(), pkWarehouseCode, pkBookID)
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
// @Tags stock-levels
// @Accept json
// @Produce  json
// @Success 200 {object} stockLevelsResponseList "List of all items"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Service failure)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Data processing failure)"
// @Router /api/v1/stock-levels [get]
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
	httputils.WriteJSON(w, http.StatusOK, stockLevelsResponseList{
		Data: dtosToPayloads,
	})
}

// FindAllPageable
//
// @Summary Get paginated list
// @Description Retrieves a paginated list using pagination parameters.
// @Tags stock-levels
// @Accept json
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param size query int false "Number of items per page (default: 10)"
// @Param sort query string false "Sort order, e.g., 'name,asc'"
// @Success 200 {object} stockLevelsResponseList "Paginated list of StockLevels"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid pagination parameters or service failure)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Data processing failure)"
// @Router /api/v1/stock-levels/pageable [get]
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
	httputils.WriteJSON(w, http.StatusOK, stockLevelsResponseList{
		Data: dtosToPayloads,
		Page: &page,
	})
}

// mappers

func mapCreateRequestToCreateInputDto(inputRequest *stockLevelsCreateRequest) (*CreateDto, error) {
	if inputRequest == nil {
		return nil, fmt.Errorf("unexpected nil input for mapping between stockLevelsCreateRequest->CreateDto")
	}
	return &CreateDto{
		WarehouseCode:    inputRequest.WarehouseCode,
		BookID:           inputRequest.BookID,
		AvailableQty:     inputRequest.AvailableQty,
		ReservedQty:      inputRequest.ReservedQty,
		ReorderThreshold: inputRequest.ReorderThreshold,
		LastCountedAt:    inputRequest.LastCountedAt,
	}, nil
}

func mapUpdateRequestToUpdateInputDto(inputRequest *stockLevelsUpdateRequest) (*UpdateDto, error) {
	if inputRequest == nil {
		return nil, fmt.Errorf("unexpected nil input for mapping between stockLevelsUpdateRequest->UpdateDto")
	}
	return &UpdateDto{
		AvailableQty:     inputRequest.AvailableQty,
		ReservedQty:      inputRequest.ReservedQty,
		ReorderThreshold: inputRequest.ReorderThreshold,
		LastCountedAt:    inputRequest.LastCountedAt,
	}, nil
}

func mapDtosToPayloads(inputDtos []Dto) ([]stockLevelsResponse, error) {
	outputResponses := make([]stockLevelsResponse, 0, len(inputDtos))
	for i := range inputDtos { // Iterate using index to avoid copying (gocritic:rangeValCopy)
		toPayload, err := mapDtoToPayload(&inputDtos[i])
		if err != nil {
			return nil, err
		}
		outputResponses = append(outputResponses, toPayload)
	}
	return outputResponses, nil
}

func mapDtoToPayload(inputDto *Dto) (stockLevelsResponse, error) {
	if inputDto == nil {
		return stockLevelsResponse{}, fmt.Errorf("unexpected nil input for mapping between Dto->stockLevelsResponse")
	}
	return stockLevelsResponse{
		WarehouseCode:    inputDto.WarehouseCode,
		BookID:           inputDto.BookID,
		AvailableQty:     inputDto.AvailableQty,
		ReservedQty:      inputDto.ReservedQty,
		ReorderThreshold: inputDto.ReorderThreshold,
		LastCountedAt:    inputDto.LastCountedAt,
	}, nil
}
