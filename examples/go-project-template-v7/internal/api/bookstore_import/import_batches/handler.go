package import_batches

import (
	"errors"
	"go-project-template-v7/pkg/apperrors"
	"go-project-template-v7/pkg/httputils"
	"go-project-template-v7/pkg/pageable"
	"go-project-template-v7/pkg/validator"
	"net/http"
)

type Handler struct {
	svc Service
}

func NewHandler(svc Service) *Handler {
	return &Handler{
		svc: svc,
	}
}

// Save
//
// @Summary Create new item
// @Description Create new item handler
// @Tags import-batches
// @Accept json
// @Produce json
// @Param request body importBatchesCreateRequest true "Create input"
// @Success 201 {object} importBatchesResponse
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid request payload)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error"
// @Router /api/v1/import-batches [post]
func (h *Handler) Save(w http.ResponseWriter, r *http.Request) {
	req := &importBatchesCreateRequest{}
	if err := httputils.ReadJSON(r, &req); err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}
	if err := validator.ValidateStruct(req); err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}
	createInput := mapCreateRequestToCreateInputDto(req)
	resp, err := h.svc.Save(r.Context(), createInput)
	if err != nil {
		httputils.WriteJSON(w, http.StatusInternalServerError, httputils.ErrorResponse{Message: err.Error()})
		return
	}
	httputils.WriteJSON(w, http.StatusCreated, mapDtoToPayload(resp))
}

// UpdateByID
//
// @Summary Update existing item
// @Description Updates an item by its ID
// @Tags import-batches
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Param request body importBatchesUpdateRequest true "Update input"
// @Success 200 {object} importBatchesResponse
// @Failure 400 {object} httputils.ErrorResponse "Bad Request"
// @Failure 404 {object} httputils.ErrorResponse "Not Found"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error"
// @Router /api/v1/import-batches/{source_name}/{batch_no} [put]
func (h *Handler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	pkSourceName, err := httputils.PathValueString(r, "source_name")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}
	pkBatchNo, err := httputils.PathValueI32(r, "batch_no")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	req := &importBatchesUpdateRequest{}
	if err := httputils.ReadJSON(r, &req); err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}
	if err := validator.ValidateStruct(req); err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}
	updateInput := mapUpdateRequestToUpdateInputDto(req)
	resp, err := h.svc.UpdateByID(r.Context(), updateInput, pkSourceName, pkBatchNo)
	if err != nil {
		if errors.Is(err, apperrors.ErrNotFound) {
			httputils.WriteJSON(w, http.StatusNotFound, httputils.ErrorResponse{Message: "record not found"})
			return
		}
		httputils.WriteJSON(w, http.StatusInternalServerError, httputils.ErrorResponse{Message: err.Error()})
		return
	}
	httputils.WriteJSON(w, http.StatusOK, mapDtoToPayload(resp))
}

// DeleteByID
//
// @Summary Delete existing item
// @Description Deletes an item by its ID
// @Tags import-batches
// @Accept json
// @Produce json
// @Param id path int true "ImportBatches ID"
// @Success 204 "No Content (Successfully deleted)"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid ID format)"
// @Failure 404 {object} httputils.ErrorResponse "Not Found"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error"
// @Router /api/v1/import-batches/{source_name}/{batch_no} [delete]
func (h *Handler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	pkSourceName, err := httputils.PathValueString(r, "source_name")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}
	pkBatchNo, err := httputils.PathValueI32(r, "batch_no")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	if err := h.svc.DeleteByID(r.Context(), pkSourceName, pkBatchNo); err != nil {
		if errors.Is(err, apperrors.ErrNotFound) {
			httputils.WriteJSON(w, http.StatusNotFound, httputils.ErrorResponse{Message: "record not found"})
			return
		}
		httputils.WriteJSON(w, http.StatusInternalServerError, httputils.ErrorResponse{Message: err.Error()})
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// FindByID retrieves an item by its ID.
//
// @Summary Get item by ID
// @Description Retrieves the details based on the provided ID in the request path.
// @Tags import-batches
// @Produce json
// @Param id path int true "Item ID"
// @Success 200 {object} importBatchesResponse "Single item"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid ID format)"
// @Failure 404 {object} httputils.ErrorResponse "Not Found"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error"
// @Router /api/v1/import-batches/{source_name}/{batch_no} [get]
func (h *Handler) FindByID(w http.ResponseWriter, r *http.Request) {
	pkSourceName, err := httputils.PathValueString(r, "source_name")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}
	pkBatchNo, err := httputils.PathValueI32(r, "batch_no")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	resp, err := h.svc.FindByID(r.Context(), pkSourceName, pkBatchNo)
	if err != nil {
		if errors.Is(err, apperrors.ErrNotFound) {
			httputils.WriteJSON(w, http.StatusNotFound, httputils.ErrorResponse{Message: "record not found"})
			return
		}
		httputils.WriteJSON(w, http.StatusInternalServerError, httputils.ErrorResponse{Message: err.Error()})
		return
	}
	httputils.WriteJSON(w, http.StatusOK, mapDtoToPayload(resp))
}

// FindAll
//
// @Summary Get all
// @Description Retrieves a list without pagination.
// @Tags import-batches
// @Accept json
// @Produce  json
// @Success 200 {object} importBatchesResponseList "List of all items"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error"
// @Router /api/v1/import-batches [get]
func (h *Handler) FindAll(w http.ResponseWriter, r *http.Request) {
	resp, err := h.svc.FindAll(r.Context())
	if err != nil {
		httputils.WriteJSON(w, http.StatusInternalServerError, httputils.ErrorResponse{Message: err.Error()})
		return
	}
	httputils.WriteJSON(w, http.StatusOK, importBatchesResponseList{
		Data: mapDtosToPayloads(resp),
	})
}

// FindAllPageable
//
// @Summary Get paginated list
// @Description Retrieves a paginated list using pagination parameters.
// @Tags import-batches
// @Accept json
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param size query int false "Number of items per page (default: 10)"
// @Param sort query string false "Sort order, e.g., 'name,asc'"
// @Success 200 {object} importBatchesResponseList "Paginated list of ImportBatches"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid pagination parameters)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error"
// @Router /api/v1/import-batches/pageable [get]
func (h *Handler) FindAllPageable(w http.ResponseWriter, r *http.Request) {
	pq, err := pageable.GetPaginationFromCtx(r)
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}
	resp, page, err := h.svc.FindAllPageable(r.Context(), pq)
	if err != nil {
		httputils.WriteJSON(w, http.StatusInternalServerError, httputils.ErrorResponse{Message: err.Error()})
		return
	}
	httputils.WriteJSON(w, http.StatusOK, importBatchesResponseList{
		Data: mapDtosToPayloads(resp),
		Page: &page,
	})
}

// mappers

func mapCreateRequestToCreateInputDto(inputRequest *importBatchesCreateRequest) *CreateDto {
	return &CreateDto{
		SourceName: inputRequest.SourceName,
		BatchNo:    inputRequest.BatchNo,
		StartedAt:  inputRequest.StartedAt,
		FinishedAt: inputRequest.FinishedAt,
		FileName:   inputRequest.FileName,
		RowCount:   inputRequest.RowCount,
		Metadata:   inputRequest.Metadata,
	}
}

func mapUpdateRequestToUpdateInputDto(inputRequest *importBatchesUpdateRequest) *UpdateDto {
	return &UpdateDto{
		StartedAt:  inputRequest.StartedAt,
		FinishedAt: inputRequest.FinishedAt,
		FileName:   inputRequest.FileName,
		RowCount:   inputRequest.RowCount,
		Metadata:   inputRequest.Metadata,
	}
}

func mapDtosToPayloads(inputDtos []Dto) []importBatchesResponse {
	outputResponses := make([]importBatchesResponse, 0, len(inputDtos))
	for i := range inputDtos {
		outputResponses = append(outputResponses, mapDtoToPayload(&inputDtos[i]))
	}
	return outputResponses
}

func mapDtoToPayload(inputDto *Dto) importBatchesResponse {
	return importBatchesResponse{
		SourceName: inputDto.SourceName,
		BatchNo:    inputDto.BatchNo,
		StartedAt:  inputDto.StartedAt,
		FinishedAt: inputDto.FinishedAt,
		FileName:   inputDto.FileName,
		RowCount:   inputDto.RowCount,
		Metadata:   inputDto.Metadata,
	}
}
