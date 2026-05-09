package app

import "strings"

type GenSvc struct {
	Dto     string
	Service string
}

func GenService(s TableToStructInfo) GenSvc {
	pk := NewPrimaryKeyView(s.PrimaryKeys)

	data := map[string]any{
		"StructName":        s.StructName,
		"PackageName":       strings.ToLower(s.DbTableName),
		"ParametersByPkeys": pk.Params,
		"ArgumentsByPkeys":  pk.Args,
		"DtoFieldsFull":     s.FullFields(),
		"DtoFieldsCreate":   s.InsertFields(),
		"DtoFieldsUpdate":   s.UpdateFields(),
		"HasPrimaryKey":     s.HasPrimaryKey,
		"HasUpdateFields":   s.HasUpdateFields,
	}

	return GenSvc{
		Dto:     PrintFormatted(ExecTemplate("dto", data, FuncMap)),
		Service: PrintFormatted(ExecTemplate("service", data, FuncMap)),
	}
}
