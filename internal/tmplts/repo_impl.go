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
	UpdateByID(ctx context.Context, entityId int, inputEntity *dbModel.{{.StructName}}) (*dbModel.{{.StructName}}, error)
	DeleteByID(ctx context.Context, entityId int) error
	FindByID(ctx context.Context, entityId int) (*dbModel.{{.StructName}}, error)
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

func (r *{{.ImplName}}) UpdateByID(ctx context.Context, entityId int, inputEntity *dbModel.{{.StructName}}) (*dbModel.{{.StructName}}, error) {
	tag := "{{.ImplName}}.UpdateByID"

	query := ` + "`{{.RepoUpdateQuery | AddPadding2}}`" + `

	var scannedEntity dbModel.{{.StructName}}
	err := r.db.Pool.QueryRow(ctx, query,
		entityId,
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

func (r *{{.ImplName}}) DeleteByID(ctx context.Context, entityId int) error {
	tag := "{{.ImplName}}.DeleteByID"

	query := ` + "`{{.RepoDeleteQuery | AddPadding2}}`" + `

	cmdTag, err := r.db.Pool.Exec(ctx, query, entityId)
	if err != nil || cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("%s. no rows deleted for id: %v, %w", tag, entityId, err)
	}
	return nil
}

func (r *{{.ImplName}}) FindByID(ctx context.Context, entityId int) (*dbModel.{{.StructName}}, error) {
	tag := "{{.ImplName}}.FindByID"

	query := ` + "`{{.RepoGetByIdQuery | AddPadding2}}`" + `

	var scannedEntity dbModel.{{.StructName}}
	err := r.db.Pool.QueryRow(ctx, query, entityId).Scan(
{{- range .StructFieldsWithPKeys}}
		&scannedEntity.{{.}},
{{- end }}
	)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return &scannedEntity, nil
}

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
		var scannedEntity dbModel.{{.StructName}}
		err = rows.Scan(
{{- range .StructFieldsWithPKeys}}
			&scannedEntity.{{.}},
{{- end }}
		)
		if err != nil {
			return nil, pageable.Page{}, fmt.Errorf("%s: %w", tag, err)
		}
		scannedEntities = append(scannedEntities, scannedEntity)
	}

	if rows.Err() != nil {
		return nil, page, rows.Err()
	}
	return scannedEntities, page, nil
}
`
