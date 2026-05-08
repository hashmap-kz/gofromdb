package v1

import (
	"fmt"
	"net/http"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/composite_pk/dto"
	"go-project-template-v5/internal/api/composite_pk/service"

	"go-project-template-v5/pkg/httputils"
	"go-project-template-v5/pkg/validator"
)

type CompositePkHTTPHandler struct {
	compositePkService service.CompositePkService
}

func NewCompositePkHTTPHandler(compositePkService service.CompositePkService) *CompositePkHTTPHandler {
	return &CompositePkHTTPHandler{
		compositePkService: compositePkService,
	}
}

// Save
//
// @Summary Create new item
// @Description Create new item handler
// @Tags composite-pk
// @Accept json
// @Produce json
// @Param request body compositePkCreateRequest true "Create input"
// @Success 201 {object} compositePkResponse
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid request payload)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error"
// @Router /api/v1/composite-pk [post]
func (h *CompositePkHTTPHandler) Save(w http.ResponseWriter, r *http.Request) {
	// read RequestBody
	req := &compositePkCreateRequest{}
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
	resp, err := h.compositePkService.Save(r.Context(), createInput)
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
// @Tags composite-pk
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Param request body compositePkUpdateRequest true "Update input"
// @Success 201 {object} compositePkResponse
// @Failure 400 {object} httputils.ErrorResponse "Bad Request"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error"
// @Router /api/v1/composite-pk/{tenant_id}/{code} [put]
func (h *CompositePkHTTPHandler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	pkTenantID, err := httputils.PathValueI64(r, "tenant_id")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}
	pkCode, err := httputils.PathValueString(r, "code")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	req := &compositePkUpdateRequest{}
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
	resp, err := h.compositePkService.UpdateByID(r.Context(), updateInput, pkTenantID, pkCode)
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
// @Tags composite-pk
// @Accept json
// @Produce json
// @Param id path int true "CompositePk ID"
// @Success 204 "No Content (Successfully deleted)"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid ID format)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Deletion failed)"
// @Router /api/v1/composite-pk/{tenant_id}/{code} [delete]
func (h *CompositePkHTTPHandler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	pkTenantID, err := httputils.PathValueI64(r, "tenant_id")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}
	pkCode, err := httputils.PathValueString(r, "code")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	err = h.compositePkService.DeleteByID(r.Context(), pkTenantID, pkCode)
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
// @Tags composite-pk
// @Produce json
// @Param id path int true "Item ID"
// @Success 200 {object} compositePkResponse "Single item"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid ID format)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Deletion failed)"
// @Router /api/v1/composite-pk/{tenant_id}/{code} [get]
func (h *CompositePkHTTPHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	pkTenantID, err := httputils.PathValueI64(r, "tenant_id")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}
	pkCode, err := httputils.PathValueString(r, "code")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	resp, err := h.compositePkService.FindByID(r.Context(), pkTenantID, pkCode)
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
// @Tags composite-pk
// @Accept json
// @Produce  json
// @Success 200 {object} compositePkResponseList "List of all items"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Service failure)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Data processing failure)"
// @Router /api/v1/composite-pk [get]
func (h *CompositePkHTTPHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	// call service
	resp, err := h.compositePkService.FindAll(r.Context())
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
	httputils.WriteJSON(w, http.StatusOK, compositePkResponseList{
		Data: dtosToPayloads,
	})
}

// FindAllPageable
//
// @Summary Get paginated list
// @Description Retrieves a paginated list using pagination parameters.
// @Tags composite-pk
// @Accept json
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param size query int false "Number of items per page (default: 10)"
// @Param sort query string false "Sort order, e.g., 'name,asc'"
// @Success 200 {object} compositePkResponseList "Paginated list of CompositePk"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid pagination parameters or service failure)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Data processing failure)"
// @Router /api/v1/composite-pk/pageable [get]
func (h *CompositePkHTTPHandler) FindAllPageable(w http.ResponseWriter, r *http.Request) {
	pq, err := pageable.GetPaginationFromCtx(r)
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// call service
	resp, page, err := h.compositePkService.FindAllPageable(r.Context(), pq)
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
	httputils.WriteJSON(w, http.StatusOK, compositePkResponseList{
		Data: dtosToPayloads,
		Page: &page,
	})
}

// mappers

func mapCreateRequestToCreateInputDto(inputRequest *compositePkCreateRequest) (*dto.CompositePkCreateDto, error) {
	if inputRequest == nil {
		return nil, fmt.Errorf("unexpected nil input for mapping between compositePkCreateRequest->CompositePkCreateDto")
	}
	return &dto.CompositePkCreateDto{
		TenantID: inputRequest.TenantID,
		Code:     inputRequest.Code,
		Name:     inputRequest.Name,
	}, nil
}

func mapUpdateRequestToUpdateInputDto(inputRequest *compositePkUpdateRequest) (*dto.CompositePkUpdateDto, error) {
	if inputRequest == nil {
		return nil, fmt.Errorf("unexpected nil input for mapping between compositePkUpdateRequest->CompositePkUpdateDto")
	}
	return &dto.CompositePkUpdateDto{
		Name: inputRequest.Name,
	}, nil
}

func mapDtosToPayloads(inputDtos []dto.CompositePkDto) ([]compositePkResponse, error) {
	outputResponses := make([]compositePkResponse, 0, len(inputDtos))
	for i := range inputDtos { // Iterate using index to avoid copying (gocritic:rangeValCopy)
		toPayload, err := mapDtoToPayload(&inputDtos[i])
		if err != nil {
			return nil, err
		}
		outputResponses = append(outputResponses, toPayload)
	}
	return outputResponses, nil
}

func mapDtoToPayload(inputDto *dto.CompositePkDto) (compositePkResponse, error) {
	if inputDto == nil {
		return compositePkResponse{}, fmt.Errorf("unexpected nil input for mapping between CompositePkDto->compositePkResponse")
	}
	return compositePkResponse{
		TenantID: inputDto.TenantID,
		Code:     inputDto.Code,
		Name:     inputDto.Name,
	}, nil
}
