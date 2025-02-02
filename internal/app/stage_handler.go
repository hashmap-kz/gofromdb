package app

import (
	"fmt"
	"strings"

	"genpg-v5/internal/tmplts"
)

type GenHandl struct {
	HandlerDtos string
	HandlerImpl string
}

func GenHandler(s TableToStructInfo) GenHandl {
	clause := genPathIDSClause(s.PrimaryKeys)
	argumentsByPkeys := genArgumentsByPkeys(s)

	data := map[string]any{
		"StructName":                  s.StructName,
		"StructComment":               s.StructComment,
		"PackageName":                 strings.ToLower(s.DbTableName),
		"ImplName":                    s.StructName + "HTTPHandler",
		"PathIDSClause":               clause,
		"PkeysURLPath":                s.PkeysURLPath,
		"ArgumentsByPkeys":            argumentsByPkeys,
		"StructNameLowerFirstLetter":  s.StructNameLowerFirstLetter,
		"StructNamePluralRequestPath": s.StructNamePluralRequestPath,
		"ServiceVarName":              s.StructNameLowerFirstLetter + "Service",
		"ServiceInterfaceName":        s.StructName + "Service",
		"CreateRequestName":           s.StructNameLowerFirstLetter + "CreateRequest",
		"UpdateRequestName":           s.StructNameLowerFirstLetter + "UpdateRequest",
		"ResponseName":                s.StructNameLowerFirstLetter + "Response",
		"ResponseListName":            s.StructNameLowerFirstLetter + "ResponseList",
		"DtoName":                     s.StructName + "Dto",
		"DtoUpdateName":               s.StructName + "UpdateDto",
		"DtoCreateName":               s.StructName + "CreateDto",
		"DtoFieldsFull": s.GetStructFields(Filters{
			WithInsertableOnly: false,
			WithInternals:      true,
		}),
		"DtoFieldsNoPkeysNoDefaults": s.GetStructFields(Filters{
			WithInsertableOnly: true,
			WithInternals:      false,
		}),
	}

	dtosResult := ExecTemplate("handler-dtos", tmplts.HandlerPayloadsTmpl, data, FuncMap)
	implResult := ExecTemplate("handler-impl", tmplts.HandlerImpl, data, FuncMap)

	return GenHandl{
		HandlerDtos: PrintFormatted(dtosResult),
		HandlerImpl: PrintFormatted(implResult),
	}
}

func genPathIDSClause(pkeys []TableToStructFieldInfo) string {
	tmpl := `
	pk%s, err := %s(r, "%s")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}`

	sb := strings.Builder{}
	for _, f := range pkeys {
		if f.FieldType == "int" {
			sb.WriteString(fmt.Sprintf(tmpl, f.FieldName, "httputils.PathValueI32", f.DbFieldName))
		} else if f.FieldType == "int64" {
			sb.WriteString(fmt.Sprintf(tmpl, f.FieldName, "httputils.PathValueI32", f.DbFieldName))
		} else {
			sb.WriteString(fmt.Sprintf(tmpl, f.FieldName, "httputils.PathValueString", f.DbFieldName))
		}
	}
	return sb.String()
}
