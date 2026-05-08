package app

import "strings"

type GenSvc struct {
	ServiceDtos      string
	ServiceInterface string
	ServiceImpl      string
}

func GenService(s TableToStructInfo) GenSvc {
	pk := NewPrimaryKeyView(s.PrimaryKeys)

	data := map[string]any{
		"StructName":              s.StructName,
		"PackageName":             strings.ToLower(s.DbTableName),
		"InterfaceName":           s.StructName + "Service",
		"ImplName":                LowerFirstLetter(s.StructName) + "Service",
		"ParametersByPkeys":       pk.Params,
		"ArgumentsByPkeys":        pk.Args,
		"RepositoryInterfaceName": s.StructName + "Repository",
		"RepositoryVarName":       s.StructNameLowerFirstLetter + "Repository",
		"DtoName":                 s.StructName + "Dto",
		"DtoUpdateName":           s.StructName + "UpdateDto",
		"DtoCreateName":           s.StructName + "CreateDto",
		"DtoFieldsFull":           s.FullFields(),
		"DtoFieldsCreate":         s.InsertFields(),
		"DtoFieldsUpdate":         s.UpdateFields(),
		"HasPrimaryKey":           s.HasPrimaryKey,
		"HasUpdateFields":         s.HasUpdateFields,
	}

	return GenSvc{
		ServiceDtos:      PrintFormatted(ExecTemplate("service_dtos", data, FuncMap)),
		ServiceInterface: PrintFormatted(ExecTemplate("service_interface", data, FuncMap)),
		ServiceImpl:      PrintFormatted(ExecTemplate("service_impl", data, FuncMap)),
	}
}
