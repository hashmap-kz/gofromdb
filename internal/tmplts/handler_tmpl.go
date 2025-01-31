package tmplts

var HandlerPayloadsTmpl = `
package v1

import (
	"go-project-template-v5/pkg/pageable"
	"time"
)

{{- if .StructComment}}
// {{.CreateRequestName}} {{.StructComment | ToLower}}
{{- end}}
type {{.CreateRequestName}} struct {
{{- range .DtoFieldsNoPkeysNoDefaults}}
	{{- if .FieldComment}}
	// {{.FieldComment}}
	{{- end}}
	{{.FieldName}} {{.FieldType}} ` + "`json:\"{{.DbFieldName}}\"`" + `
	{{- if .FieldComment}}
		{{print "\n"}}
	{{- end}}
{{- end}}
}

{{- if .StructComment}}
// {{.UpdateRequestName}} {{.StructComment | ToLower}}
{{- end}}
type {{.UpdateRequestName}} struct {
{{- range .DtoFieldsNoPkeysNoDefaults}}
	{{- if .FieldComment}}
	// {{.FieldComment}}
	{{- end}}
	{{.FieldName}} {{.FieldType}} ` + "`json:\"{{.DbFieldName}}\"`" + `
	{{- if .FieldComment}}
		{{print "\n"}}
	{{- end}}
{{- end}}
}

{{- if .StructComment}}
// {{.ResponseName}} {{.StructComment | ToLower}}
{{- end}}
type {{.ResponseName}} struct {
{{- range .DtoFieldsFull}}
	{{- if .FieldComment}}
	// {{.FieldComment}}
	{{- end}}
	{{.FieldName}} {{.FieldType}} ` + "`json:\"{{.DbFieldName}}\"`" + `
	{{- if .FieldComment}}
		{{print "\n"}}
	{{- end}}
{{- end}}
}

// {{.ResponseListName}} response list
type {{.ResponseListName}} struct {
	// Page information (if present)
	Page *pageable.Page ` + "`json:\"page,omitempty\"`" + `
	
	// Payload
	Data []{{.ResponseName}} ` + "`json:\"data\"`" + `
}
`

var HandlerImpl = `
package v1

import (
	"fmt"
	"net/http"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/{{.PackageName}}/dto"
	"go-project-template-v5/internal/api/{{.PackageName}}/service"

	"go-project-template-v5/pkg/httputils"
	"go-project-template-v5/pkg/validator"
)

type {{.ImplName}} struct {
	{{.ServiceVarName}} service.{{.ServiceInterfaceName}}
}

func New{{.ImplName}}({{.ServiceVarName}} service.{{.ServiceInterfaceName}}) *{{.ImplName}} {
	return &{{.ImplName}}{
		{{.ServiceVarName}}: {{.ServiceVarName}},
	}
}

// Save
//
// @Summary Create new item
// @Description Create new item handler
// @Tags {{.StructNamePluralRequestPath}}
// @Accept json
// @Produce json
// @Param request body {{.CreateRequestName}} true "Create input"
// @Success 201 {object} {{.ResponseName}}
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid request payload)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error"
// @Router /api/v1/{{.StructNamePluralRequestPath}} [post]
func (h *{{.ImplName}}) Save(w http.ResponseWriter, r *http.Request) {
	// read RequestBody
	req := &{{.CreateRequestName}}{}
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
	resp, err := h.{{.ServiceVarName}}.Save(r.Context(), createInput)
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

// GetAll
//
// @Summary Get all
// @Description Retrieves a list without pagination.
// @Tags {{.StructNamePluralRequestPath}}
// @Accept json
// @Produce  json
// @Success 200 {object} {{.ResponseListName}} "List of all BuyItems"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Service failure)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Data processing failure)"
// @Router /api/v1/{{.StructNamePluralRequestPath}} [get]
func (h *{{.ImplName}}) GetAll(w http.ResponseWriter, r *http.Request) {
	// call service
	resp, err := h.{{.ServiceVarName}}.GetAll(r.Context())
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
	httputils.WriteJSON(w, http.StatusOK, {{.ResponseListName}}{
		Data: dtosToPayloads,
	})
}

// GetAllPaginated
//
// @Summary Get paginated list
// @Description Retrieves a paginated list using pagination parameters.
// @Tags {{.StructNamePluralRequestPath}}
// @Accept json
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param size query int false "Number of items per page (default: 10)"
// @Param sort query string false "Sort order, e.g., 'name,asc'"
// @Success 200 {object} {{.ResponseListName}} "Paginated list of {{.StructName}}"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid pagination parameters or service failure)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Data processing failure)"
// @Router /api/v1/{{.StructNamePluralRequestPath}}/pageable [get]
func (h *{{.ImplName}}) GetAllPaginated(w http.ResponseWriter, r *http.Request) {
	pq, err := pageable.GetPaginationFromCtx(r)
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// call service
	resp, page, err := h.{{.ServiceVarName}}.GetAllPaginated(r.Context(), pq)
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
	httputils.WriteJSON(w, http.StatusOK, {{.ResponseListName}}{
		Data: dtosToPayloads,
		Page: &page,
	})
}

// Update
//
// @Summary Update existing item
// @Description Updates an item by its ID
// @Tags {{.StructNamePluralRequestPath}}
// @Accept json
// @Produce json
// @Param request body {{.UpdateRequestName}} true "Update input"
// @Success 201 {object} {{.ResponseName}}
// @Failure 400 {object} httputils.ErrorResponse "Bad Request"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error"
// @Router /api/v1/{{.StructNamePluralRequestPath}} [put]
func (h *{{.ImplName}}) Update(w http.ResponseWriter, r *http.Request) {
	id, err := httputils.PathValueI64(r, "id")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	req := &{{.UpdateRequestName}}{}
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
	// TODO: types - int(id)
	resp, err := h.{{.ServiceVarName}}.Update(r.Context(), int(id), updateInput)
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

// Delete
//
// @Summary Delete existing item
// @Description Deletes an item by its ID
// @Tags {{.StructNamePluralRequestPath}}
// @Accept json
// @Produce json
// @Param id path int true "{{.StructName}} ID"
// @Success 204 "No Content (Successfully deleted)"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid ID format)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Deletion failed)"
// @Router /api/v1/{{.StructNamePluralRequestPath}} [delete]
func (h *{{.ImplName}}) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := httputils.PathValueI64(r, "id")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	err = h.{{.ServiceVarName}}.Delete(r.Context(), int(id))
	if err != nil {
		httputils.WriteJSON(w, http.StatusInternalServerError, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// 204 OK
	w.WriteHeader(http.StatusNoContent)
}

func (h *{{.ImplName}}) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := httputils.PathValueI64(r, "id")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	resp, err := h.{{.ServiceVarName}}.GetByID(r.Context(), int(id))
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

func mapCreateRequestToCreateInputDto(inputRequest *{{.CreateRequestName}}) (*dto.{{.DtoCreateName}}, error) {
	if inputRequest == nil {
		return nil, fmt.Errorf("unexpected nil input for mapping between {{.CreateRequestName}}->{{.DtoCreateName}}")
	}
	return &dto.{{.DtoCreateName}} {
{{- range .DtoFieldsNoPkeysNoDefaults}}
	{{.FieldName}}: inputRequest.{{.FieldName}},
{{- end}}
	}, nil
}

func mapUpdateRequestToUpdateInputDto(inputRequest *{{.UpdateRequestName}}) (*dto.{{.DtoUpdateName}}, error) {
	if inputRequest == nil {
		return nil, fmt.Errorf("unexpected nil input for mapping between {{.UpdateRequestName}}->{{.DtoUpdateName}}")
	}
	return &dto.{{.DtoUpdateName}} {
{{- range .DtoFieldsNoPkeysNoDefaults}}
	{{.FieldName}}: inputRequest.{{.FieldName}},
{{- end}}
	}, nil
}

func mapDtosToPayloads(inputDtos []dto.{{.DtoName}}) ([]{{.ResponseName}}, error) {
	var outputResponses []{{.ResponseName}}
	for _, inputDto := range inputDtos {
		toPayload, err := mapDtoToPayload(&inputDto)
		if err != nil {
			return nil, err
		}
		outputResponses = append(outputResponses, toPayload)
	}
	return outputResponses, nil
}

func mapDtoToPayload(inputDto *dto.{{.DtoName}}) ({{.ResponseName}}, error) {
	if inputDto == nil {
		return {{.ResponseName}}{}, fmt.Errorf("unexpected nil input for mapping between {{.DtoName}}->{{.ResponseName}}")
	}
	return {{.ResponseName}}{
{{- range .DtoFieldsFull}}
		{{.FieldName}}: inputDto.{{.FieldName}},
{{- end }}
	}, nil
}
`
