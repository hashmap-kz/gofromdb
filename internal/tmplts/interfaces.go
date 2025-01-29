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

// Init

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

	clientServiceInterface "go-project-template-v5/internal/api/{{.DbTableName}}/service"
	clientServiceImpl "go-project-template-v5/internal/api/{{.DbTableName}}/service/impl"
)

// Init

type Services struct {
	// TODO: other service interfaces here
	ClientService clientServiceInterface.ClientService
}

type Deps struct {
	// TODO: other deps here
	Repos *Repositories
}

func NewServices(ctx context.Context, deps Deps) *Services {
	return &Services{
		// TODO: other service impls here
		ClientService: clientServiceImpl.NewClientService(ctx, deps.Repos.ClientRepository),
	}
}
`
