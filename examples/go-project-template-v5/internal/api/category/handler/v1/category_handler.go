package v1

import (
	"fmt"
	"net/http"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/category/dto"
	"go-project-template-v5/internal/api/category/service"

	"go-project-template-v5/pkg/httputils"
	"go-project-template-v5/pkg/validator"
)

type CategoryHTTPHandler struct {
	categoryService service.CategoryService
}

func NewCategoryHTTPHandler(categoryService service.CategoryService) *CategoryHTTPHandler {
	return &CategoryHTTPHandler{
		categoryService: categoryService,
	}
}

// Save
//
// @Summary Create new Category
// @Description Create Category handler
// @Tags Category
// @Accept json
// @Produce json
// @Param request body categoryCreateRequest true "Create input"
// @Success 201 {object} categoryResponse
// @Router /api/v1/categories [post]
func (h *CategoryHTTPHandler) Save(w http.ResponseWriter, r *http.Request) {
	// read RequestBody
	req := &categoryCreateRequest{}
	if err := httputils.ReadJSON(r, &req); err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// check RequestBody
	if err := validator.ValidateStruct(req); err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// call service
	resp, err := h.categoryService.Save(r.Context(), &dto.CategoryCreateDto{
		Name:     req.Name,
		ParentID: req.ParentID,
	})
	if err != nil {
		httputils.WriteJSON(w, http.StatusInternalServerError, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// convert service-model to handler-payload
	dtoToPayload, err := mapDtoToPayload(resp)
	if err != nil {
		httputils.WriteJSON(w, http.StatusInternalServerError, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// 200 OK
	httputils.WriteJSON(w, http.StatusOK, dtoToPayload)
}

func (h *CategoryHTTPHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	// call service
	resp, err := h.categoryService.GetAll(r.Context())
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
	httputils.WriteJSON(w, http.StatusOK, categoryResponseList{
		Data: dtosToPayloads,
	})
}

func (h *CategoryHTTPHandler) GetAllPaginated(w http.ResponseWriter, r *http.Request) {
	pq, err := pageable.GetPaginationFromCtx(r)
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// call service
	resp, page, err := h.categoryService.GetAllPaginated(r.Context(), pq)
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
	httputils.WriteJSON(w, http.StatusOK, categoryResponseList{
		Data: dtosToPayloads,
		Page: &page,
	})
}

// Update
//
// @Summary Update existing Category
// @Description Update Category handler
// @Tags Category
// @Accept json
// @Produce json
// @Param request body categoryUpdateRequest true "Update input"
// @Success 201 {object} categoryResponse
// @Router /api/v1/categories [put]
func (h *CategoryHTTPHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := httputils.PathValueI64(r, "id")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	req := &categoryUpdateRequest{}
	if err := httputils.ReadJSON(r, &req); err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// TODO: types - int(id)
	resp, err := h.categoryService.Update(r.Context(), int(id), &dto.CategoryUpdateDto{
		Name:     req.Name,
		ParentID: req.ParentID,
	})
	if err != nil {
		httputils.WriteJSON(w, http.StatusInternalServerError, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// convert service-model to handler-payload
	dtoToPayload, err := mapDtoToPayload(resp)
	if err != nil {
		httputils.WriteJSON(w, http.StatusInternalServerError, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// 201 OK
	httputils.WriteJSON(w, http.StatusCreated, dtoToPayload)
}

func (h *CategoryHTTPHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := httputils.PathValueI64(r, "id")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	err = h.categoryService.Delete(r.Context(), int(id))
	if err != nil {
		httputils.WriteJSON(w, http.StatusInternalServerError, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// 204 OK
	w.WriteHeader(http.StatusNoContent)
}

func (h *CategoryHTTPHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := httputils.PathValueI64(r, "id")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	resp, err := h.categoryService.GetByID(r.Context(), int(id))
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

// mappers

func mapDtosToPayloads(inputDtos []dto.CategoryDto) ([]categoryResponse, error) {
	var outputResponses []categoryResponse
	for _, inputDto := range inputDtos {
		toPayload, err := mapDtoToPayload(&inputDto)
		if err != nil {
			return nil, err
		}
		outputResponses = append(outputResponses, toPayload)
	}
	return outputResponses, nil
}

func mapDtoToPayload(inputDto *dto.CategoryDto) (categoryResponse, error) {
	if inputDto == nil {
		return categoryResponse{}, fmt.Errorf("unexpected nil input for mapping between CategoryDto->categoryResponse")
	}
	return categoryResponse{
		RecordID:  inputDto.RecordID,
		Name:      inputDto.Name,
		ParentID:  inputDto.ParentID,
		CreatedAt: inputDto.CreatedAt,
		UpdatedAt: inputDto.UpdatedAt,
		Guid:      inputDto.Guid,
	}, nil
}
