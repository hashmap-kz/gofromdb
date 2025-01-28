package tmplts

var EntityTemplate = `
type {{.StructName}} struct {
{{- range .Columns}}
	{{.FieldName}} {{.FieldType}} ` + "`json:\"{{.DbFieldName}}\" db:\"{{.DbFieldName}}\"`" + `
{{- end}}
}
`
