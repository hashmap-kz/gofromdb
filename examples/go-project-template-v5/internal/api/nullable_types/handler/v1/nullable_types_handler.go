package v1

import (
	"fmt"
	"net/http"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/nullable_types/dto"
	"go-project-template-v5/internal/api/nullable_types/service"

	"go-project-template-v5/pkg/httputils"
	"go-project-template-v5/pkg/validator"
)

type NullableTypesHTTPHandler struct {
	nullableTypesService service.NullableTypesService
}

func NewNullableTypesHTTPHandler(nullableTypesService service.NullableTypesService) *NullableTypesHTTPHandler {
	return &NullableTypesHTTPHandler{
		nullableTypesService: nullableTypesService,
	}
}

// Save
//
// @Summary Create new item
// @Description Create new item handler
// @Tags nullable-types
// @Accept json
// @Produce json
// @Param request body nullableTypesCreateRequest true "Create input"
// @Success 201 {object} nullableTypesResponse
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid request payload)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error"
// @Router /api/v1/nullable-types [post]
func (h *NullableTypesHTTPHandler) Save(w http.ResponseWriter, r *http.Request) {
	// read RequestBody
	req := &nullableTypesCreateRequest{}
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
	resp, err := h.nullableTypesService.Save(r.Context(), createInput)
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
// @Tags nullable-types
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Param request body nullableTypesUpdateRequest true "Update input"
// @Success 201 {object} nullableTypesResponse
// @Failure 400 {object} httputils.ErrorResponse "Bad Request"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error"
// @Router /api/v1/nullable-types/{id} [put]
func (h *NullableTypesHTTPHandler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	pkID, err := httputils.PathValueI64(r, "id")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	req := &nullableTypesUpdateRequest{}
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
	resp, err := h.nullableTypesService.UpdateByID(r.Context(), updateInput, pkID)
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
// @Tags nullable-types
// @Accept json
// @Produce json
// @Param id path int true "NullableTypes ID"
// @Success 204 "No Content (Successfully deleted)"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid ID format)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Deletion failed)"
// @Router /api/v1/nullable-types/{id} [delete]
func (h *NullableTypesHTTPHandler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	pkID, err := httputils.PathValueI64(r, "id")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	err = h.nullableTypesService.DeleteByID(r.Context(), pkID)
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
// @Tags nullable-types
// @Produce json
// @Param id path int true "Item ID"
// @Success 200 {object} nullableTypesResponse "Single item"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid ID format)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Deletion failed)"
// @Router /api/v1/nullable-types/{id} [get]
func (h *NullableTypesHTTPHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	pkID, err := httputils.PathValueI64(r, "id")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	resp, err := h.nullableTypesService.FindByID(r.Context(), pkID)
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
// @Tags nullable-types
// @Accept json
// @Produce  json
// @Success 200 {object} nullableTypesResponseList "List of all items"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Service failure)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Data processing failure)"
// @Router /api/v1/nullable-types [get]
func (h *NullableTypesHTTPHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	// call service
	resp, err := h.nullableTypesService.FindAll(r.Context())
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
	httputils.WriteJSON(w, http.StatusOK, nullableTypesResponseList{
		Data: dtosToPayloads,
	})
}

// FindAllPageable
//
// @Summary Get paginated list
// @Description Retrieves a paginated list using pagination parameters.
// @Tags nullable-types
// @Accept json
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param size query int false "Number of items per page (default: 10)"
// @Param sort query string false "Sort order, e.g., 'name,asc'"
// @Success 200 {object} nullableTypesResponseList "Paginated list of NullableTypes"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid pagination parameters or service failure)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Data processing failure)"
// @Router /api/v1/nullable-types/pageable [get]
func (h *NullableTypesHTTPHandler) FindAllPageable(w http.ResponseWriter, r *http.Request) {
	pq, err := pageable.GetPaginationFromCtx(r)
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// call service
	resp, page, err := h.nullableTypesService.FindAllPageable(r.Context(), pq)
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
	httputils.WriteJSON(w, http.StatusOK, nullableTypesResponseList{
		Data: dtosToPayloads,
		Page: &page,
	})
}

// mappers

func mapCreateRequestToCreateInputDto(inputRequest *nullableTypesCreateRequest) (*dto.NullableTypesCreateDto, error) {
	if inputRequest == nil {
		return nil, fmt.Errorf("unexpected nil input for mapping between nullableTypesCreateRequest->NullableTypesCreateDto")
	}
	return &dto.NullableTypesCreateDto{
		Name:    inputRequest.Name,
		Amount:  inputRequest.Amount,
		Payload: inputRequest.Payload,
		Tags:    inputRequest.Tags,
		Active:  inputRequest.Active,
	}, nil
}

func mapUpdateRequestToUpdateInputDto(inputRequest *nullableTypesUpdateRequest) (*dto.NullableTypesUpdateDto, error) {
	if inputRequest == nil {
		return nil, fmt.Errorf("unexpected nil input for mapping between nullableTypesUpdateRequest->NullableTypesUpdateDto")
	}
	return &dto.NullableTypesUpdateDto{
		Name:    inputRequest.Name,
		Amount:  inputRequest.Amount,
		Payload: inputRequest.Payload,
		Tags:    inputRequest.Tags,
		Active:  inputRequest.Active,
	}, nil
}

func mapDtosToPayloads(inputDtos []dto.NullableTypesDto) ([]nullableTypesResponse, error) {
	outputResponses := make([]nullableTypesResponse, 0, len(inputDtos))
	for i := range inputDtos { // Iterate using index to avoid copying (gocritic:rangeValCopy)
		toPayload, err := mapDtoToPayload(&inputDtos[i])
		if err != nil {
			return nil, err
		}
		outputResponses = append(outputResponses, toPayload)
	}
	return outputResponses, nil
}

func mapDtoToPayload(inputDto *dto.NullableTypesDto) (nullableTypesResponse, error) {
	if inputDto == nil {
		return nullableTypesResponse{}, fmt.Errorf("unexpected nil input for mapping between NullableTypesDto->nullableTypesResponse")
	}
	return nullableTypesResponse{
		ID:        inputDto.ID,
		Name:      inputDto.Name,
		Amount:    inputDto.Amount,
		Payload:   inputDto.Payload,
		Tags:      inputDto.Tags,
		Active:    inputDto.Active,
		CreatedAt: inputDto.CreatedAt,
	}, nil
}
