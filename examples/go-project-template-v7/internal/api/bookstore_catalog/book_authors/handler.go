package book_authors

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
// @Tags book-authors
// @Accept json
// @Produce json
// @Param request body bookAuthorsCreateRequest true "Create input"
// @Success 201 {object} bookAuthorsResponse
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid request payload)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error"
// @Router /api/v1/book-authors [post]
func (h *Handler) Save(w http.ResponseWriter, r *http.Request) {
	req := &bookAuthorsCreateRequest{}
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
// @Tags book-authors
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Param request body bookAuthorsUpdateRequest true "Update input"
// @Success 200 {object} bookAuthorsResponse
// @Failure 400 {object} httputils.ErrorResponse "Bad Request"
// @Failure 404 {object} httputils.ErrorResponse "Not Found"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error"
// @Router /api/v1/book-authors/{book_id}/{author_id} [put]
func (h *Handler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	pkBookID, err := httputils.PathValueI64(r, "book_id")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}
	pkAuthorID, err := httputils.PathValueString(r, "author_id")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	req := &bookAuthorsUpdateRequest{}
	if err := httputils.ReadJSON(r, &req); err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}
	if err := validator.ValidateStruct(req); err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}
	updateInput := mapUpdateRequestToUpdateInputDto(req)
	resp, err := h.svc.UpdateByID(r.Context(), updateInput, pkBookID, pkAuthorID)
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
// @Tags book-authors
// @Accept json
// @Produce json
// @Param id path int true "BookAuthors ID"
// @Success 204 "No Content (Successfully deleted)"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid ID format)"
// @Failure 404 {object} httputils.ErrorResponse "Not Found"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error"
// @Router /api/v1/book-authors/{book_id}/{author_id} [delete]
func (h *Handler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	pkBookID, err := httputils.PathValueI64(r, "book_id")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}
	pkAuthorID, err := httputils.PathValueString(r, "author_id")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	if err := h.svc.DeleteByID(r.Context(), pkBookID, pkAuthorID); err != nil {
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
// @Tags book-authors
// @Produce json
// @Param id path int true "Item ID"
// @Success 200 {object} bookAuthorsResponse "Single item"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid ID format)"
// @Failure 404 {object} httputils.ErrorResponse "Not Found"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error"
// @Router /api/v1/book-authors/{book_id}/{author_id} [get]
func (h *Handler) FindByID(w http.ResponseWriter, r *http.Request) {
	pkBookID, err := httputils.PathValueI64(r, "book_id")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}
	pkAuthorID, err := httputils.PathValueString(r, "author_id")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	resp, err := h.svc.FindByID(r.Context(), pkBookID, pkAuthorID)
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
// @Tags book-authors
// @Accept json
// @Produce  json
// @Success 200 {object} bookAuthorsResponseList "List of all items"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error"
// @Router /api/v1/book-authors [get]
func (h *Handler) FindAll(w http.ResponseWriter, r *http.Request) {
	resp, err := h.svc.FindAll(r.Context())
	if err != nil {
		httputils.WriteJSON(w, http.StatusInternalServerError, httputils.ErrorResponse{Message: err.Error()})
		return
	}
	httputils.WriteJSON(w, http.StatusOK, bookAuthorsResponseList{
		Data: mapDtosToPayloads(resp),
	})
}

// FindAllPageable
//
// @Summary Get paginated list
// @Description Retrieves a paginated list using pagination parameters.
// @Tags book-authors
// @Accept json
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param size query int false "Number of items per page (default: 10)"
// @Param sort query string false "Sort order, e.g., 'name,asc'"
// @Success 200 {object} bookAuthorsResponseList "Paginated list of BookAuthors"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid pagination parameters)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error"
// @Router /api/v1/book-authors/pageable [get]
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
	httputils.WriteJSON(w, http.StatusOK, bookAuthorsResponseList{
		Data: mapDtosToPayloads(resp),
		Page: &page,
	})
}

// mappers

func mapCreateRequestToCreateInputDto(inputRequest *bookAuthorsCreateRequest) *CreateDto {
	return &CreateDto{
		BookID:            inputRequest.BookID,
		AuthorID:          inputRequest.AuthorID,
		ContributionOrder: inputRequest.ContributionOrder,
		Role:              inputRequest.Role,
		Notes:             inputRequest.Notes,
	}
}

func mapUpdateRequestToUpdateInputDto(inputRequest *bookAuthorsUpdateRequest) *UpdateDto {
	return &UpdateDto{
		ContributionOrder: inputRequest.ContributionOrder,
		Role:              inputRequest.Role,
		Notes:             inputRequest.Notes,
	}
}

func mapDtosToPayloads(inputDtos []Dto) []bookAuthorsResponse {
	outputResponses := make([]bookAuthorsResponse, 0, len(inputDtos))
	for i := range inputDtos {
		outputResponses = append(outputResponses, mapDtoToPayload(&inputDtos[i]))
	}
	return outputResponses
}

func mapDtoToPayload(inputDto *Dto) bookAuthorsResponse {
	return bookAuthorsResponse{
		BookID:            inputDto.BookID,
		AuthorID:          inputDto.AuthorID,
		ContributionOrder: inputDto.ContributionOrder,
		Role:              inputDto.Role,
		Notes:             inputDto.Notes,
	}
}
