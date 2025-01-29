package tmplts

var EntityTemplate = `
{{- if .StructComment}}
// {{.StructName}}. {{.StructComment}}
{{- end}}
type {{.StructName}} struct {
{{- range .Columns}}
	{{- if .FieldComment}}
	// {{.FieldComment}}
	{{- end}}
	{{.FieldName}} {{.FieldType}} ` + "`json:\"{{.DbFieldName}}\" db:\"{{.DbFieldName}}\"`" + `
	{{- if .FieldComment}}
		{{print "\n"}}
	{{- end}}
{{- end}}
}
`
