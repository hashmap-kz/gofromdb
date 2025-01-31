package v1

import (
	"fmt"
	"net/http"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/buy_item/dto"
	"go-project-template-v5/internal/api/buy_item/service"

	"go-project-template-v5/pkg/httputils"
	"go-project-template-v5/pkg/validator"
)

type BuyItemHTTPHandler struct {
	buyItemService service.BuyItemService
}

func NewBuyItemHTTPHandler(buyItemService service.BuyItemService) *BuyItemHTTPHandler {
	return &BuyItemHTTPHandler{
		buyItemService: buyItemService,
	}
}

// Save
//
// @Summary Create new item
// @Description Create new item handler
// @Tags buy-items
// @Accept json
// @Produce json
// @Param request body buyItemCreateRequest true "Create input"
// @Success 201 {object} buyItemResponse
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid request payload)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error"
// @Router /api/v1/buy-items [post]
func (h *BuyItemHTTPHandler) Save(w http.ResponseWriter, r *http.Request) {
	// read RequestBody
	req := &buyItemCreateRequest{}
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
	resp, err := h.buyItemService.Save(r.Context(), createInput)
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

// GetAll
//
// @Summary Get all
// @Description Retrieves a list without pagination.
// @Tags buy-items
// @Accept json
// @Produce  json
// @Success 200 {object} buyItemResponseList "List of all BuyItems"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Service failure)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Data processing failure)"
// @Router /api/v1/buy-items [get]
func (h *BuyItemHTTPHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	// call service
	resp, err := h.buyItemService.GetAll(r.Context())
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
	httputils.WriteJSON(w, http.StatusOK, buyItemResponseList{
		Data: dtosToPayloads,
	})
}

// GetAllPaginated
//
// @Summary Get paginated list
// @Description Retrieves a paginated list using pagination parameters.
// @Tags buy-items
// @Accept json
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param size query int false "Number of items per page (default: 10)"
// @Param sort query string false "Sort order, e.g., 'name,asc'"
// @Success 200 {object} buyItemResponseList "Paginated list of BuyItem"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid pagination parameters or service failure)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Data processing failure)"
// @Router /api/v1/buy-items/pageable [get]
func (h *BuyItemHTTPHandler) GetAllPaginated(w http.ResponseWriter, r *http.Request) {
	pq, err := pageable.GetPaginationFromCtx(r)
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// call service
	resp, page, err := h.buyItemService.GetAllPaginated(r.Context(), pq)
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
	httputils.WriteJSON(w, http.StatusOK, buyItemResponseList{
		Data: dtosToPayloads,
		Page: &page,
	})
}

// Update
//
// @Summary Update existing item
// @Description Updates an item by its ID
// @Tags buy-items
// @Accept json
// @Produce json
// @Param request body buyItemUpdateRequest true "Update input"
// @Success 201 {object} buyItemResponse
// @Failure 400 {object} httputils.ErrorResponse "Bad Request"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error"
// @Router /api/v1/buy-items [put]
func (h *BuyItemHTTPHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := httputils.PathValueI64(r, "id")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	req := &buyItemUpdateRequest{}
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
	// TODO: types - int(id)
	resp, err := h.buyItemService.Update(r.Context(), int(id), updateInput)
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

// Delete
//
// @Summary Delete existing item
// @Description Deletes an item by its ID
// @Tags buy-items
// @Accept json
// @Produce json
// @Param id path int true "BuyItem ID"
// @Success 204 "No Content (Successfully deleted)"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid ID format)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Deletion failed)"
// @Router /api/v1/buy-items [delete]
func (h *BuyItemHTTPHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := httputils.PathValueI64(r, "id")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	err = h.buyItemService.Delete(r.Context(), int(id))
	if err != nil {
		httputils.WriteJSON(w, http.StatusInternalServerError, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// 204 OK
	w.WriteHeader(http.StatusNoContent)
}

func (h *BuyItemHTTPHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := httputils.PathValueI64(r, "id")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	resp, err := h.buyItemService.GetByID(r.Context(), int(id))
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

func mapCreateRequestToCreateInputDto(inputRequest *buyItemCreateRequest) (*dto.BuyItemCreateDto, error) {
	if inputRequest == nil {
		return nil, fmt.Errorf("unexpected nil input for mapping between buyItemCreateRequest->BuyItemCreateDto")
	}
	return &dto.BuyItemCreateDto{
		BuyID:     inputRequest.BuyID,
		ProductID: inputRequest.ProductID,
		Quantity:  inputRequest.Quantity,
		Price:     inputRequest.Price,
	}, nil
}

func mapUpdateRequestToUpdateInputDto(inputRequest *buyItemUpdateRequest) (*dto.BuyItemUpdateDto, error) {
	if inputRequest == nil {
		return nil, fmt.Errorf("unexpected nil input for mapping between buyItemUpdateRequest->BuyItemUpdateDto")
	}
	return &dto.BuyItemUpdateDto{
		BuyID:     inputRequest.BuyID,
		ProductID: inputRequest.ProductID,
		Quantity:  inputRequest.Quantity,
		Price:     inputRequest.Price,
	}, nil
}

func mapDtosToPayloads(inputDtos []dto.BuyItemDto) ([]buyItemResponse, error) {
	var outputResponses []buyItemResponse
	for _, inputDto := range inputDtos {
		toPayload, err := mapDtoToPayload(&inputDto)
		if err != nil {
			return nil, err
		}
		outputResponses = append(outputResponses, toPayload)
	}
	return outputResponses, nil
}

func mapDtoToPayload(inputDto *dto.BuyItemDto) (buyItemResponse, error) {
	if inputDto == nil {
		return buyItemResponse{}, fmt.Errorf("unexpected nil input for mapping between BuyItemDto->buyItemResponse")
	}
	return buyItemResponse{
		RecordID:  inputDto.RecordID,
		BuyID:     inputDto.BuyID,
		ProductID: inputDto.ProductID,
		Quantity:  inputDto.Quantity,
		Price:     inputDto.Price,
		CreatedAt: inputDto.CreatedAt,
		UpdatedAt: inputDto.UpdatedAt,
		Guid:      inputDto.Guid,
	}, nil
}
