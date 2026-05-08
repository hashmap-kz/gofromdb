package v1

import (
	"fmt"
	"net/http"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/no_pk/dto"
	"go-project-template-v5/internal/api/no_pk/service"

	"go-project-template-v5/pkg/httputils"
	"go-project-template-v5/pkg/validator"
)

type NoPkHTTPHandler struct {
	noPkService service.NoPkService
}

func NewNoPkHTTPHandler(noPkService service.NoPkService) *NoPkHTTPHandler {
	return &NoPkHTTPHandler{
		noPkService: noPkService,
	}
}

// Save
//
// @Summary Create new item
// @Description Create new item handler
// @Tags no-pk
// @Accept json
// @Produce json
// @Param request body noPkCreateRequest true "Create input"
// @Success 201 {object} noPkResponse
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid request payload)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error"
// @Router /api/v1/no-pk [post]
func (h *NoPkHTTPHandler) Save(w http.ResponseWriter, r *http.Request) {
	// read RequestBody
	req := &noPkCreateRequest{}
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
	resp, err := h.noPkService.Save(r.Context(), createInput)
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

// FindAll
//
// @Summary Get all
// @Description Retrieves a list without pagination.
// @Tags no-pk
// @Accept json
// @Produce  json
// @Success 200 {object} noPkResponseList "List of all items"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Service failure)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Data processing failure)"
// @Router /api/v1/no-pk [get]
func (h *NoPkHTTPHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	// call service
	resp, err := h.noPkService.FindAll(r.Context())
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
	httputils.WriteJSON(w, http.StatusOK, noPkResponseList{
		Data: dtosToPayloads,
	})
}

// FindAllPageable
//
// @Summary Get paginated list
// @Description Retrieves a paginated list using pagination parameters.
// @Tags no-pk
// @Accept json
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param size query int false "Number of items per page (default: 10)"
// @Param sort query string false "Sort order, e.g., 'name,asc'"
// @Success 200 {object} noPkResponseList "Paginated list of NoPk"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid pagination parameters or service failure)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Data processing failure)"
// @Router /api/v1/no-pk/pageable [get]
func (h *NoPkHTTPHandler) FindAllPageable(w http.ResponseWriter, r *http.Request) {
	pq, err := pageable.GetPaginationFromCtx(r)
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// call service
	resp, page, err := h.noPkService.FindAllPageable(r.Context(), pq)
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
	httputils.WriteJSON(w, http.StatusOK, noPkResponseList{
		Data: dtosToPayloads,
		Page: &page,
	})
}

// mappers

func mapCreateRequestToCreateInputDto(inputRequest *noPkCreateRequest) (*dto.NoPkCreateDto, error) {
	if inputRequest == nil {
		return nil, fmt.Errorf("unexpected nil input for mapping between noPkCreateRequest->NoPkCreateDto")
	}
	return &dto.NoPkCreateDto{
		EventTime: inputRequest.EventTime,
		Message:   inputRequest.Message,
	}, nil
}

func mapUpdateRequestToUpdateInputDto(inputRequest *noPkUpdateRequest) (*dto.NoPkUpdateDto, error) {
	if inputRequest == nil {
		return nil, fmt.Errorf("unexpected nil input for mapping between noPkUpdateRequest->NoPkUpdateDto")
	}
	return &dto.NoPkUpdateDto{
		EventTime: inputRequest.EventTime,
		Message:   inputRequest.Message,
	}, nil
}

func mapDtosToPayloads(inputDtos []dto.NoPkDto) ([]noPkResponse, error) {
	outputResponses := make([]noPkResponse, 0, len(inputDtos))
	for i := range inputDtos { // Iterate using index to avoid copying (gocritic:rangeValCopy)
		toPayload, err := mapDtoToPayload(&inputDtos[i])
		if err != nil {
			return nil, err
		}
		outputResponses = append(outputResponses, toPayload)
	}
	return outputResponses, nil
}

func mapDtoToPayload(inputDto *dto.NoPkDto) (noPkResponse, error) {
	if inputDto == nil {
		return noPkResponse{}, fmt.Errorf("unexpected nil input for mapping between NoPkDto->noPkResponse")
	}
	return noPkResponse{
		EventTime: inputDto.EventTime,
		Message:   inputDto.Message,
	}, nil
}
