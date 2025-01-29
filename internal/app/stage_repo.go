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

	repoSaveQueryResult := ExecTemplate("query-save", tmplts.RepoSaveQueryTemplate,
		map[string]any{
			"SchemaName":         "public",
			"TableName":          s.DbTableName,
			"FieldsNoPKeys":      strings.Join(fieldsWithoutPkeysAndDefaults, ",\n"),
			"FieldsWithPKeys":    strings.Join(fieldsWithPkeysAndDefaults, ",\n"),
			"ValuesPlaceholders": strings.Join(CreatePlaceholders(len(fieldsWithoutPkeysAndDefaults)), ", "),
		}, FuncMap)

	repoUpdateQueryResult := ExecTemplate("query-update", tmplts.RepoUpdateQueryTemplate,
		map[string]any{
			"SchemaName":                    "public",
			"TableName":                     s.DbTableName,
			"FieldsNoPKeysWithPlaceholders": strings.Join(updateSets, ",\n"),
			"FieldsWithPKeys":               strings.Join(fieldsWithPkeysAndDefaults, ",\n"),
			"PkeyFieldName":                 pkeyDatabaseFieldName,
		}, FuncMap)

	repoDeleteQueryResult := ExecTemplate("query-delete", tmplts.RepoDeleteQueryTemplate,
		map[string]any{
			"SchemaName":    "public",
			"TableName":     s.DbTableName,
			"PkeyFieldName": pkeyDatabaseFieldName,
		}, FuncMap)

	repoGetByIdQueryResult := ExecTemplate("query-get-by-id", tmplts.RepoGetByIdQueryTemplate,
		map[string]any{
			"SchemaName":      "public",
			"TableName":       s.DbTableName,
			"FieldsWithPKeys": strings.Join(fieldsWithPkeysAndDefaults, ",\n"),
			"PkeyFieldName":   pkeyDatabaseFieldName,
		}, FuncMap)

	repoGetAllQueryResult := ExecTemplate("query-get-all", tmplts.RepoGetAllQueryTemplate,
		map[string]any{
			"SchemaName":      "public",
			"TableName":       s.DbTableName,
			"FieldsWithPKeys": strings.Join(fieldsWithPkeysAndDefaults, ",\n"),
			"PkeyFieldName":   pkeyDatabaseFieldName,
		}, FuncMap)

	repoCountQueryResult := ExecTemplate("query-count", tmplts.RepoCountQueryTemplate,
		map[string]any{
			"SchemaName":      "public",
			"TableName":       s.DbTableName,
			"FieldsWithPKeys": strings.Join(fieldsWithPkeysAndDefaults, ",\n"),
			"PkeyFieldName":   pkeyDatabaseFieldName,
		}, FuncMap)

	repoGetAllPaginatedQueryResult := ExecTemplate("query-get-all-paginated", tmplts.RepoGetAllPaginatedQueryTemplate,
		map[string]any{
			"SchemaName":      "public",
			"TableName":       s.DbTableName,
			"FieldsWithPKeys": strings.Join(fieldsWithPkeysAndDefaults, ",\n"),
			"PkeyFieldName":   pkeyDatabaseFieldName,
		}, FuncMap)

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
