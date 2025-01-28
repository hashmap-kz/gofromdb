package tmplts

var RepoSaveQueryTemplate = `
INSERT INTO {{.SchemaName}}.{{.TableName}} (
{{.FieldsNoPKeys | AddPadding}}
)
VALUES ({{.ValuesPlaceholders}})
RETURNING
{{.FieldsWithPKeys | AddPadding}}
`

var RepoSaveFuncTemplate = `
func (r *{{.StructName | ToCamel}}Repository) Save(ctx context.Context, entity *model.{{.StructName}}) (*model.{{.StructName}}, error) {
	query := ` + "`{{.Query}}`" + `

	var i model.{{.StructName}}
	if err := r.db.Pool.QueryRow(ctx, query,
{{.StructFieldsNoPKeys | AddPadding}}
	).Scan(
{{.StructFieldsWithPKeys | AddPadding}}
	); err != nil {
		return nil, err
	}

	return &i, nil
}
`
