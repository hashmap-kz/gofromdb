package tmplts

var RepoInterfaceTemplate = `
package repository

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	dbModel "go-project-template-v5/internal/api/{{.PackageName}}/entity/postgres"
)

type {{.InterfaceName}} interface {
	Save(ctx context.Context, inputEntity *dbModel.{{.StructName}}) (*dbModel.{{.StructName}}, error)
{{- if .HasUpdateFields}}
	UpdateByID(ctx context.Context, inputEntity *dbModel.{{.StructName}}, {{.ParametersByPkeys}}) (*dbModel.{{.StructName}}, error)
{{- end}}
{{- if .HasPrimaryKey}}
	DeleteByID(ctx context.Context, {{.ParametersByPkeys}}) error
	FindByID(ctx context.Context, {{.ParametersByPkeys}}) (*dbModel.{{.StructName}}, error)
{{- end}}
	FindAll(ctx context.Context) ([]dbModel.{{.StructName}}, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.{{.StructName}}, pageable.Page, error)
}
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
	
	"github.com/jackc/pgx/v5"
)

type {{.ImplName}} struct {
	db *postgres.Postgres
}

var _ repository.{{.InterfaceName}} = &{{.ImplName}}{}

func New{{.InterfaceName}}(_ context.Context, db *postgres.Postgres) repository.{{.InterfaceName}} {
	return &{{.ImplName}}{
		db: db,
	}
}

func (r *{{.ImplName}}) Save(ctx context.Context, inputEntity *dbModel.{{.StructName}}) (*dbModel.{{.StructName}}, error) {
	tag := "{{.ImplName}}.Save"

	query := ` + "`{{.RepoSaveQuery | AddPadding2}}`" + `

	row := r.db.Pool.QueryRow(ctx, query,
{{- range .DtoFieldsCreate}}
		inputEntity.{{.FieldName}},
{{- end}}
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

{{- if .HasUpdateFields}}
func (r *{{.ImplName}}) UpdateByID(ctx context.Context, inputEntity *dbModel.{{.StructName}}, {{.ParametersByPkeys}}) (*dbModel.{{.StructName}}, error) {
	tag := "{{.ImplName}}.UpdateByID"

	query := ` + "`{{.RepoUpdateQuery | AddPadding2}}`" + `

	row := r.db.Pool.QueryRow(ctx, query,
		{{.ArgumentsByPkeys}},
{{- range .DtoFieldsUpdate}}
		inputEntity.{{.FieldName}},
{{- end}}
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

{{- end}}

{{- if .HasPrimaryKey}}
func (r *{{.ImplName}}) DeleteByID(ctx context.Context, {{.ParametersByPkeys}}) error {
	tag := "{{.ImplName}}.DeleteByID"

	query := ` + "`{{.RepoDeleteQuery | AddPadding2}}`" + `

	cmdTag, err := r.db.Pool.Exec(ctx, query, {{.ArgumentsByPkeys}})
	if err != nil || cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("%s. no rows were deleted: %w", tag, err)
	}
	return nil
}

func (r *{{.ImplName}}) FindByID(ctx context.Context, {{.ParametersByPkeys}}) (*dbModel.{{.StructName}}, error) {
	tag := "{{.ImplName}}.FindByID"

	query := ` + "`{{.RepoGetByIdQuery | AddPadding2}}`" + `
	
	row := r.db.Pool.QueryRow(ctx, query, {{.ArgumentsByPkeys}})

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

{{- end}}

func (r *{{.ImplName}}) FindAll(ctx context.Context) ([]dbModel.{{.StructName}}, error) {
	tag := "{{.ImplName}}.FindAll"

	query := ` + "`{{.RepoGetAllQuery | AddPadding2}}`" + `

	rows, err := r.db.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	var scannedEntities []dbModel.{{.StructName}}
	for rows.Next() {
		scannedEntity, err := scanFullRow(rows)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", tag, err)
		}
		scannedEntities = append(scannedEntities, *scannedEntity)
	}
	
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return scannedEntities, nil
}

func (r *{{.ImplName}}) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.{{.StructName}}, pageable.Page, error) {
	tag := "{{.ImplName}}.FindAllPageable"

	// retrieve total count
	queryCnt := ` + "`{{.RepoCountQuery}}`" + `
	var totalCount int
	if err := r.db.Pool.QueryRow(ctx, queryCnt).Scan(&totalCount); err != nil {
		return nil, pageable.Page{}, err
	}

	// init page
	page := pageable.CreatePage(pq, totalCount)

	// handle empty result
	if totalCount == 0 {
		return nil, page, nil
	}

	// select entities
	query := ` + "`{{.RepoGetAllPaginatedQuery | AddPadding2}}`" + `

	rows, err := r.db.Pool.Query(ctx, query, pq.GetOffset(), pq.GetLimit())
	if err != nil {
		return nil, pageable.Page{}, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	var scannedEntities []dbModel.{{.StructName}}
	for rows.Next() {
		scannedEntity, err := scanFullRow(rows)
		if err != nil {
			return nil, pageable.Page{}, fmt.Errorf("%s: %w", tag, err)
		}
		scannedEntities = append(scannedEntities, *scannedEntity)
	}

	if rows.Err() != nil {
		return nil, page, rows.Err()
	}
	return scannedEntities, page, nil
}

// scan utils

// scanFullRow is expected to scan all columns from a table.
// For simplicity, most methods scan the entire row of the table into the result entity.
// You should adapt methods as needed (e.g., if business logic requires returning only an ID after an UPDATE).
func scanFullRow(row pgx.Row) (*dbModel.{{.StructName}}, error) {
	var scannedEntity dbModel.{{.StructName}}
	err := row.Scan(
{{- range .DtoFieldsFull}}
		&scannedEntity.{{.FieldName}},
{{- end}}
	)
	if err != nil {
		return nil, err
	}
	return &scannedEntity, nil
}
`
