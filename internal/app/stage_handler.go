package app

import (
	"fmt"
)

type GenHandl struct {
	Payload string
	Handler string
}

func GenHandler(s TableToStructInfo) (GenHandl, error) {
	pk, err := NewPrimaryKeyView(s.PrimaryKeys)
	if err != nil {
		return GenHandl{}, err
	}

	data := map[string]any{
		"StructName":                  s.StructName,
		"StructComment":               s.StructComment,
		"PackageName":                 s.OutDirName,
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

	exec := func(name string) (string, error) {
		out, err := ExecTemplate(name, data, FuncMap)
		if err != nil {
			return "", fmt.Errorf("handler %s: %w", name, err)
		}
		return PrintFormatted(out), nil
	}

	payload, err := exec("payload")
	if err != nil {
		return GenHandl{}, err
	}
	handler, err := exec("handler")
	if err != nil {
		return GenHandl{}, err
	}

	return GenHandl{
		Payload: payload,
		Handler: handler,
	}, nil
}
