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
	data := map[string]any{
		"StructName":     s.StructName,
		"PackageName":    strings.ToLower(s.DbTableName),
		"InterfaceName":  s.StructName + "Service",
		"RepositoryName": s.StructName + "Repository",
		"ImplName":       LowerFirstLetter(s.StructName) + "Service",
		"DtoName":        s.StructName + "Dto",
		"DtoUpdateName":  s.StructName + "UpdateDto",
		"DtoCreateName":  s.StructName + "CreateDto",
		"DtoFieldsFull": s.GetStructFields(Filters{
			WithInsertableOnly: false,
			WithInternals:      true,
		}),
		"DtoFieldsNoPkeysNoDefaults": s.GetStructFields(Filters{
			WithInsertableOnly: true,
			WithInternals:      false,
		}),
	}

	interfaceResult := ExecTemplate("service-interface", tmplts.ServiceInterfaceTemplate, data, FuncMap)
	dtosResult := ExecTemplate("service-dtos", tmplts.ServiceDtosTemplate, data, FuncMap)
	implResult := ExecTemplate("service-impl", tmplts.ServiceImplTemplate, data, FuncMap)

	return GenSvc{
		ServiceDtos:      PrintFormatted(dtosResult),
		ServiceInterface: PrintFormatted(interfaceResult),
		ServiceImpl:      PrintFormatted(implResult),
	}
}
