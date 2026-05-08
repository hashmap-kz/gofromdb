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
where {{.QWhereClause}}
returning 
{{.QFieldsWithPKeys | AddPadding}}
`

var RepoDeleteQueryTemplate = `
delete from only {{.QSchemaName}}.{{.QTableName}}
where {{.QWhereClause}}
`

var RepoGetByIdQueryTemplate = `
select
{{.QFieldsWithPKeys | AddPadding}}
from {{.QSchemaName}}.{{.QTableName}}
where {{.QWhereClause}}
{{.QOrderClause}}
`

var RepoGetAllQueryTemplate = `
select
{{.QFieldsWithPKeys | AddPadding}}
from {{.QSchemaName}}.{{.QTableName}}
{{.QOrderClause}}
`

var RepoCountQueryTemplate = `select count(*) from {{.QSchemaName}}.{{.QTableName}}`

var RepoGetAllPaginatedQueryTemplate = `
select
{{.QFieldsWithPKeys | AddPadding}}
from {{.QSchemaName}}.{{.QTableName}}
{{.QOrderClause}}
offset $1 limit $2
`
