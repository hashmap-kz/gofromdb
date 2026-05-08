package app

import (
	"strings"

	"genpg-v5/internal/tmplts"
)

type RepoTemplateData struct {
	StructName           string
	StructComment        string
	PackageName          string
	InterfaceName        string
	ImplName             string
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

func NewRepoTemplateData(s TableToStructInfo) RepoTemplateData {
	pk := NewPrimaryKeyView(s.PrimaryKeys)
	queries := NewRepoQueries(s, pk)

	return RepoTemplateData{
		StructName:               s.StructName,
		StructComment:            s.StructComment,
		PackageName:              strings.ToLower(s.DbTableName),
		InterfaceName:            s.StructName + "Repository",
		ImplName:                 LowerFirstLetter(s.StructName) + "Repository",
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
	}
}

func NewRepoQueries(s TableToStructInfo, pk PrimaryKeyView) RepoQueries {
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

	return RepoQueries{
		Save:            ExecTemplate("query-save", tmplts.RepoSaveQueryTemplate, queryTemplatesData, FuncMap),
		UpdateByID:      ExecTemplate("query-update", tmplts.RepoUpdateQueryTemplate, queryTemplatesData, FuncMap),
		DeleteByID:      ExecTemplate("query-delete", tmplts.RepoDeleteQueryTemplate, queryTemplatesData, FuncMap),
		GetByID:         ExecTemplate("query-get-by-id", tmplts.RepoGetByIdQueryTemplate, queryTemplatesData, FuncMap),
		GetAll:          ExecTemplate("query-get-all", tmplts.RepoGetAllQueryTemplate, queryTemplatesData, FuncMap),
		Count:           ExecTemplate("query-count", tmplts.RepoCountQueryTemplate, queryTemplatesData, FuncMap),
		GetAllPaginated: ExecTemplate("query-get-all-paginated", tmplts.RepoGetAllPaginatedQueryTemplate, queryTemplatesData, FuncMap),
	}
}
