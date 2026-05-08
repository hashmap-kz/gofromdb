package app

import "strings"

type GenHandl struct {
	HandlerDtos string
	HandlerImpl string
}

func GenHandler(s TableToStructInfo) GenHandl {
	pk := NewPrimaryKeyView(s.PrimaryKeys)

	data := map[string]any{
		"StructName":                  s.StructName,
		"StructComment":               s.StructComment,
		"PackageName":                 strings.ToLower(s.DbTableName),
		"ImplName":                    s.StructName + "HTTPHandler",
		"PathIDSClause":               pk.PathRead,
		"PkeysURLPath":                pk.URLPath,
		"ArgumentsByPkeys":            pk.Args,
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
		"DtoFieldsFull":               s.FullFields(),
		"DtoFieldsCreate":             s.InsertFields(),
		"DtoFieldsUpdate":             s.UpdateFields(),
		"HasPrimaryKey":               s.HasPrimaryKey,
		"HasUpdateFields":             s.HasUpdateFields,
	}

	return GenHandl{
		HandlerDtos: PrintFormatted(ExecTemplate("handler_payloads", data, FuncMap)),
		HandlerImpl: PrintFormatted(ExecTemplate("handler_impl", data, FuncMap)),
	}
}
