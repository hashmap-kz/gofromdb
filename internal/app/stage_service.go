package app

import (
	"strings"

	"genpg-v5/internal/tmplts"
)

type GenSvc struct {
	ServiceDtos      string
	ServiceInterface string
	ServiceImpl      string
}

func GenService(s TableToStructInfo) GenSvc {
	parametersByPkeys := genParametersByPkeys(s)
	argumentsByPkeys := genArgumentsByPkeys(s)

	data := map[string]any{
		"StructName":              s.StructName,
		"PackageName":             strings.ToLower(s.DbTableName),
		"InterfaceName":           s.StructName + "Service",
		"ImplName":                LowerFirstLetter(s.StructName) + "Service",
		"ParametersByPkeys":       parametersByPkeys,
		"ArgumentsByPkeys":        argumentsByPkeys,
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

	interfaceRes := ExecTemplate("service-interface", tmplts.ServiceInterfaceTemplate, data, FuncMap)
	modelsRes := ExecTemplate("service-dtos", tmplts.ServiceDtosTemplate, data, FuncMap)
	implRes := ExecTemplate("service-impl", tmplts.ServiceImplTemplate, data, FuncMap)

	return GenSvc{
		ServiceDtos:      PrintFormatted(modelsRes),
		ServiceInterface: PrintFormatted(interfaceRes),
		ServiceImpl:      PrintFormatted(implRes),
	}
}
