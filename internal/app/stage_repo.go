package app

import (
	"strings"

	"genpg-v5/internal/tmplts"
)

type GenRepo struct {
	RepoEntity    string
	RepoInterface string
	RepoImpl      string
}

type queriesResults struct {
	repoSaveQueryResult            string
	repoUpdateQueryResult          string
	repoDeleteQueryResult          string
	repoGetByIdQueryResult         string
	repoGetAllQueryResult          string
	repoCountQueryResult           string
	repoGetAllPaginatedQueryResult string
}

func GenRepository(s TableToStructInfo) GenRepo {
	parametersByPkeys := genParametersByPkeys(s)
	argumentsByPkeys := genArgumentsByPkeys(s)

	queries := genQueries(s)
	data := map[string]any{
		"StructName":               s.StructName,
		"StructComment":            s.StructComment,
		"PackageName":              strings.ToLower(s.DbTableName),
		"InterfaceName":            s.StructName + "Repository",
		"ImplName":                 LowerFirstLetter(s.StructName) + "Repository",
		"ParametersByPkeys":        parametersByPkeys,
		"ArgumentsByPkeys":         argumentsByPkeys,
		"RepoSaveQuery":            queries.repoSaveQueryResult,
		"RepoUpdateQuery":          queries.repoUpdateQueryResult,
		"RepoDeleteQuery":          queries.repoDeleteQueryResult,
		"RepoGetByIdQuery":         queries.repoGetByIdQueryResult,
		"RepoGetAllQuery":          queries.repoGetAllQueryResult,
		"RepoCountQuery":           queries.repoCountQueryResult,
		"RepoGetAllPaginatedQuery": queries.repoGetAllPaginatedQueryResult,
		"DtoFieldsCreate": s.GetStructFields(Filters{
			WithInsertableOnly: true,
			WithInternals:      false,
		}),
		"DtoFieldsUpdate": s.GetStructFields(Filters{
			WithInsertableOnly: true,
			WithInternals:      false,
			WithoutPrimaryKeys: true,
		}),
		"HasPrimaryKey":   s.HasPrimaryKey,
		"HasUpdateFields": s.HasUpdateFields,
		"DtoFieldsFull": s.GetStructFields(Filters{
			WithInsertableOnly: false,
			WithInternals:      true,
		}),
	}

	interfaceRes := ExecTemplate("entity-interface", tmplts.RepoInterfaceTemplate, data, FuncMap)
	modelsRes := ExecTemplate("entity", tmplts.EntityTemplate, data, FuncMap)
	implRes := ExecTemplate("funcs", tmplts.RepoImplTemplate, data, FuncMap)

	return GenRepo{
		RepoEntity:    PrintFormatted(modelsRes),
		RepoInterface: PrintFormatted(interfaceRes),
		RepoImpl:      PrintFormatted(implRes),
	}
}

func genQueries(s TableToStructInfo) queriesResults {
	insertFields := s.GetDbFieldsAsString(Filters{
		WithInsertableOnly: true,
		WithInternals:      false,
	})

	fieldsWithPkeysAndDefaults := s.GetDbFieldsAsString(Filters{
		WithInsertableOnly: false,
		WithInternals:      true,
	})

	updateSets := GenUpdateSets(s.GetStructFields(Filters{
		WithInsertableOnly: true,
		WithInternals:      false,
		WithoutPrimaryKeys: true,
	}), len(s.PrimaryKeys))

	queryTemplatesData := map[string]any{
		"QSchemaName":         s.DbSchemaName,
		"QTableName":          s.DbTableName,
		"QFieldsNoPKeys":      strings.Join(insertFields, ",\n"),
		"QFieldsWithPKeys":    strings.Join(fieldsWithPkeysAndDefaults, ",\n"),
		"QWhereClause":        genWhereClauseByPkeys(s),
		"QOrderClause":        genOrderBySQLByPkeys(s),
		"QInsertPlaceholders": strings.Join(CreatePlaceholders(len(insertFields)), ", "),
		"QUpdateSets":         strings.Join(updateSets, ",\n"),
	}

	repoSaveQueryResult := ExecTemplate("query-save", tmplts.RepoSaveQueryTemplate, queryTemplatesData, FuncMap)
	repoUpdateQueryResult := ExecTemplate("query-update", tmplts.RepoUpdateQueryTemplate, queryTemplatesData, FuncMap)
	repoDeleteQueryResult := ExecTemplate("query-delete", tmplts.RepoDeleteQueryTemplate, queryTemplatesData, FuncMap)
	repoGetByIdQueryResult := ExecTemplate("query-get-by-id", tmplts.RepoGetByIdQueryTemplate, queryTemplatesData, FuncMap)
	repoGetAllQueryResult := ExecTemplate("query-get-all", tmplts.RepoGetAllQueryTemplate, queryTemplatesData, FuncMap)
	repoCountQueryResult := ExecTemplate("query-count", tmplts.RepoCountQueryTemplate, queryTemplatesData, FuncMap)
	repoGetAllPaginatedQueryResult := ExecTemplate("query-get-all-paginated", tmplts.RepoGetAllPaginatedQueryTemplate, queryTemplatesData, FuncMap)

	return queriesResults{
		repoSaveQueryResult:            repoSaveQueryResult,
		repoUpdateQueryResult:          repoUpdateQueryResult,
		repoDeleteQueryResult:          repoDeleteQueryResult,
		repoGetByIdQueryResult:         repoGetByIdQueryResult,
		repoGetAllQueryResult:          repoGetAllQueryResult,
		repoCountQueryResult:           repoCountQueryResult,
		repoGetAllPaginatedQueryResult: repoGetAllPaginatedQueryResult,
	}
}
