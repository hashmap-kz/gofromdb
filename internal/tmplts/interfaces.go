package tmplts

var RepoInterfaceGeneral = `
package api

import (
	"context"

{{- range .Structs}}
	{{.StructNameLowerFirstLetter}}Repo "go-project-template-v5/internal/api/{{.DbTableName}}/repository"
	{{.StructNameLowerFirstLetter}}Impl "go-project-template-v5/internal/api/{{.DbTableName}}/repository/impl"
{{- end }}

	"go-project-template-v5/pkg/storage/postgres"
)

// Init all repos

type Repositories struct {
{{- range .Structs}}
	{{.StructName}}Repository {{.StructNameLowerFirstLetter}}Repo.{{.StructName}}Repository
{{- end }}
}

func NewRepositories(ctx context.Context, db *postgres.Postgres) *Repositories {
	return &Repositories{
{{- range .Structs}}
		{{.StructName}}Repository: {{.StructNameLowerFirstLetter}}Impl.New{{.StructName}}Repository(ctx, db),
{{- end }}
	}
}
`

var ServiceInterfaceGeneral = `
package api

import (
	"context"
{{- range .Structs}}
	{{.StructNameLowerFirstLetter}}Serv "go-project-template-v5/internal/api/{{.DbTableName}}/service"
	{{.StructNameLowerFirstLetter}}Impl "go-project-template-v5/internal/api/{{.DbTableName}}/service/impl"
{{- end }}
)

// Init all services

type Services struct {
{{- range .Structs}}
	{{.StructName}}Service {{.StructNameLowerFirstLetter}}Serv.{{.StructName}}Service
{{- end }}
}

type Deps struct {
	// TODO: other deps here
	Repos *Repositories
}

func NewServices(ctx context.Context, deps Deps) *Services {
	return &Services{
{{- range .Structs}}
		{{.StructName}}Service: {{.StructNameLowerFirstLetter}}Impl.New{{.StructName}}Service(ctx, deps.Repos.{{.StructName}}Repository),
{{- end }}
	}
}
`

var HandlerInterfaceGeneral = `
package api

import (
	"net/http"
{{- range .Structs}}
	{{.StructNameLowerFirstLetter}}v1 "go-project-template-v5/internal/api/{{.DbTableName}}/handler/v1"
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
	{{.StructNameLowerFirstLetter}}Handler := {{.StructNameLowerFirstLetter}}v1.New{{.StructName}}HTTPHandler(h.Services.{{.StructName}}Service)
	router.HandleFunc("POST /api/v1/{{.StructNamePluralRequestPath}}"						, {{.StructNameLowerFirstLetter}}Handler.Save)
{{- if .HasUpdateFields}}
	router.HandleFunc("PUT /api/v1/{{.StructNamePluralRequestPath}}/{{.PkeysURLPath}}"		, {{.StructNameLowerFirstLetter}}Handler.UpdateByID)
{{- end}}
{{- if .HasPrimaryKey}}
	router.HandleFunc("DELETE /api/v1/{{.StructNamePluralRequestPath}}/{{.PkeysURLPath}}"	, {{.StructNameLowerFirstLetter}}Handler.DeleteByID)
	router.HandleFunc("GET /api/v1/{{.StructNamePluralRequestPath}}/{{.PkeysURLPath}}"		, {{.StructNameLowerFirstLetter}}Handler.FindByID)
{{- end}}
	router.HandleFunc("GET /api/v1/{{.StructNamePluralRequestPath}}"						, {{.StructNameLowerFirstLetter}}Handler.FindAll)
	router.HandleFunc("GET /api/v1/{{.StructNamePluralRequestPath}}/pageable"				, {{.StructNameLowerFirstLetter}}Handler.FindAllPageable)
{{print "\n"}}
{{- end }}
}
`
