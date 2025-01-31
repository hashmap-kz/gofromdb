package tmplts

var RepoSaveQueryTemplate = `
insert into {{.QSchemaName}}.{{.QTableName}} (
{{.QFieldsNoPKeys | AddPadding}}
)
values ({{.QInsertPlaceholders}})
returning
{{.QFieldsWithPKeys | AddPadding}}
`

var RepoUpdateQueryTemplate = `
update {{.QSchemaName}}.{{.QTableName}}
set 
{{.QUpdateSets | AddPadding}}
where {{.QPkeyFieldName}} = $1
returning 
{{.QFieldsWithPKeys | AddPadding}}
`

var RepoDeleteQueryTemplate = `
delete from only {{.QSchemaName}}.{{.QTableName}}
where {{.QPkeyFieldName}} = $1
`

var RepoGetByIdQueryTemplate = `
select
{{.QFieldsWithPKeys | AddPadding}}
from {{.QSchemaName}}.{{.QTableName}}
where {{.QPkeyFieldName}} = $1
order by {{.QPkeyFieldName}}
`

var RepoGetAllQueryTemplate = `
select
{{.QFieldsWithPKeys | AddPadding}}
from {{.QSchemaName}}.{{.QTableName}}
order by {{.QPkeyFieldName}}
`

var RepoCountQueryTemplate = `select count({{.QPkeyFieldName}}) from {{.QSchemaName}}.{{.QTableName}}`

var RepoGetAllPaginatedQueryTemplate = `
select
{{.QFieldsWithPKeys | AddPadding}}
from {{.QSchemaName}}.{{.QTableName}}
order by {{.QPkeyFieldName}}
offset $1 limit $2
`
