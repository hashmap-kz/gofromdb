package v1

import (
	"fmt"
	"net/http"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/purchase_items/dto"
	"go-project-template-v5/internal/api/purchase_items/service"

	"go-project-template-v5/pkg/httputils"
	"go-project-template-v5/pkg/validator"
)

type PurchaseItemsHTTPHandler struct {
	purchaseItemsService service.PurchaseItemsService
}

func NewPurchaseItemsHTTPHandler(purchaseItemsService service.PurchaseItemsService) *PurchaseItemsHTTPHandler {
	return &PurchaseItemsHTTPHandler{
		purchaseItemsService: purchaseItemsService,
	}
}

// Save
//
// @Summary Create new item
// @Description Create new item handler
// @Tags purchase-items
// @Accept json
// @Produce json
// @Param request body purchaseItemsCreateRequest true "Create input"
// @Success 201 {object} purchaseItemsResponse
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid request payload)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error"
// @Router /api/v1/purchase-items [post]
func (h *PurchaseItemsHTTPHandler) Save(w http.ResponseWriter, r *http.Request) {
	// read RequestBody
	req := &purchaseItemsCreateRequest{}
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
	resp, err := h.purchaseItemsService.Save(r.Context(), createInput)
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
// @Tags purchase-items
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Param request body purchaseItemsUpdateRequest true "Update input"
// @Success 201 {object} purchaseItemsResponse
// @Failure 400 {object} httputils.ErrorResponse "Bad Request"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error"
// @Router /api/v1/purchase-items/{record_id} [put]
func (h *PurchaseItemsHTTPHandler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	pkRecordID, err := httputils.PathValueI32(r, "record_id")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	req := &purchaseItemsUpdateRequest{}
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
	resp, err := h.purchaseItemsService.UpdateByID(r.Context(), updateInput, pkRecordID)
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
// @Tags purchase-items
// @Accept json
// @Produce json
// @Param id path int true "PurchaseItems ID"
// @Success 204 "No Content (Successfully deleted)"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid ID format)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Deletion failed)"
// @Router /api/v1/purchase-items/{record_id} [delete]
func (h *PurchaseItemsHTTPHandler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	pkRecordID, err := httputils.PathValueI32(r, "record_id")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	err = h.purchaseItemsService.DeleteByID(r.Context(), pkRecordID)
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
// @Tags purchase-items
// @Produce json
// @Param id path int true "Item ID"
// @Success 200 {object} purchaseItemsResponse "Single item"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid ID format)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Deletion failed)"
// @Router /api/v1/purchase-items/{record_id} [get]
func (h *PurchaseItemsHTTPHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	pkRecordID, err := httputils.PathValueI32(r, "record_id")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	resp, err := h.purchaseItemsService.FindByID(r.Context(), pkRecordID)
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
// @Tags purchase-items
// @Accept json
// @Produce  json
// @Success 200 {object} purchaseItemsResponseList "List of all items"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Service failure)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Data processing failure)"
// @Router /api/v1/purchase-items [get]
func (h *PurchaseItemsHTTPHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	// call service
	resp, err := h.purchaseItemsService.FindAll(r.Context())
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
	httputils.WriteJSON(w, http.StatusOK, purchaseItemsResponseList{
		Data: dtosToPayloads,
	})
}

// FindAllPageable
//
// @Summary Get paginated list
// @Description Retrieves a paginated list using pagination parameters.
// @Tags purchase-items
// @Accept json
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param size query int false "Number of items per page (default: 10)"
// @Param sort query string false "Sort order, e.g., 'name,asc'"
// @Success 200 {object} purchaseItemsResponseList "Paginated list of PurchaseItems"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid pagination parameters or service failure)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Data processing failure)"
// @Router /api/v1/purchase-items/pageable [get]
func (h *PurchaseItemsHTTPHandler) FindAllPageable(w http.ResponseWriter, r *http.Request) {
	pq, err := pageable.GetPaginationFromCtx(r)
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// call service
	resp, page, err := h.purchaseItemsService.FindAllPageable(r.Context(), pq)
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
	httputils.WriteJSON(w, http.StatusOK, purchaseItemsResponseList{
		Data: dtosToPayloads,
		Page: &page,
	})
}

// mappers

func mapCreateRequestToCreateInputDto(inputRequest *purchaseItemsCreateRequest) (*dto.PurchaseItemsCreateDto, error) {
	if inputRequest == nil {
		return nil, fmt.Errorf("unexpected nil input for mapping between purchaseItemsCreateRequest->PurchaseItemsCreateDto")
	}
	return &dto.PurchaseItemsCreateDto{
		PurchaseID: inputRequest.PurchaseID,
		ProductID:  inputRequest.ProductID,
		Quantity:   inputRequest.Quantity,
		Price:      inputRequest.Price,
	}, nil
}

func mapUpdateRequestToUpdateInputDto(inputRequest *purchaseItemsUpdateRequest) (*dto.PurchaseItemsUpdateDto, error) {
	if inputRequest == nil {
		return nil, fmt.Errorf("unexpected nil input for mapping between purchaseItemsUpdateRequest->PurchaseItemsUpdateDto")
	}
	return &dto.PurchaseItemsUpdateDto{
		PurchaseID: inputRequest.PurchaseID,
		ProductID:  inputRequest.ProductID,
		Quantity:   inputRequest.Quantity,
		Price:      inputRequest.Price,
	}, nil
}

func mapDtosToPayloads(inputDtos []dto.PurchaseItemsDto) ([]purchaseItemsResponse, error) {
	var outputResponses []purchaseItemsResponse
	for _, inputDto := range inputDtos {
		toPayload, err := mapDtoToPayload(&inputDto)
		if err != nil {
			return nil, err
		}
		outputResponses = append(outputResponses, toPayload)
	}
	return outputResponses, nil
}

func mapDtoToPayload(inputDto *dto.PurchaseItemsDto) (purchaseItemsResponse, error) {
	if inputDto == nil {
		return purchaseItemsResponse{}, fmt.Errorf("unexpected nil input for mapping between PurchaseItemsDto->purchaseItemsResponse")
	}
	return purchaseItemsResponse{
		RecordID:   inputDto.RecordID,
		PurchaseID: inputDto.PurchaseID,
		ProductID:  inputDto.ProductID,
		Quantity:   inputDto.Quantity,
		Price:      inputDto.Price,
		CreatedAt:  inputDto.CreatedAt,
		UpdatedAt:  inputDto.UpdatedAt,
		Guid:       inputDto.Guid,
	}, nil
}
