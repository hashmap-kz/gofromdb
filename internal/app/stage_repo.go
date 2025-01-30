package app

import (
	"strings"

	"genpg-v5/internal/tmplts"
)

const (
	// used for update/delete handling
	pkeyDatabaseFieldName = "record_id"
	pkeyStructFieldName   = "RecordID"
)

// Paths example:
// internal/api/client/entity/postgres/client_entity_pg.go
// internal/api/client/repository/client_repository.go
// internal/api/client/repository/impl/client_repository_pg.go

type GenRepo struct {
	RepoEntity    string
	RepoInterface string
	RepoImpl      string
}

func GenRepository(s TableToStructInfo) GenRepo {
	// Entity template
	entityTemplateResult := ExecTemplate("entity", tmplts.EntityTemplate,
		map[string]any{
			"StructName":    s.StructName,
			"StructComment": s.StructComment,
			"Columns":       s.Fields,
		}, FuncMap)

	// Interface template
	interfaceTemplateResult := ExecTemplate("entity-interface", tmplts.RepoInterfaceTemplate,
		map[string]any{
			"StructName":    s.StructName,
			"PackageName":   strings.ToLower(s.DbTableName),
			"InterfaceName": s.StructName + "Repository",
		}, FuncMap)

	// Impl

	fieldsWithoutPkeysAndDefaults := s.GetDbFieldsAsString(false, true)
	fieldsWithPkeysAndDefaults := s.GetDbFieldsAsString(true, false)
	updateSets := GenUpdateSets(fieldsWithoutPkeysAndDefaults)

	queryTemplatesData := map[string]any{
		"SchemaName":                    "public",
		"TableName":                     s.DbTableName,
		"FieldsNoPKeys":                 strings.Join(fieldsWithoutPkeysAndDefaults, ",\n"),
		"FieldsWithPKeys":               strings.Join(fieldsWithPkeysAndDefaults, ",\n"),
		"PkeyFieldName":                 pkeyDatabaseFieldName,
		"ValuesPlaceholders":            strings.Join(CreatePlaceholders(len(fieldsWithoutPkeysAndDefaults)), ", "),
		"FieldsNoPKeysWithPlaceholders": strings.Join(updateSets, ",\n"),
	}

	repoSaveQueryResult := ExecTemplate("query-save", tmplts.RepoSaveQueryTemplate, queryTemplatesData, FuncMap)
	repoUpdateQueryResult := ExecTemplate("query-update", tmplts.RepoUpdateQueryTemplate, queryTemplatesData, FuncMap)
	repoDeleteQueryResult := ExecTemplate("query-delete", tmplts.RepoDeleteQueryTemplate, queryTemplatesData, FuncMap)
	repoGetByIdQueryResult := ExecTemplate("query-get-by-id", tmplts.RepoGetByIdQueryTemplate, queryTemplatesData, FuncMap)
	repoGetAllQueryResult := ExecTemplate("query-get-all", tmplts.RepoGetAllQueryTemplate, queryTemplatesData, FuncMap)
	repoCountQueryResult := ExecTemplate("query-count", tmplts.RepoCountQueryTemplate, queryTemplatesData, FuncMap)
	repoGetAllPaginatedQueryResult := ExecTemplate("query-get-all-paginated", tmplts.RepoGetAllPaginatedQueryTemplate, queryTemplatesData, FuncMap)

	// Function template

	structFieldsWithoutPkeysAndDefaults := s.GetStructFieldsAsString(false, true)
	structFieldsWithPkeysAndDefaults := s.GetStructFieldsAsString(true, false)

	structFieldsUpdate := []string{pkeyStructFieldName}
	structFieldsUpdate = append(structFieldsUpdate, structFieldsWithoutPkeysAndDefaults...)

	repoImplTemplateResult := ExecTemplate("funcs", tmplts.RepoImplTemplate,
		map[string]any{
			"RepoSaveQuery":            repoSaveQueryResult,
			"RepoUpdateQuery":          repoUpdateQueryResult,
			"RepoDeleteQuery":          repoDeleteQueryResult,
			"RepoGetByIdQuery":         repoGetByIdQueryResult,
			"RepoGetAllQuery":          repoGetAllQueryResult,
			"RepoCountQuery":           repoCountQueryResult,
			"RepoGetAllPaginatedQuery": repoGetAllPaginatedQueryResult,
			"StructName":               s.StructName,
			"PackageName":              strings.ToLower(s.DbTableName),
			"InterfaceName":            s.StructName + "Repository",
			"ImplName":                 LowerFirstLetter(s.StructName) + "Repository",
			"StructFieldsUpdate":       structFieldsUpdate,
			"StructFieldsNoPKeys":      structFieldsWithoutPkeysAndDefaults,
			"StructFieldsWithPKeys":    structFieldsWithPkeysAndDefaults,
		}, FuncMap)

	return GenRepo{
		RepoEntity:    PrintFormatted(entityTemplateResult),
		RepoInterface: PrintFormatted(interfaceTemplateResult),
		RepoImpl:      PrintFormatted(repoImplTemplateResult),
	}
}
