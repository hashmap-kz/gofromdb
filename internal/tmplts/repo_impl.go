package tmplts

var RepoSaveQueryTemplate = `
insert into {{.SchemaName}}.{{.TableName}} (
{{.FieldsNoPKeys | AddPadding}}
)
values ({{.ValuesPlaceholders}})
returning
{{.FieldsWithPKeys | AddPadding}}
`

var RepoUpdateQueryTemplate = `
update {{.SchemaName}}.{{.TableName}}
set 
{{.FieldsNoPKeysWithPlaceholders | AddPadding}}
where {{.PkeyFieldName}} = $1
returning 
{{.FieldsWithPKeys | AddPadding}}
`

var RepoDeleteQueryTemplate = `
delete from only {{.SchemaName}}.{{.TableName}}
where {{.PkeyFieldName}} = $1
`

var RepoGetByIdQueryTemplate = `
select
{{.FieldsWithPKeys | AddPadding}}
from {{.SchemaName}}.{{.TableName}}
where {{.PkeyFieldName}} = $1
order by {{.PkeyFieldName}}
`

var RepoGetAllQueryTemplate = `
select
{{.FieldsWithPKeys | AddPadding}}
from {{.SchemaName}}.{{.TableName}}
order by {{.PkeyFieldName}}
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

func (r *{{.ImplName}}) Save(ctx context.Context, inputEntity *dbModel.{{.StructName}}) (*dbModel.{{.StructName}}, error) {
	tag := "{{.ImplName}}.Save"

	query := ` + "`{{.RepoSaveQuery | AddPadding2}}`" + `

	var scannedEntity dbModel.{{.StructName}}
	err := r.db.Pool.QueryRow(ctx, query,
{{- range .StructFieldsNoPKeys}}
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

func (r *{{.ImplName}}) Update(ctx context.Context, inputEntity *dbModel.{{.StructName}}) (*dbModel.{{.StructName}}, error) {
	tag := "{{.ImplName}}.Update"

	query := ` + "`{{.RepoUpdateQuery | AddPadding2}}`" + `

	var scannedEntity dbModel.{{.StructName}}
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

func (r *{{.ImplName}}) Delete(ctx context.Context, id int) error {
	tag := "{{.ImplName}}.Delete"

	query := ` + "`{{.RepoDeleteQuery | AddPadding2}}`" + `

	cmdTag, err := r.db.Pool.Exec(ctx, query, id)
	if err != nil || cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("%s. no rows deleted for id: %v, %w", tag, id, err)
	}
	return nil
}

func (r *{{.ImplName}}) GetByID(ctx context.Context, id int) (*dbModel.{{.StructName}}, error) {
	tag := "{{.ImplName}}.GetByID"

	query := ` + "`{{.RepoGetByIdQuery | AddPadding2}}`" + `

	var scannedEntity dbModel.{{.StructName}}
	err := r.db.Pool.QueryRow(ctx, query, id).Scan(
{{- range .StructFieldsWithPKeys}}
		&scannedEntity.{{.}},
{{- end }}
	)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return &scannedEntity, nil
}

func (r *{{.ImplName}}) GetAll(ctx context.Context) ([]dbModel.{{.StructName}}, error) {
	tag := "{{.ImplName}}.GetAll"

	query := ` + "`{{.RepoGetAllQuery | AddPadding2}}`" + `

	rows, err := r.db.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	var scannedEntities []dbModel.{{.StructName}}
	for rows.Next() {
		var scannedEntity dbModel.{{.StructName}}
		err = rows.Scan(
{{- range .StructFieldsWithPKeys}}
			&scannedEntity.{{.}},
{{- end }}
		)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", tag, err)
		}
		scannedEntities = append(scannedEntities, scannedEntity)
	}
	
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return scannedEntities, nil
}
`
