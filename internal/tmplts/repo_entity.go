package tmplts

var EntityTemplate = `
package postgres
import "time"

{{- if .StructComment}}
// {{.StructName}} {{.StructComment | ToLower}}
{{- end}}
type {{.StructName}} struct {
{{- range .DtoFieldsFull}}
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
