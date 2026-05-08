package app

import (
	"fmt"
	"log"
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
		"DtoFieldsCreate": s.GetStructFields(Filters{
			WithInsertableOnly: true,
			WithInternals:      false,
		}),
		"DtoFieldsUpdate": s.GetStructFields(Filters{
			WithInsertableOnly: true,
			WithInternals:      false,
			WithoutPrimaryKeys: true,
		}),
		"HasPrimaryKey":   s.HasPrimaryKey,
		"HasUpdateFields": s.HasUpdateFields,
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
		parser, ok := pathValueParser(f.FieldType)
		if !ok {
			log.Fatalf("cannot generate handler path parser for primary key column %q with Go type %q", f.DbFieldName, f.FieldType)
		}
		sb.WriteString(fmt.Sprintf(tmpl, f.FieldName, parser, f.DbFieldName))
	}
	return sb.String()
}

func pathValueParser(fieldType string) (string, bool) {
	switch fieldType {
	case "int", "int32":
		return "httputils.PathValueI32", true
	case "int16":
		return "httputils.PathValueI16", true
	case "int64":
		return "httputils.PathValueI64", true
	case "uint32":
		return "httputils.PathValueU32", true
	case "string":
		return "httputils.PathValueString", true
	default:
		return "", false
	}
}
