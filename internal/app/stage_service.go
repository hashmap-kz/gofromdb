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
	// Interface template

	interfaceResult := ExecTemplate("service-interface", tmplts.ServiceInterfaceTemplate,
		map[string]any{
			"StructName":    s.StructName,
			"PackageName":   strings.ToLower(s.DbTableName),
			"InterfaceName": s.StructName + "Service",
			"ImplName":      LowerFirstLetter(s.StructName) + "Service",
			"DtoName":       s.StructName + "Dto",
			"DtoUpdateName": s.StructName + "UpdateDto",
			"DtoCreateName": s.StructName + "CreateDto",
		}, FuncMap)

	// Dto template

	structFieldsWithoutPkeysAndDefaults := s.GetStructFields(Filters{
		WithPkeys:    false,
		SkipInternal: true,
	})

	dtosResult := ExecTemplate("service-dtos", tmplts.ServiceDtosTemplate,
		map[string]any{
			"StructName":                 s.StructName,
			"PackageName":                strings.ToLower(s.DbTableName),
			"InterfaceName":              s.StructName + "Service",
			"ImplName":                   LowerFirstLetter(s.StructName) + "Service",
			"DtoName":                    s.StructName + "Dto",
			"DtoUpdateName":              s.StructName + "UpdateDto",
			"DtoCreateName":              s.StructName + "CreateDto",
			"DtoFieldsFull":              s.Fields,
			"DtoFieldsNoPkeysNoDefaults": structFieldsWithoutPkeysAndDefaults,
		}, FuncMap)

	// Functions template

	structFieldsWithoutPkeysAndDefaultsStr := s.GetStructFieldsAsString(Filters{
		WithPkeys:    false,
		SkipInternal: true,
	})
	structFieldsWithPkeysAndDefaultsStr := s.GetStructFieldsAsString(Filters{
		WithPkeys:    true,
		SkipInternal: false,
	})

	implResult := ExecTemplate("service-impl", tmplts.ServiceImplTemplate,
		map[string]any{
			"StructName":            s.StructName,
			"PackageName":           strings.ToLower(s.DbTableName),
			"InterfaceName":         s.StructName + "Service",
			"RepositoryName":        s.StructName + "Repository",
			"ImplName":              LowerFirstLetter(s.StructName) + "Service",
			"DtoName":               s.StructName + "Dto",
			"DtoUpdateName":         s.StructName + "UpdateDto",
			"DtoCreateName":         s.StructName + "CreateDto",
			"StructFieldsNoPKeys":   structFieldsWithoutPkeysAndDefaultsStr,
			"StructFieldsWithPKeys": structFieldsWithPkeysAndDefaultsStr,
		}, FuncMap)

	return GenSvc{
		ServiceDtos:      PrintFormatted(dtosResult),
		ServiceInterface: PrintFormatted(interfaceResult),
		ServiceImpl:      PrintFormatted(implResult),
	}
}
