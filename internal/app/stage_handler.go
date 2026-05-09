package app

import "strings"

type GenHandl struct {
	Payload string
	Handler string
}

func GenHandler(s TableToStructInfo) GenHandl {
	pk := NewPrimaryKeyView(s.PrimaryKeys)

	data := map[string]any{
		"StructName":                  s.StructName,
		"StructComment":               s.StructComment,
		"PackageName":                 strings.ToLower(s.DbTableName),
		"PathIDSClause":               pk.PathRead,
		"PkeysURLPath":                pk.URLPath,
		"ArgumentsByPkeys":            pk.Args,
		"StructNameLowerFirstLetter":  s.StructNameLowerFirstLetter,
		"StructNamePluralRequestPath": s.StructNamePluralRequestPath,
		"ServiceVarName":              s.StructNameLowerFirstLetter + "Service",
		"CreateRequestName":           s.StructNameLowerFirstLetter + "CreateRequest",
		"UpdateRequestName":           s.StructNameLowerFirstLetter + "UpdateRequest",
		"ResponseName":                s.StructNameLowerFirstLetter + "Response",
		"ResponseListName":            s.StructNameLowerFirstLetter + "ResponseList",
		"DtoFieldsFull":               s.FullFields(),
		"DtoFieldsCreate":             s.InsertFields(),
		"DtoFieldsUpdate":             s.UpdateFields(),
		"HasPrimaryKey":               s.HasPrimaryKey,
		"HasUpdateFields":             s.HasUpdateFields,
	}

	return GenHandl{
		Payload: PrintFormatted(ExecTemplate("payload", data, FuncMap)),
		Handler: PrintFormatted(ExecTemplate("handler", data, FuncMap)),
	}
}
