package v1

import (
	"fmt"
	"net/http"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/buy/dto"
	"go-project-template-v5/internal/api/buy/service"

	"go-project-template-v5/pkg/httputils"
	"go-project-template-v5/pkg/validator"
)

type BuyHTTPHandler struct {
	buyService service.BuyService
}

func NewBuyHTTPHandler(buyService service.BuyService) *BuyHTTPHandler {
	return &BuyHTTPHandler{
		buyService: buyService,
	}
}

// Save
//
// @Summary Create new Buy
// @Description Create Buy handler
// @Tags Buy
// @Accept json
// @Produce json
// @Param request body buyCreateRequest true "Create input"
// @Success 201 {object} buyResponse
// @Router /api/v1/buys [post]
func (h *BuyHTTPHandler) Save(w http.ResponseWriter, r *http.Request) {
	// read RequestBody
	req := &buyCreateRequest{}
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
	resp, err := h.buyService.Save(r.Context(), &dto.BuyCreateDto{
		ClientID:    req.ClientID,
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

func (h *BuyHTTPHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	// call service
	resp, err := h.buyService.GetAll(r.Context())
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
	httputils.WriteJSON(w, http.StatusOK, buyResponseList{
		Data: dtosToPayloads,
	})
}

func (h *BuyHTTPHandler) GetAllPaginated(w http.ResponseWriter, r *http.Request) {
	pq, err := pageable.GetPaginationFromCtx(r)
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// call service
	resp, page, err := h.buyService.GetAllPaginated(r.Context(), pq)
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
	httputils.WriteJSON(w, http.StatusOK, buyResponseList{
		Data: dtosToPayloads,
		Page: &page,
	})
}

// Update
//
// @Summary Update existing Buy
// @Description Update Buy handler
// @Tags Buy
// @Accept json
// @Produce json
// @Param request body buyUpdateRequest true "Update input"
// @Success 201 {object} buyResponse
// @Router /api/v1/buys [put]
func (h *BuyHTTPHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := httputils.PathValueI64(r, "id")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	req := &buyUpdateRequest{}
	if err := httputils.ReadJSON(r, &req); err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// TODO: types - int(id)
	resp, err := h.buyService.Update(r.Context(), int(id), &dto.BuyUpdateDto{
		ClientID:    req.ClientID,
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

func (h *BuyHTTPHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := httputils.PathValueI64(r, "id")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	err = h.buyService.Delete(r.Context(), int(id))
	if err != nil {
		httputils.WriteJSON(w, http.StatusInternalServerError, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// 204 OK
	w.WriteHeader(http.StatusNoContent)
}

func (h *BuyHTTPHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := httputils.PathValueI64(r, "id")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	resp, err := h.buyService.GetByID(r.Context(), int(id))
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

func mapDtosToPayloads(inputDtos []dto.BuyDto) ([]buyResponse, error) {
	var outputResponses []buyResponse
	for _, inputDto := range inputDtos {
		toPayload, err := mapDtoToPayload(&inputDto)
		if err != nil {
			return nil, err
		}
		outputResponses = append(outputResponses, toPayload)
	}
	return outputResponses, nil
}

func mapDtoToPayload(inputDto *dto.BuyDto) (buyResponse, error) {
	if inputDto == nil {
		return buyResponse{}, fmt.Errorf("unexpected nil input for mapping between BuyDto->buyResponse")
	}
	return buyResponse{
		RecordID:    inputDto.RecordID,
		ClientID:    inputDto.ClientID,
		Description: inputDto.Description,
		CreatedAt:   inputDto.CreatedAt,
		UpdatedAt:   inputDto.UpdatedAt,
		Guid:        inputDto.Guid,
	}, nil
}
