package app

import (
	"genpg-v5/internal/tmplts"
)

type GenHandl struct {
	HandlerDtos string
	HandlerImpl string
}

func GenHandler(s TableToStructInfo) GenHandl {
	// Dto template

	structFieldsWithoutPkeysAndDefaults := s.GetStructFields(false, true)
	structFieldsWithPkeysAndWithoutDefaults := s.GetStructFields(true, true)

	dtosResult := ExecTemplate("handler-dtos", tmplts.HandlerPayloadsTmpl,
		map[string]any{
			"StructNameLowerFirstLetter": s.StructNameLowerFirstLetter,
			"DtoFieldsFull":              s.Fields,
			"DtoFieldsCreate":            structFieldsWithoutPkeysAndDefaults,
			"DtoFieldsUpdate":            structFieldsWithPkeysAndWithoutDefaults,
		}, FuncMap)

	return GenHandl{
		HandlerDtos: PrintFormatted(dtosResult),
		HandlerImpl: "",
	}
}
