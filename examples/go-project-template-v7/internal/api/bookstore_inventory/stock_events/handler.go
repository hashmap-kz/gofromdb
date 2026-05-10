package stock_events

import (
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
// @Tags stock-events
// @Accept json
// @Produce json
// @Param request body stockEventsCreateRequest true "Create input"
// @Success 201 {object} stockEventsResponse
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid request payload)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error"
// @Router /api/v1/stock-events [post]
func (h *Handler) Save(w http.ResponseWriter, r *http.Request) {
	req := &stockEventsCreateRequest{}
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

// FindAll
//
// @Summary Get all
// @Description Retrieves a list without pagination.
// @Tags stock-events
// @Accept json
// @Produce  json
// @Success 200 {object} stockEventsResponseList "List of all items"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error"
// @Router /api/v1/stock-events [get]
func (h *Handler) FindAll(w http.ResponseWriter, r *http.Request) {
	resp, err := h.svc.FindAll(r.Context())
	if err != nil {
		httputils.WriteJSON(w, http.StatusInternalServerError, httputils.ErrorResponse{Message: err.Error()})
		return
	}
	httputils.WriteJSON(w, http.StatusOK, stockEventsResponseList{
		Data: mapDtosToPayloads(resp),
	})
}

// FindAllPageable
//
// @Summary Get paginated list
// @Description Retrieves a paginated list using pagination parameters.
// @Tags stock-events
// @Accept json
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param size query int false "Number of items per page (default: 10)"
// @Param sort query string false "Sort order, e.g., 'name,asc'"
// @Success 200 {object} stockEventsResponseList "Paginated list of StockEvents"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid pagination parameters)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error"
// @Router /api/v1/stock-events/pageable [get]
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
	httputils.WriteJSON(w, http.StatusOK, stockEventsResponseList{
		Data: mapDtosToPayloads(resp),
		Page: &page,
	})
}

// mappers

func mapCreateRequestToCreateInputDto(inputRequest *stockEventsCreateRequest) *CreateDto {
	return &CreateDto{
		HappenedAt:    inputRequest.HappenedAt,
		WarehouseCode: inputRequest.WarehouseCode,
		BookID:        inputRequest.BookID,
		DeltaQty:      inputRequest.DeltaQty,
		Reason:        inputRequest.Reason,
		Payload:       inputRequest.Payload,
	}
}

func mapDtosToPayloads(inputDtos []Dto) []stockEventsResponse {
	outputResponses := make([]stockEventsResponse, 0, len(inputDtos))
	for i := range inputDtos {
		outputResponses = append(outputResponses, mapDtoToPayload(&inputDtos[i]))
	}
	return outputResponses
}

func mapDtoToPayload(inputDto *Dto) stockEventsResponse {
	return stockEventsResponse{
		HappenedAt:    inputDto.HappenedAt,
		WarehouseCode: inputDto.WarehouseCode,
		BookID:        inputDto.BookID,
		DeltaQty:      inputDto.DeltaQty,
		Reason:        inputDto.Reason,
		Payload:       inputDto.Payload,
	}
}
