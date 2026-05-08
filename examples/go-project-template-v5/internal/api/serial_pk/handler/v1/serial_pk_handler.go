package v1

import (
	"fmt"
	"net/http"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/serial_pk/dto"
	"go-project-template-v5/internal/api/serial_pk/service"

	"go-project-template-v5/pkg/httputils"
	"go-project-template-v5/pkg/validator"
)

type SerialPkHTTPHandler struct {
	serialPkService service.SerialPkService
}

func NewSerialPkHTTPHandler(serialPkService service.SerialPkService) *SerialPkHTTPHandler {
	return &SerialPkHTTPHandler{
		serialPkService: serialPkService,
	}
}

// Save
//
// @Summary Create new item
// @Description Create new item handler
// @Tags serial-pk
// @Accept json
// @Produce json
// @Param request body serialPkCreateRequest true "Create input"
// @Success 201 {object} serialPkResponse
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid request payload)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error"
// @Router /api/v1/serial-pk [post]
func (h *SerialPkHTTPHandler) Save(w http.ResponseWriter, r *http.Request) {
	// read RequestBody
	req := &serialPkCreateRequest{}
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
	resp, err := h.serialPkService.Save(r.Context(), createInput)
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
// @Tags serial-pk
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Param request body serialPkUpdateRequest true "Update input"
// @Success 201 {object} serialPkResponse
// @Failure 400 {object} httputils.ErrorResponse "Bad Request"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error"
// @Router /api/v1/serial-pk/{id} [put]
func (h *SerialPkHTTPHandler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	pkID, err := httputils.PathValueI64(r, "id")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	req := &serialPkUpdateRequest{}
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
	resp, err := h.serialPkService.UpdateByID(r.Context(), updateInput, pkID)
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
// @Tags serial-pk
// @Accept json
// @Produce json
// @Param id path int true "SerialPk ID"
// @Success 204 "No Content (Successfully deleted)"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid ID format)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Deletion failed)"
// @Router /api/v1/serial-pk/{id} [delete]
func (h *SerialPkHTTPHandler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	pkID, err := httputils.PathValueI64(r, "id")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	err = h.serialPkService.DeleteByID(r.Context(), pkID)
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
// @Tags serial-pk
// @Produce json
// @Param id path int true "Item ID"
// @Success 200 {object} serialPkResponse "Single item"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid ID format)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Deletion failed)"
// @Router /api/v1/serial-pk/{id} [get]
func (h *SerialPkHTTPHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	pkID, err := httputils.PathValueI64(r, "id")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	resp, err := h.serialPkService.FindByID(r.Context(), pkID)
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
// @Tags serial-pk
// @Accept json
// @Produce  json
// @Success 200 {object} serialPkResponseList "List of all items"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Service failure)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Data processing failure)"
// @Router /api/v1/serial-pk [get]
func (h *SerialPkHTTPHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	// call service
	resp, err := h.serialPkService.FindAll(r.Context())
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
	httputils.WriteJSON(w, http.StatusOK, serialPkResponseList{
		Data: dtosToPayloads,
	})
}

// FindAllPageable
//
// @Summary Get paginated list
// @Description Retrieves a paginated list using pagination parameters.
// @Tags serial-pk
// @Accept json
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param size query int false "Number of items per page (default: 10)"
// @Param sort query string false "Sort order, e.g., 'name,asc'"
// @Success 200 {object} serialPkResponseList "Paginated list of SerialPk"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid pagination parameters or service failure)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Data processing failure)"
// @Router /api/v1/serial-pk/pageable [get]
func (h *SerialPkHTTPHandler) FindAllPageable(w http.ResponseWriter, r *http.Request) {
	pq, err := pageable.GetPaginationFromCtx(r)
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// call service
	resp, page, err := h.serialPkService.FindAllPageable(r.Context(), pq)
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
	httputils.WriteJSON(w, http.StatusOK, serialPkResponseList{
		Data: dtosToPayloads,
		Page: &page,
	})
}

// mappers

func mapCreateRequestToCreateInputDto(inputRequest *serialPkCreateRequest) (*dto.SerialPkCreateDto, error) {
	if inputRequest == nil {
		return nil, fmt.Errorf("unexpected nil input for mapping between serialPkCreateRequest->SerialPkCreateDto")
	}
	return &dto.SerialPkCreateDto{
		Name: inputRequest.Name,
	}, nil
}

func mapUpdateRequestToUpdateInputDto(inputRequest *serialPkUpdateRequest) (*dto.SerialPkUpdateDto, error) {
	if inputRequest == nil {
		return nil, fmt.Errorf("unexpected nil input for mapping between serialPkUpdateRequest->SerialPkUpdateDto")
	}
	return &dto.SerialPkUpdateDto{
		Name: inputRequest.Name,
	}, nil
}

func mapDtosToPayloads(inputDtos []dto.SerialPkDto) ([]serialPkResponse, error) {
	outputResponses := make([]serialPkResponse, 0, len(inputDtos))
	for i := range inputDtos { // Iterate using index to avoid copying (gocritic:rangeValCopy)
		toPayload, err := mapDtoToPayload(&inputDtos[i])
		if err != nil {
			return nil, err
		}
		outputResponses = append(outputResponses, toPayload)
	}
	return outputResponses, nil
}

func mapDtoToPayload(inputDto *dto.SerialPkDto) (serialPkResponse, error) {
	if inputDto == nil {
		return serialPkResponse{}, fmt.Errorf("unexpected nil input for mapping between SerialPkDto->serialPkResponse")
	}
	return serialPkResponse{
		ID:   inputDto.ID,
		Name: inputDto.Name,
	}, nil
}
