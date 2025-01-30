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
	structFieldsWithoutPkeysAndDefaults := s.GetStructFields(false, true)

	data := map[string]any{
		"PackageName":                strings.ToLower(s.DbTableName),
		"StructNameLowerFirstLetter": s.StructNameLowerFirstLetter,
		"StructComment":              s.StructComment,
		"ImplName":                   s.StructName + "HTTPHandler",
		"ServiceVarName":             s.StructNameLowerFirstLetter + "Service",
		"ServiceInterfaceName":       s.StructName + "Service",
		"CreateRequestName":          s.StructNameLowerFirstLetter + "CreateRequest",
		"UpdateRequestName":          s.StructNameLowerFirstLetter + "UpdateRequest",
		"ResponseName":               s.StructNameLowerFirstLetter + "Response",
		"ResponseListName":           s.StructNameLowerFirstLetter + "ResponseList",
		"DtoName":                    s.StructName + "Dto",
		"DtoUpdateName":              s.StructName + "UpdateDto",
		"DtoCreateName":              s.StructName + "CreateDto",
		"DtoFieldsFull":              s.Fields,
		"DtoFieldsNoPkeysNoDefaults": structFieldsWithoutPkeysAndDefaults,
	}

	dtosResult := ExecTemplate("handler-dtos", tmplts.HandlerPayloadsTmpl, data, FuncMap)
	implResult := ExecTemplate("handler-impl", tmplts.HandlerImpl, data, FuncMap)

	return GenHandl{
		HandlerDtos: PrintFormatted(dtosResult),
		HandlerImpl: PrintFormatted(implResult),
	}
}
