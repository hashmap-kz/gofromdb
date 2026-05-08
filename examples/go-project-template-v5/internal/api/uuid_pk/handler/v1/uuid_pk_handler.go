package v1

import (
	"fmt"
	"net/http"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/uuid_pk/dto"
	"go-project-template-v5/internal/api/uuid_pk/service"

	"go-project-template-v5/pkg/httputils"
	"go-project-template-v5/pkg/validator"
)

type UUIDPkHTTPHandler struct {
	uUIDPkService service.UUIDPkService
}

func NewUUIDPkHTTPHandler(uUIDPkService service.UUIDPkService) *UUIDPkHTTPHandler {
	return &UUIDPkHTTPHandler{
		uUIDPkService: uUIDPkService,
	}
}

// Save
//
// @Summary Create new item
// @Description Create new item handler
// @Tags uuid-pk
// @Accept json
// @Produce json
// @Param request body uUIDPkCreateRequest true "Create input"
// @Success 201 {object} uUIDPkResponse
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid request payload)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error"
// @Router /api/v1/uuid-pk [post]
func (h *UUIDPkHTTPHandler) Save(w http.ResponseWriter, r *http.Request) {
	// read RequestBody
	req := &uUIDPkCreateRequest{}
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
	resp, err := h.uUIDPkService.Save(r.Context(), createInput)
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
// @Tags uuid-pk
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Param request body uUIDPkUpdateRequest true "Update input"
// @Success 201 {object} uUIDPkResponse
// @Failure 400 {object} httputils.ErrorResponse "Bad Request"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error"
// @Router /api/v1/uuid-pk/{id} [put]
func (h *UUIDPkHTTPHandler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	pkID, err := httputils.PathValueString(r, "id")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	req := &uUIDPkUpdateRequest{}
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
	resp, err := h.uUIDPkService.UpdateByID(r.Context(), updateInput, pkID)
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
// @Tags uuid-pk
// @Accept json
// @Produce json
// @Param id path int true "UUIDPk ID"
// @Success 204 "No Content (Successfully deleted)"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid ID format)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Deletion failed)"
// @Router /api/v1/uuid-pk/{id} [delete]
func (h *UUIDPkHTTPHandler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	pkID, err := httputils.PathValueString(r, "id")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	err = h.uUIDPkService.DeleteByID(r.Context(), pkID)
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
// @Tags uuid-pk
// @Produce json
// @Param id path int true "Item ID"
// @Success 200 {object} uUIDPkResponse "Single item"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid ID format)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Deletion failed)"
// @Router /api/v1/uuid-pk/{id} [get]
func (h *UUIDPkHTTPHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	pkID, err := httputils.PathValueString(r, "id")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	resp, err := h.uUIDPkService.FindByID(r.Context(), pkID)
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
// @Tags uuid-pk
// @Accept json
// @Produce  json
// @Success 200 {object} uUIDPkResponseList "List of all items"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Service failure)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Data processing failure)"
// @Router /api/v1/uuid-pk [get]
func (h *UUIDPkHTTPHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	// call service
	resp, err := h.uUIDPkService.FindAll(r.Context())
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
	httputils.WriteJSON(w, http.StatusOK, uUIDPkResponseList{
		Data: dtosToPayloads,
	})
}

// FindAllPageable
//
// @Summary Get paginated list
// @Description Retrieves a paginated list using pagination parameters.
// @Tags uuid-pk
// @Accept json
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param size query int false "Number of items per page (default: 10)"
// @Param sort query string false "Sort order, e.g., 'name,asc'"
// @Success 200 {object} uUIDPkResponseList "Paginated list of UUIDPk"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid pagination parameters or service failure)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Data processing failure)"
// @Router /api/v1/uuid-pk/pageable [get]
func (h *UUIDPkHTTPHandler) FindAllPageable(w http.ResponseWriter, r *http.Request) {
	pq, err := pageable.GetPaginationFromCtx(r)
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// call service
	resp, page, err := h.uUIDPkService.FindAllPageable(r.Context(), pq)
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
	httputils.WriteJSON(w, http.StatusOK, uUIDPkResponseList{
		Data: dtosToPayloads,
		Page: &page,
	})
}

// mappers

func mapCreateRequestToCreateInputDto(inputRequest *uUIDPkCreateRequest) (*dto.UUIDPkCreateDto, error) {
	if inputRequest == nil {
		return nil, fmt.Errorf("unexpected nil input for mapping between uUIDPkCreateRequest->UUIDPkCreateDto")
	}
	return &dto.UUIDPkCreateDto{
		ID:   inputRequest.ID,
		Name: inputRequest.Name,
	}, nil
}

func mapUpdateRequestToUpdateInputDto(inputRequest *uUIDPkUpdateRequest) (*dto.UUIDPkUpdateDto, error) {
	if inputRequest == nil {
		return nil, fmt.Errorf("unexpected nil input for mapping between uUIDPkUpdateRequest->UUIDPkUpdateDto")
	}
	return &dto.UUIDPkUpdateDto{
		Name: inputRequest.Name,
	}, nil
}

func mapDtosToPayloads(inputDtos []dto.UUIDPkDto) ([]uUIDPkResponse, error) {
	outputResponses := make([]uUIDPkResponse, 0, len(inputDtos))
	for i := range inputDtos { // Iterate using index to avoid copying (gocritic:rangeValCopy)
		toPayload, err := mapDtoToPayload(&inputDtos[i])
		if err != nil {
			return nil, err
		}
		outputResponses = append(outputResponses, toPayload)
	}
	return outputResponses, nil
}

func mapDtoToPayload(inputDto *dto.UUIDPkDto) (uUIDPkResponse, error) {
	if inputDto == nil {
		return uUIDPkResponse{}, fmt.Errorf("unexpected nil input for mapping between UUIDPkDto->uUIDPkResponse")
	}
	return uUIDPkResponse{
		ID:   inputDto.ID,
		Name: inputDto.Name,
	}, nil
}
