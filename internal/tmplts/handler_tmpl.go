package tmplts

var HandlerPayloadsTmpl = `
package v1

import (
	"go-project-template-v5/pkg/pageable"
	"time"
)

type {{.StructNameLowerFirstLetter}}CreateRequest struct {
{{- range .DtoFieldsCreate}}
	{{.FieldName}} {{.FieldType}} ` + "`json:\"{{.DbFieldName}}\"`" + `
{{- end}}
}

type {{.StructNameLowerFirstLetter}}Response struct {
{{- range .DtoFieldsFull}}
	{{.FieldName}} {{.FieldType}} ` + "`json:\"{{.DbFieldName}}\"`" + `
{{- end}}
}

type {{.StructNameLowerFirstLetter}}ResponseList struct {
	Page pageable.Page 
	Data []{{.StructNameLowerFirstLetter}}Response ` + "`json:\"data\"`" + `
}

type {{.StructNameLowerFirstLetter}}UpdateRequest struct {
{{- range .DtoFieldsUpdate}}
	{{.FieldName}} {{.FieldType}} ` + "`json:\"{{.DbFieldName}}\"`" + `
{{- end}}
}
`
