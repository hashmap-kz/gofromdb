package core

import "strings"

type HandlerTemplateData struct {
	StructName                  string
	StructComment               string
	PackageName                 string
	PathIDSClause               string
	PkeysURLPath                string
	ArgumentsByPkeys            string
	StructNameLowerFirstLetter  string
	StructNamePluralRequestPath string
	ServiceVarName              string
	CreateRequestName           string
	UpdateRequestName           string
	ResponseName                string
	ResponseListName            string

	HasPrimaryKey   bool
	HasUpdateFields bool

	DtoFieldsCreate []TableToStructFieldInfo
	DtoFieldsUpdate []TableToStructFieldInfo
	DtoFieldsFull   []TableToStructFieldInfo
}

func NewHandlerTemplateData(s TableToStructInfo) (HandlerTemplateData, error) {
	pk, err := NewPrimaryKeyView(s.PrimaryKeys)
	if err != nil {
		return HandlerTemplateData{}, err
	}

	return HandlerTemplateData{
		StructName:                  s.StructName,
		StructComment:               s.StructComment,
		PackageName:                 strings.ToLower(s.DbTableName),
		PathIDSClause:               pk.PathRead,
		PkeysURLPath:                pk.URLPath,
		ArgumentsByPkeys:            pk.Args,
		StructNameLowerFirstLetter:  s.StructNameLowerFirstLetter,
		StructNamePluralRequestPath: s.StructNamePluralRequestPath,
		ServiceVarName:              s.StructNameLowerFirstLetter + "Service",
		CreateRequestName:           s.StructNameLowerFirstLetter + "CreateRequest",
		UpdateRequestName:           s.StructNameLowerFirstLetter + "UpdateRequest",
		ResponseName:                s.StructNameLowerFirstLetter + "Response",
		ResponseListName:            s.StructNameLowerFirstLetter + "ResponseList",
		HasPrimaryKey:               s.HasPrimaryKey,
		HasUpdateFields:             s.HasUpdateFields,
		DtoFieldsCreate:             s.InsertFields(),
		DtoFieldsUpdate:             s.UpdateFields(),
		DtoFieldsFull:               s.FullFields(),
	}, nil
}
