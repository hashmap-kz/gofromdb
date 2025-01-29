package tmplts

var HandlerPayloadsTmpl = `
package v1

import (
	"go-project-template-v5/pkg/pageable"
	"time"
)

{{- if .StructComment}}
// {{.StructNameLowerFirstLetter}}CreateRequest {{.StructComment | ToLower}}
{{- end}}
type {{.StructNameLowerFirstLetter}}CreateRequest struct {
{{- range .DtoFieldsCreate}}
	{{- if .FieldComment}}
	// {{.FieldComment}}
	{{- end}}
	{{.FieldName}} {{.FieldType}} ` + "`json:\"{{.DbFieldName}}\"`" + `
	{{- if .FieldComment}}
		{{print "\n"}}
	{{- end}}
{{- end}}
}

{{- if .StructComment}}
// {{.StructNameLowerFirstLetter}}UpdateRequest {{.StructComment | ToLower}}
{{- end}}
type {{.StructNameLowerFirstLetter}}UpdateRequest struct {
{{- range .DtoFieldsUpdate}}
	{{- if .FieldComment}}
	// {{.FieldComment}}
	{{- end}}
	{{.FieldName}} {{.FieldType}} ` + "`json:\"{{.DbFieldName}}\"`" + `
	{{- if .FieldComment}}
		{{print "\n"}}
	{{- end}}
{{- end}}
}

{{- if .StructComment}}
// {{.StructNameLowerFirstLetter}}Response {{.StructComment | ToLower}}
{{- end}}
type {{.StructNameLowerFirstLetter}}Response struct {
{{- range .DtoFieldsFull}}
	{{- if .FieldComment}}
	// {{.FieldComment}}
	{{- end}}
	{{.FieldName}} {{.FieldType}} ` + "`json:\"{{.DbFieldName}}\"`" + `
	{{- if .FieldComment}}
		{{print "\n"}}
	{{- end}}
{{- end}}
}

// {{.StructNameLowerFirstLetter}}ResponseList response list
type {{.StructNameLowerFirstLetter}}ResponseList struct {
	// Page information (if present)
	Page pageable.Page ` + "`json:\"page,omitempty\"`" + `
	
	// Payload
	Data []{{.StructNameLowerFirstLetter}}Response ` + "`json:\"data\"`" + `
}
`
