package stock_levels

import (
	"errors"
	"go-project-template-v7/pkg/apperrors"
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
	req := &stockLevelsCreateRequest{}
	if err := httputils.ReadJSON(r, &req); err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}
	if err := validator.ValidateStruct(req); err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}
	createInput := mapCreateRequestToCreateInputDto(req)
	resp, err := h.svc.Save(r.Context(), createInput)
	if err != nil {
		httputils.WriteJSON(w, http.StatusInternalServerError, httputils.ErrorResponse{Message: err.Error()})
		return
	}
	httputils.WriteJSON(w, http.StatusCreated, mapDtoToPayload(resp))
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
// @Success 200 {object} stockLevelsResponse
// @Failure 400 {object} httputils.ErrorResponse "Bad Request"
// @Failure 404 {object} httputils.ErrorResponse "Not Found"
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
	updateInput := mapUpdateRequestToUpdateInputDto(req)
	resp, err := h.svc.UpdateByID(r.Context(), updateInput, pkWarehouseCode, pkBookID)
	if err != nil {
		if errors.Is(err, apperrors.ErrNotFound) {
			httputils.WriteJSON(w, http.StatusNotFound, httputils.ErrorResponse{Message: "record not found"})
			return
		}
		httputils.WriteJSON(w, http.StatusInternalServerError, httputils.ErrorResponse{Message: err.Error()})
		return
	}
	httputils.WriteJSON(w, http.StatusOK, mapDtoToPayload(resp))
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
// @Failure 404 {object} httputils.ErrorResponse "Not Found"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error"
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

	if err := h.svc.DeleteByID(r.Context(), pkWarehouseCode, pkBookID); err != nil {
		if errors.Is(err, apperrors.ErrNotFound) {
			httputils.WriteJSON(w, http.StatusNotFound, httputils.ErrorResponse{Message: "record not found"})
			return
		}
		httputils.WriteJSON(w, http.StatusInternalServerError, httputils.ErrorResponse{Message: err.Error()})
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// FindByID retrieves an item by its ID.
//
// @Summary Get item by ID
// @Description Retrieves the details based on the provided ID in the request path.
// @Tags stock-levels
// @Produce json
// @Param id path int true "Item ID"
// @Success 200 {object} stockLevelsResponse "Single item"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid ID format)"
// @Failure 404 {object} httputils.ErrorResponse "Not Found"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error"
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
		if errors.Is(err, apperrors.ErrNotFound) {
			httputils.WriteJSON(w, http.StatusNotFound, httputils.ErrorResponse{Message: "record not found"})
			return
		}
		httputils.WriteJSON(w, http.StatusInternalServerError, httputils.ErrorResponse{Message: err.Error()})
		return
	}
	httputils.WriteJSON(w, http.StatusOK, mapDtoToPayload(resp))
}

// FindAll
//
// @Summary Get all
// @Description Retrieves a list without pagination.
// @Tags stock-levels
// @Accept json
// @Produce  json
// @Success 200 {object} stockLevelsResponseList "List of all items"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error"
// @Router /api/v1/stock-levels [get]
func (h *Handler) FindAll(w http.ResponseWriter, r *http.Request) {
	resp, err := h.svc.FindAll(r.Context())
	if err != nil {
		httputils.WriteJSON(w, http.StatusInternalServerError, httputils.ErrorResponse{Message: err.Error()})
		return
	}
	httputils.WriteJSON(w, http.StatusOK, stockLevelsResponseList{
		Data: mapDtosToPayloads(resp),
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
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid pagination parameters)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error"
// @Router /api/v1/stock-levels/pageable [get]
func (h *Handler) FindAllPageable(w http.ResponseWriter, r *http.Request) {
	pq, err := pageable.GetPaginationFromCtx(r)
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}
	resp, page, err := h.svc.FindAllPageable(r.Context(), pq)
	if err != nil {
		httputils.WriteJSON(w, http.StatusInternalServerError, httputils.ErrorResponse{Message: err.Error()})
		return
	}
	httputils.WriteJSON(w, http.StatusOK, stockLevelsResponseList{
		Data: mapDtosToPayloads(resp),
		Page: &page,
	})
}

// mappers

func mapCreateRequestToCreateInputDto(inputRequest *stockLevelsCreateRequest) *CreateDto {
	return &CreateDto{
		WarehouseCode:    inputRequest.WarehouseCode,
		BookID:           inputRequest.BookID,
		AvailableQty:     inputRequest.AvailableQty,
		ReservedQty:      inputRequest.ReservedQty,
		ReorderThreshold: inputRequest.ReorderThreshold,
		LastCountedAt:    inputRequest.LastCountedAt,
	}
}

func mapUpdateRequestToUpdateInputDto(inputRequest *stockLevelsUpdateRequest) *UpdateDto {
	return &UpdateDto{
		AvailableQty:     inputRequest.AvailableQty,
		ReservedQty:      inputRequest.ReservedQty,
		ReorderThreshold: inputRequest.ReorderThreshold,
		LastCountedAt:    inputRequest.LastCountedAt,
	}
}

func mapDtosToPayloads(inputDtos []Dto) []stockLevelsResponse {
	outputResponses := make([]stockLevelsResponse, 0, len(inputDtos))
	for i := range inputDtos {
		outputResponses = append(outputResponses, mapDtoToPayload(&inputDtos[i]))
	}
	return outputResponses
}

func mapDtoToPayload(inputDto *Dto) stockLevelsResponse {
	return stockLevelsResponse{
		WarehouseCode:    inputDto.WarehouseCode,
		BookID:           inputDto.BookID,
		AvailableQty:     inputDto.AvailableQty,
		ReservedQty:      inputDto.ReservedQty,
		ReorderThreshold: inputDto.ReorderThreshold,
		LastCountedAt:    inputDto.LastCountedAt,
	}
}
