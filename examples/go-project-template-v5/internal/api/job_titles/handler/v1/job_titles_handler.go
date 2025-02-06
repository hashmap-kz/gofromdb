package v1

import (
	"fmt"
	"net/http"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/job_titles/dto"
	"go-project-template-v5/internal/api/job_titles/service"

	"go-project-template-v5/pkg/httputils"
	"go-project-template-v5/pkg/validator"
)

type JobTitlesHTTPHandler struct {
	jobTitlesService service.JobTitlesService
}

func NewJobTitlesHTTPHandler(jobTitlesService service.JobTitlesService) *JobTitlesHTTPHandler {
	return &JobTitlesHTTPHandler{
		jobTitlesService: jobTitlesService,
	}
}

// Save
//
// @Summary Create new item
// @Description Create new item handler
// @Tags job-titles
// @Accept json
// @Produce json
// @Param request body jobTitlesCreateRequest true "Create input"
// @Success 201 {object} jobTitlesResponse
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid request payload)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error"
// @Router /api/v1/job-titles [post]
func (h *JobTitlesHTTPHandler) Save(w http.ResponseWriter, r *http.Request) {
	// read RequestBody
	req := &jobTitlesCreateRequest{}
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
	resp, err := h.jobTitlesService.Save(r.Context(), createInput)
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
// @Tags job-titles
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Param request body jobTitlesUpdateRequest true "Update input"
// @Success 201 {object} jobTitlesResponse
// @Failure 400 {object} httputils.ErrorResponse "Bad Request"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error"
// @Router /api/v1/job-titles/{record_id} [put]
func (h *JobTitlesHTTPHandler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	pkRecordID, err := httputils.PathValueI32(r, "record_id")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	req := &jobTitlesUpdateRequest{}
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
	resp, err := h.jobTitlesService.UpdateByID(r.Context(), updateInput, pkRecordID)
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
// @Tags job-titles
// @Accept json
// @Produce json
// @Param id path int true "JobTitles ID"
// @Success 204 "No Content (Successfully deleted)"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid ID format)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Deletion failed)"
// @Router /api/v1/job-titles/{record_id} [delete]
func (h *JobTitlesHTTPHandler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	pkRecordID, err := httputils.PathValueI32(r, "record_id")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	err = h.jobTitlesService.DeleteByID(r.Context(), pkRecordID)
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
// @Tags job-titles
// @Produce json
// @Param id path int true "Item ID"
// @Success 200 {object} jobTitlesResponse "Single item"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid ID format)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Deletion failed)"
// @Router /api/v1/job-titles/{record_id} [get]
func (h *JobTitlesHTTPHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	pkRecordID, err := httputils.PathValueI32(r, "record_id")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	resp, err := h.jobTitlesService.FindByID(r.Context(), pkRecordID)
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
// @Tags job-titles
// @Accept json
// @Produce  json
// @Success 200 {object} jobTitlesResponseList "List of all items"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Service failure)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Data processing failure)"
// @Router /api/v1/job-titles [get]
func (h *JobTitlesHTTPHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	// call service
	resp, err := h.jobTitlesService.FindAll(r.Context())
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
	httputils.WriteJSON(w, http.StatusOK, jobTitlesResponseList{
		Data: dtosToPayloads,
	})
}

// FindAllPageable
//
// @Summary Get paginated list
// @Description Retrieves a paginated list using pagination parameters.
// @Tags job-titles
// @Accept json
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param size query int false "Number of items per page (default: 10)"
// @Param sort query string false "Sort order, e.g., 'name,asc'"
// @Success 200 {object} jobTitlesResponseList "Paginated list of JobTitles"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid pagination parameters or service failure)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Data processing failure)"
// @Router /api/v1/job-titles/pageable [get]
func (h *JobTitlesHTTPHandler) FindAllPageable(w http.ResponseWriter, r *http.Request) {
	pq, err := pageable.GetPaginationFromCtx(r)
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// call service
	resp, page, err := h.jobTitlesService.FindAllPageable(r.Context(), pq)
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
	httputils.WriteJSON(w, http.StatusOK, jobTitlesResponseList{
		Data: dtosToPayloads,
		Page: &page,
	})
}

// mappers

func mapCreateRequestToCreateInputDto(inputRequest *jobTitlesCreateRequest) (*dto.JobTitlesCreateDto, error) {
	if inputRequest == nil {
		return nil, fmt.Errorf("unexpected nil input for mapping between jobTitlesCreateRequest->JobTitlesCreateDto")
	}
	return &dto.JobTitlesCreateDto{
		TitleName: inputRequest.TitleName,
	}, nil
}

func mapUpdateRequestToUpdateInputDto(inputRequest *jobTitlesUpdateRequest) (*dto.JobTitlesUpdateDto, error) {
	if inputRequest == nil {
		return nil, fmt.Errorf("unexpected nil input for mapping between jobTitlesUpdateRequest->JobTitlesUpdateDto")
	}
	return &dto.JobTitlesUpdateDto{
		TitleName: inputRequest.TitleName,
	}, nil
}

func mapDtosToPayloads(inputDtos []dto.JobTitlesDto) ([]jobTitlesResponse, error) {
	var outputResponses []jobTitlesResponse
	for _, inputDto := range inputDtos {
		toPayload, err := mapDtoToPayload(&inputDto)
		if err != nil {
			return nil, err
		}
		outputResponses = append(outputResponses, toPayload)
	}
	return outputResponses, nil
}

func mapDtoToPayload(inputDto *dto.JobTitlesDto) (jobTitlesResponse, error) {
	if inputDto == nil {
		return jobTitlesResponse{}, fmt.Errorf("unexpected nil input for mapping between JobTitlesDto->jobTitlesResponse")
	}
	return jobTitlesResponse{
		RecordID:  inputDto.RecordID,
		TitleName: inputDto.TitleName,
		CreatedAt: inputDto.CreatedAt,
		UpdatedAt: inputDto.UpdatedAt,
		Guid:      inputDto.Guid,
	}, nil
}
