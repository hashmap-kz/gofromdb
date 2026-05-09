package core

import (
	"fmt"
	"strings"
)

type GenSvc struct {
	Dto     string
	Service string
}

func GenService(s TableToStructInfo) (GenSvc, error) {
	pk, err := NewPrimaryKeyView(s.PrimaryKeys)
	if err != nil {
		return GenSvc{}, err
	}

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

	exec := func(name string) (string, error) {
		out, err := ExecTemplate(name, data, FuncMap)
		if err != nil {
			return "", fmt.Errorf("service %s: %w", name, err)
		}
		return PrintFormatted(out), nil
	}

	dto, err := exec("dto")
	if err != nil {
		return GenSvc{}, err
	}
	svc, err := exec("service")
	if err != nil {
		return GenSvc{}, err
	}

	return GenSvc{
		Dto:     dto,
		Service: svc,
	}, nil
}
