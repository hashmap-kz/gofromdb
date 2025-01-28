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
func (r *clientRepository) Save(ctx context.Context, entity *model.Client) (*model.Client, error) {
	query := ` + "`%s`" + `

	var i model.Client
	if err := r.db.Pool.QueryRow(ctx, query,
		entity.Email,
	).Scan(
		&i.RecordID,
		&i.Email,
	); err != nil {
		return nil, err
	}

	return &i, nil
}
`
