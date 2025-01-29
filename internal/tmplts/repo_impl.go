package tmplts

var RepoSaveQueryTemplate = `
INSERT INTO {{.SchemaName}}.{{.TableName}} (
{{.FieldsNoPKeys | AddPadding}}
)
VALUES ({{.ValuesPlaceholders}})
RETURNING
{{.FieldsWithPKeys | AddPadding}}
`

var RepoUpdateQueryTemplate = `
UPDATE {{.SchemaName}}.{{.TableName}}
SET 
{{.FieldsNoPKeysWithPlaceholders | AddPadding}}
WHERE {{.PkeyFieldName}} = $1
RETURNING 
{{.FieldsWithPKeys | AddPadding}}
`

var RepoImplTemplate = `
package impl

import (
	"context"
	"fmt"

	dbModel "go-project-template-v5/internal/api/{{.PackageName}}/entity/postgres"

	"go-project-template-v5/internal/api/{{.PackageName}}/repository"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/pkg/storage/postgres"
)

type {{.ImplName}} struct {
	db *postgres.Postgres
}

var _ repository.{{.InterfaceName}} = &{{.ImplName}}{}

func NewClientRepository(_ context.Context, db *postgres.Postgres) repository.{{.InterfaceName}} {
	return &{{.ImplName}}{
		db: db,
	}
}

func (r *{{.ImplName}}) Save(ctx context.Context, inputEntity *model.{{.StructName}}) (*model.{{.StructName}}, error) {
	tag := "{{.ImplName}}.Save"

	query := ` + "`{{.RepoSaveQuery | AddPadding2}}`" + `

	var scannedEntity model.{{.StructName}}
	err := r.db.Pool.QueryRow(ctx, query,
{{- range .StructFieldsNoPkeys}}
		inputEntity.{{.}},
{{- end }}
	).Scan(
{{- range .StructFieldsWithPKeys}}
		&scannedEntity.{{.}},
{{- end }}
	) 

	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return &scannedEntity, nil
}

func (r *clientRepository) Update(ctx context.Context, inputEntity *dbModel.Client) (*dbModel.Client, error) {
	tag := "clientRepository.Update"

	query := ` + "`{{.RepoUpdateQuery | AddPadding2}}`" + `

	var scannedEntity model.{{.StructName}}
	err := r.db.Pool.QueryRow(ctx, query,
{{- range .StructFieldsUpdate}}
		inputEntity.{{.}},
{{- end }}
	).Scan(
{{- range .StructFieldsWithPKeys}}
		&scannedEntity.{{.}},
{{- end }}
	)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return &scannedEntity, nil
}
`
