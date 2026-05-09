package app

import (
	"fmt"
	"strings"
)

type RepoTemplateData struct {
	StructName           string
	StructComment        string
	PackageName          string
	StructNameLowerFirst string

	HasPrimaryKey   bool
	HasUpdateFields bool

	PK PrimaryKeyView

	ParametersByPkeys string
	ArgumentsByPkeys  string

	RepoSaveQuery            string
	RepoUpdateQuery          string
	RepoDeleteQuery          string
	RepoGetByIdQuery         string
	RepoGetAllQuery          string
	RepoCountQuery           string
	RepoGetAllPaginatedQuery string

	DtoFieldsCreate []TableToStructFieldInfo
	DtoFieldsUpdate []TableToStructFieldInfo
	DtoFieldsFull   []TableToStructFieldInfo
}

type RepoQueries struct {
	Save            string
	UpdateByID      string
	DeleteByID      string
	GetByID         string
	GetAll          string
	Count           string
	GetAllPaginated string
}

func NewRepoTemplateData(s TableToStructInfo) (RepoTemplateData, error) {
	pk, err := NewPrimaryKeyView(s.PrimaryKeys)
	if err != nil {
		return RepoTemplateData{}, err
	}
	queries, err := NewRepoQueries(s, pk)
	if err != nil {
		return RepoTemplateData{}, err
	}

	return RepoTemplateData{
		StructName:               s.StructName,
		StructComment:            s.StructComment,
		PackageName:              strings.ToLower(s.DbTableName),
		StructNameLowerFirst:     s.StructNameLowerFirstLetter,
		HasPrimaryKey:            s.HasPrimaryKey,
		HasUpdateFields:          s.HasUpdateFields,
		PK:                       pk,
		ParametersByPkeys:        pk.Params,
		ArgumentsByPkeys:         pk.Args,
		RepoSaveQuery:            queries.Save,
		RepoUpdateQuery:          queries.UpdateByID,
		RepoDeleteQuery:          queries.DeleteByID,
		RepoGetByIdQuery:         queries.GetByID,
		RepoGetAllQuery:          queries.GetAll,
		RepoCountQuery:           queries.Count,
		RepoGetAllPaginatedQuery: queries.GetAllPaginated,
		DtoFieldsCreate:          s.InsertFields(),
		DtoFieldsUpdate:          s.UpdateFields(),
		DtoFieldsFull:            s.FullFields(),
	}, nil
}

func NewRepoQueries(s TableToStructInfo, pk PrimaryKeyView) (RepoQueries, error) {
	insertFields := s.InsertDBFields()
	fieldsWithPkeysAndDefaults := s.ScanDBFields()
	updateSets := GenUpdateSets(s.UpdateFields(), len(s.PrimaryKeys))

	queryTemplatesData := map[string]any{
		"QSchemaName":         s.DbSchemaName,
		"QTableName":          s.DbTableName,
		"QFieldsNoPKeys":      strings.Join(insertFields, ",\n"),
		"QFieldsWithPKeys":    strings.Join(fieldsWithPkeysAndDefaults, ",\n"),
		"QWhereClause":        pk.WhereClause,
		"QOrderClause":        pk.OrderSQL,
		"QInsertPlaceholders": strings.Join(CreatePlaceholders(len(insertFields)), ", "),
		"QUpdateSets":         strings.Join(updateSets, ",\n"),
	}

	exec := func(name string) (string, error) {
		s, err := ExecTemplate(name, queryTemplatesData, FuncMap)
		if err != nil {
			return "", fmt.Errorf("repo query %s: %w", name, err)
		}
		return s, nil
	}

	save, err := exec("query_save")
	if err != nil {
		return RepoQueries{}, err
	}
	update, err := exec("query_update")
	if err != nil {
		return RepoQueries{}, err
	}
	del, err := exec("query_delete")
	if err != nil {
		return RepoQueries{}, err
	}
	getByID, err := exec("query_get_by_id")
	if err != nil {
		return RepoQueries{}, err
	}
	getAll, err := exec("query_get_all")
	if err != nil {
		return RepoQueries{}, err
	}
	count, err := exec("query_count")
	if err != nil {
		return RepoQueries{}, err
	}
	getAllPaginated, err := exec("query_get_all_paginated")
	if err != nil {
		return RepoQueries{}, err
	}

	return RepoQueries{
		Save:            save,
		UpdateByID:      update,
		DeleteByID:      del,
		GetByID:         getByID,
		GetAll:          getAll,
		Count:           count,
		GetAllPaginated: getAllPaginated,
	}, nil
}
