package app

import (
	"strings"

	"genpg-v5/internal/tmplts"
)

type GenHandl struct {
	HandlerDtos string
	HandlerImpl string
}

func GenHandler(s TableToStructInfo) GenHandl {
	data := map[string]any{
		"StructName":                  s.StructName,
		"StructComment":               s.StructComment,
		"PackageName":                 strings.ToLower(s.DbTableName),
		"ImplName":                    s.StructName + "HTTPHandler",
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
