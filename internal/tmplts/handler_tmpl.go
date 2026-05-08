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
{{- range .DtoFieldsCreate}}
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
{{- range .DtoFieldsUpdate}}
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
	if err != nil {
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

{{- if .HasUpdateFields}}
// UpdateByID
//
// @Summary Update existing item
// @Description Updates an item by its ID
// @Tags {{.StructNamePluralRequestPath}}
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Param request body {{.UpdateRequestName}} true "Update input"
// @Success 201 {object} {{.ResponseName}}
// @Failure 400 {object} httputils.ErrorResponse "Bad Request"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error"
// @Router /api/v1/{{.StructNamePluralRequestPath}}/{{.PkeysURLPath}} [put]
func (h *{{.ImplName}}) UpdateByID(w http.ResponseWriter, r *http.Request) {
	{{.PathIDSClause}}

	req := &{{.UpdateRequestName}}{}
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
	resp, err := h.{{.ServiceVarName}}.UpdateByID(r.Context(), updateInput, {{.ArgumentsByPkeys}})
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

{{- end}}

{{- if .HasPrimaryKey}}
// DeleteByID
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
// @Router /api/v1/{{.StructNamePluralRequestPath}}/{{.PkeysURLPath}} [delete]
func (h *{{.ImplName}}) DeleteByID(w http.ResponseWriter, r *http.Request) {
	{{.PathIDSClause}}

	err = h.{{.ServiceVarName}}.DeleteByID(r.Context(), {{.ArgumentsByPkeys}})
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
// @Tags {{.StructNamePluralRequestPath}}
// @Produce json
// @Param id path int true "Item ID"
// @Success 200 {object} {{.ResponseName}} "Single item"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid ID format)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Deletion failed)"
// @Router /api/v1/{{.StructNamePluralRequestPath}}/{{.PkeysURLPath}} [get]
func (h *{{.ImplName}}) FindByID(w http.ResponseWriter, r *http.Request) {
	{{.PathIDSClause}}

	resp, err := h.{{.ServiceVarName}}.FindByID(r.Context(), {{.ArgumentsByPkeys}})
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

{{- end}}

// FindAll
//
// @Summary Get all
// @Description Retrieves a list without pagination.
// @Tags {{.StructNamePluralRequestPath}}
// @Accept json
// @Produce  json
// @Success 200 {object} {{.ResponseListName}} "List of all items"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Service failure)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Data processing failure)"
// @Router /api/v1/{{.StructNamePluralRequestPath}} [get]
func (h *{{.ImplName}}) FindAll(w http.ResponseWriter, r *http.Request) {
	// call service
	resp, err := h.{{.ServiceVarName}}.FindAll(r.Context())
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

// FindAllPageable
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
func (h *{{.ImplName}}) FindAllPageable(w http.ResponseWriter, r *http.Request) {
	pq, err := pageable.GetPaginationFromCtx(r)
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// call service
	resp, page, err := h.{{.ServiceVarName}}.FindAllPageable(r.Context(), pq)
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

// mappers

func mapCreateRequestToCreateInputDto(inputRequest *{{.CreateRequestName}}) (*dto.{{.DtoCreateName}}, error) {
	if inputRequest == nil {
		return nil, fmt.Errorf("unexpected nil input for mapping between {{.CreateRequestName}}->{{.DtoCreateName}}")
	}
	return &dto.{{.DtoCreateName}} {
{{- range .DtoFieldsCreate}}
	{{.FieldName}}: inputRequest.{{.FieldName}},
{{- end}}
	}, nil
}

func mapUpdateRequestToUpdateInputDto(inputRequest *{{.UpdateRequestName}}) (*dto.{{.DtoUpdateName}}, error) {
	if inputRequest == nil {
		return nil, fmt.Errorf("unexpected nil input for mapping between {{.UpdateRequestName}}->{{.DtoUpdateName}}")
	}
	return &dto.{{.DtoUpdateName}} {
{{- range .DtoFieldsUpdate}}
	{{.FieldName}}: inputRequest.{{.FieldName}},
{{- end}}
	}, nil
}

func mapDtosToPayloads(inputDtos []dto.{{.DtoName}}) ([]{{.ResponseName}}, error) {
	outputResponses := make([]{{.ResponseName}}, 0, len(inputDtos))
	for i := range inputDtos { // Iterate using index to avoid copying (gocritic:rangeValCopy)
		toPayload, err := mapDtoToPayload(&inputDtos[i])
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
