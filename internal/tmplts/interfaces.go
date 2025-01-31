package tmplts

var RepoInterfaceGeneral = `
package api

import (
	"context"

{{- range .Structs}}
	{{.StructName | ToLower}}Repo "go-project-template-v5/internal/api/{{.DbTableName}}/repository"
	{{.StructName | ToLower}}Impl "go-project-template-v5/internal/api/{{.DbTableName}}/repository/impl"
{{- end }}

	"go-project-template-v5/pkg/storage/postgres"
)

// Init all repos

type Repositories struct {
{{- range .Structs}}
	{{.StructName}}Repository {{.StructName | ToLower}}Repo.{{.StructName}}Repository
{{- end }}
}

func NewRepositories(ctx context.Context, db *postgres.Postgres) *Repositories {
	return &Repositories{
{{- range .Structs}}
		{{.StructName}}Repository: {{.StructName | ToLower}}Impl.New{{.StructName}}Repository(ctx, db),
{{- end }}
	}
}
`

var ServiceInterfaceGeneral = `
package api

import (
	"context"
{{- range .Structs}}
	{{.StructName | ToLower}}Serv "go-project-template-v5/internal/api/{{.DbTableName}}/service"
	{{.StructName | ToLower}}Impl "go-project-template-v5/internal/api/{{.DbTableName}}/service/impl"
{{- end }}
)

// Init all services

type Services struct {
{{- range .Structs}}
	{{.StructName}}Service {{.StructName | ToLower}}Serv.{{.StructName}}Service
{{- end }}
}

type Deps struct {
	// TODO: other deps here
	Repos *Repositories
}

func NewServices(ctx context.Context, deps Deps) *Services {
	return &Services{
{{- range .Structs}}
		{{.StructName}}Service: {{.StructName | ToLower}}Impl.New{{.StructName}}Service(ctx, deps.Repos.{{.StructName}}Repository),
{{- end }}
	}
}
`

var HandlerInterfaceGeneral = `
package api

import (
	"net/http"
{{- range .Structs}}
	{{.StructName | ToLower}}v1 "go-project-template-v5/internal/api/{{.DbTableName}}/handler/v1"
{{- end }}
)

type Handler struct {
	Services *Services
}

func NewHandler(services *Services) *Handler {
	return &Handler{
		Services: services,
	}
}

func (h *Handler) Init(router *http.ServeMux) {
{{- range .Structs}}
	// {{.StructName}} routes
	{{.StructNameLowerFirstLetter}}Handler := {{.StructName | ToLower}}v1.New{{.StructName}}HTTPHandler(h.Services.{{.StructName}}Service)
	router.HandleFunc("POST /api/v1/{{.StructNamePluralRequestPath}}"			, {{.StructNameLowerFirstLetter}}Handler.Save)
	router.HandleFunc("PUT /api/v1/{{.StructNamePluralRequestPath}}/{id}"		, {{.StructNameLowerFirstLetter}}Handler.UpdateByID)
	router.HandleFunc("DELETE /api/v1/{{.StructNamePluralRequestPath}}/{id}"	, {{.StructNameLowerFirstLetter}}Handler.DeleteByID)
	router.HandleFunc("GET /api/v1/{{.StructNamePluralRequestPath}}/{id}"		, {{.StructNameLowerFirstLetter}}Handler.FindByID)
	router.HandleFunc("GET /api/v1/{{.StructNamePluralRequestPath}}"			, {{.StructNameLowerFirstLetter}}Handler.FindAll)
	router.HandleFunc("GET /api/v1/{{.StructNamePluralRequestPath}}/pageable"	, {{.StructNameLowerFirstLetter}}Handler.FindAllPageable)
{{print "\n"}}
{{- end }}
}
`
