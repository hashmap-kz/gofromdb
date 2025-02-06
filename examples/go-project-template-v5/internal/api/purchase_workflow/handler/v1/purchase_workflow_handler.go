package v1

import (
	"fmt"
	"net/http"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/purchase_workflow/dto"
	"go-project-template-v5/internal/api/purchase_workflow/service"

	"go-project-template-v5/pkg/httputils"
	"go-project-template-v5/pkg/validator"
)

type PurchaseWorkflowHTTPHandler struct {
	purchaseWorkflowService service.PurchaseWorkflowService
}

func NewPurchaseWorkflowHTTPHandler(purchaseWorkflowService service.PurchaseWorkflowService) *PurchaseWorkflowHTTPHandler {
	return &PurchaseWorkflowHTTPHandler{
		purchaseWorkflowService: purchaseWorkflowService,
	}
}

// Save
//
// @Summary Create new item
// @Description Create new item handler
// @Tags purchase-workflow
// @Accept json
// @Produce json
// @Param request body purchaseWorkflowCreateRequest true "Create input"
// @Success 201 {object} purchaseWorkflowResponse
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid request payload)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error"
// @Router /api/v1/purchase-workflow [post]
func (h *PurchaseWorkflowHTTPHandler) Save(w http.ResponseWriter, r *http.Request) {
	// read RequestBody
	req := &purchaseWorkflowCreateRequest{}
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
	resp, err := h.purchaseWorkflowService.Save(r.Context(), createInput)
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
// @Tags purchase-workflow
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Param request body purchaseWorkflowUpdateRequest true "Update input"
// @Success 201 {object} purchaseWorkflowResponse
// @Failure 400 {object} httputils.ErrorResponse "Bad Request"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error"
// @Router /api/v1/purchase-workflow/{record_id} [put]
func (h *PurchaseWorkflowHTTPHandler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	pkRecordID, err := httputils.PathValueI32(r, "record_id")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	req := &purchaseWorkflowUpdateRequest{}
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
	resp, err := h.purchaseWorkflowService.UpdateByID(r.Context(), updateInput, pkRecordID)
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
// @Tags purchase-workflow
// @Accept json
// @Produce json
// @Param id path int true "PurchaseWorkflow ID"
// @Success 204 "No Content (Successfully deleted)"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid ID format)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Deletion failed)"
// @Router /api/v1/purchase-workflow/{record_id} [delete]
func (h *PurchaseWorkflowHTTPHandler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	pkRecordID, err := httputils.PathValueI32(r, "record_id")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	err = h.purchaseWorkflowService.DeleteByID(r.Context(), pkRecordID)
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
// @Tags purchase-workflow
// @Produce json
// @Param id path int true "Item ID"
// @Success 200 {object} purchaseWorkflowResponse "Single item"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid ID format)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Deletion failed)"
// @Router /api/v1/purchase-workflow/{record_id} [get]
func (h *PurchaseWorkflowHTTPHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	pkRecordID, err := httputils.PathValueI32(r, "record_id")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	resp, err := h.purchaseWorkflowService.FindByID(r.Context(), pkRecordID)
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
// @Tags purchase-workflow
// @Accept json
// @Produce  json
// @Success 200 {object} purchaseWorkflowResponseList "List of all items"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Service failure)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Data processing failure)"
// @Router /api/v1/purchase-workflow [get]
func (h *PurchaseWorkflowHTTPHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	// call service
	resp, err := h.purchaseWorkflowService.FindAll(r.Context())
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
	httputils.WriteJSON(w, http.StatusOK, purchaseWorkflowResponseList{
		Data: dtosToPayloads,
	})
}

// FindAllPageable
//
// @Summary Get paginated list
// @Description Retrieves a paginated list using pagination parameters.
// @Tags purchase-workflow
// @Accept json
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param size query int false "Number of items per page (default: 10)"
// @Param sort query string false "Sort order, e.g., 'name,asc'"
// @Success 200 {object} purchaseWorkflowResponseList "Paginated list of PurchaseWorkflow"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid pagination parameters or service failure)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Data processing failure)"
// @Router /api/v1/purchase-workflow/pageable [get]
func (h *PurchaseWorkflowHTTPHandler) FindAllPageable(w http.ResponseWriter, r *http.Request) {
	pq, err := pageable.GetPaginationFromCtx(r)
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// call service
	resp, page, err := h.purchaseWorkflowService.FindAllPageable(r.Context(), pq)
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
	httputils.WriteJSON(w, http.StatusOK, purchaseWorkflowResponseList{
		Data: dtosToPayloads,
		Page: &page,
	})
}

// mappers

func mapCreateRequestToCreateInputDto(inputRequest *purchaseWorkflowCreateRequest) (*dto.PurchaseWorkflowCreateDto, error) {
	if inputRequest == nil {
		return nil, fmt.Errorf("unexpected nil input for mapping between purchaseWorkflowCreateRequest->PurchaseWorkflowCreateDto")
	}
	return &dto.PurchaseWorkflowCreateDto{
		ValidPeriod:    inputRequest.ValidPeriod,
		BuyID:          inputRequest.BuyID,
		PurchaseStepID: inputRequest.PurchaseStepID,
	}, nil
}

func mapUpdateRequestToUpdateInputDto(inputRequest *purchaseWorkflowUpdateRequest) (*dto.PurchaseWorkflowUpdateDto, error) {
	if inputRequest == nil {
		return nil, fmt.Errorf("unexpected nil input for mapping between purchaseWorkflowUpdateRequest->PurchaseWorkflowUpdateDto")
	}
	return &dto.PurchaseWorkflowUpdateDto{
		ValidPeriod:    inputRequest.ValidPeriod,
		BuyID:          inputRequest.BuyID,
		PurchaseStepID: inputRequest.PurchaseStepID,
	}, nil
}

func mapDtosToPayloads(inputDtos []dto.PurchaseWorkflowDto) ([]purchaseWorkflowResponse, error) {
	var outputResponses []purchaseWorkflowResponse
	for _, inputDto := range inputDtos {
		toPayload, err := mapDtoToPayload(&inputDto)
		if err != nil {
			return nil, err
		}
		outputResponses = append(outputResponses, toPayload)
	}
	return outputResponses, nil
}

func mapDtoToPayload(inputDto *dto.PurchaseWorkflowDto) (purchaseWorkflowResponse, error) {
	if inputDto == nil {
		return purchaseWorkflowResponse{}, fmt.Errorf("unexpected nil input for mapping between PurchaseWorkflowDto->purchaseWorkflowResponse")
	}
	return purchaseWorkflowResponse{
		RecordID:       inputDto.RecordID,
		ValidPeriod:    inputDto.ValidPeriod,
		BuyID:          inputDto.BuyID,
		PurchaseStepID: inputDto.PurchaseStepID,
		CreatedAt:      inputDto.CreatedAt,
		UpdatedAt:      inputDto.UpdatedAt,
		Guid:           inputDto.Guid,
	}, nil
}
