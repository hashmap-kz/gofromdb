package core

import "strings"

type ServiceTemplateData struct {
	StructName        string
	PackageName       string
	ParametersByPkeys string
	ArgumentsByPkeys  string

	HasPrimaryKey   bool
	HasUpdateFields bool

	DtoFieldsCreate []TableToStructFieldInfo
	DtoFieldsUpdate []TableToStructFieldInfo
	DtoFieldsFull   []TableToStructFieldInfo
}

func NewServiceTemplateData(s TableToStructInfo) (ServiceTemplateData, error) {
	pk, err := NewPrimaryKeyView(s.PrimaryKeys)
	if err != nil {
		return ServiceTemplateData{}, err
	}

	return ServiceTemplateData{
		StructName:        s.StructName,
		PackageName:       strings.ToLower(s.DbTableName),
		ParametersByPkeys: pk.Params,
		ArgumentsByPkeys:  pk.Args,
		HasPrimaryKey:     s.HasPrimaryKey,
		HasUpdateFields:   s.HasUpdateFields,
		DtoFieldsCreate:   s.InsertFields(),
		DtoFieldsUpdate:   s.UpdateFields(),
		DtoFieldsFull:     s.FullFields(),
	}, nil
}
