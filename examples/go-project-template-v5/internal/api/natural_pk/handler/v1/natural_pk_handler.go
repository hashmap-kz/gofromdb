package v1

import (
	"fmt"
	"net/http"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/natural_pk/dto"
	"go-project-template-v5/internal/api/natural_pk/service"

	"go-project-template-v5/pkg/httputils"
	"go-project-template-v5/pkg/validator"
)

type NaturalPkHTTPHandler struct {
	naturalPkService service.NaturalPkService
}

func NewNaturalPkHTTPHandler(naturalPkService service.NaturalPkService) *NaturalPkHTTPHandler {
	return &NaturalPkHTTPHandler{
		naturalPkService: naturalPkService,
	}
}

// Save
//
// @Summary Create new item
// @Description Create new item handler
// @Tags natural-pk
// @Accept json
// @Produce json
// @Param request body naturalPkCreateRequest true "Create input"
// @Success 201 {object} naturalPkResponse
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid request payload)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error"
// @Router /api/v1/natural-pk [post]
func (h *NaturalPkHTTPHandler) Save(w http.ResponseWriter, r *http.Request) {
	// read RequestBody
	req := &naturalPkCreateRequest{}
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
	resp, err := h.naturalPkService.Save(r.Context(), createInput)
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
// @Tags natural-pk
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Param request body naturalPkUpdateRequest true "Update input"
// @Success 201 {object} naturalPkResponse
// @Failure 400 {object} httputils.ErrorResponse "Bad Request"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error"
// @Router /api/v1/natural-pk/{code} [put]
func (h *NaturalPkHTTPHandler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	pkCode, err := httputils.PathValueString(r, "code")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	req := &naturalPkUpdateRequest{}
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
	resp, err := h.naturalPkService.UpdateByID(r.Context(), updateInput, pkCode)
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
// @Tags natural-pk
// @Accept json
// @Produce json
// @Param id path int true "NaturalPk ID"
// @Success 204 "No Content (Successfully deleted)"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid ID format)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Deletion failed)"
// @Router /api/v1/natural-pk/{code} [delete]
func (h *NaturalPkHTTPHandler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	pkCode, err := httputils.PathValueString(r, "code")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	err = h.naturalPkService.DeleteByID(r.Context(), pkCode)
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
// @Tags natural-pk
// @Produce json
// @Param id path int true "Item ID"
// @Success 200 {object} naturalPkResponse "Single item"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid ID format)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Deletion failed)"
// @Router /api/v1/natural-pk/{code} [get]
func (h *NaturalPkHTTPHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	pkCode, err := httputils.PathValueString(r, "code")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	resp, err := h.naturalPkService.FindByID(r.Context(), pkCode)
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
// @Tags natural-pk
// @Accept json
// @Produce  json
// @Success 200 {object} naturalPkResponseList "List of all items"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Service failure)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Data processing failure)"
// @Router /api/v1/natural-pk [get]
func (h *NaturalPkHTTPHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	// call service
	resp, err := h.naturalPkService.FindAll(r.Context())
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
	httputils.WriteJSON(w, http.StatusOK, naturalPkResponseList{
		Data: dtosToPayloads,
	})
}

// FindAllPageable
//
// @Summary Get paginated list
// @Description Retrieves a paginated list using pagination parameters.
// @Tags natural-pk
// @Accept json
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param size query int false "Number of items per page (default: 10)"
// @Param sort query string false "Sort order, e.g., 'name,asc'"
// @Success 200 {object} naturalPkResponseList "Paginated list of NaturalPk"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid pagination parameters or service failure)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Data processing failure)"
// @Router /api/v1/natural-pk/pageable [get]
func (h *NaturalPkHTTPHandler) FindAllPageable(w http.ResponseWriter, r *http.Request) {
	pq, err := pageable.GetPaginationFromCtx(r)
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// call service
	resp, page, err := h.naturalPkService.FindAllPageable(r.Context(), pq)
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
	httputils.WriteJSON(w, http.StatusOK, naturalPkResponseList{
		Data: dtosToPayloads,
		Page: &page,
	})
}

// mappers

func mapCreateRequestToCreateInputDto(inputRequest *naturalPkCreateRequest) (*dto.NaturalPkCreateDto, error) {
	if inputRequest == nil {
		return nil, fmt.Errorf("unexpected nil input for mapping between naturalPkCreateRequest->NaturalPkCreateDto")
	}
	return &dto.NaturalPkCreateDto{
		Code: inputRequest.Code,
		Name: inputRequest.Name,
	}, nil
}

func mapUpdateRequestToUpdateInputDto(inputRequest *naturalPkUpdateRequest) (*dto.NaturalPkUpdateDto, error) {
	if inputRequest == nil {
		return nil, fmt.Errorf("unexpected nil input for mapping between naturalPkUpdateRequest->NaturalPkUpdateDto")
	}
	return &dto.NaturalPkUpdateDto{
		Name: inputRequest.Name,
	}, nil
}

func mapDtosToPayloads(inputDtos []dto.NaturalPkDto) ([]naturalPkResponse, error) {
	outputResponses := make([]naturalPkResponse, 0, len(inputDtos))
	for i := range inputDtos { // Iterate using index to avoid copying (gocritic:rangeValCopy)
		toPayload, err := mapDtoToPayload(&inputDtos[i])
		if err != nil {
			return nil, err
		}
		outputResponses = append(outputResponses, toPayload)
	}
	return outputResponses, nil
}

func mapDtoToPayload(inputDto *dto.NaturalPkDto) (naturalPkResponse, error) {
	if inputDto == nil {
		return naturalPkResponse{}, fmt.Errorf("unexpected nil input for mapping between NaturalPkDto->naturalPkResponse")
	}
	return naturalPkResponse{
		Code: inputDto.Code,
		Name: inputDto.Name,
	}, nil
}
