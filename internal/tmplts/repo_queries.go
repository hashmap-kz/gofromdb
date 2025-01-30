package tmplts

var RepoSaveQueryTemplate = `
insert into {{.SchemaName}}.{{.TableName}} (
{{.FieldsNoPKeys | AddPadding}}
)
values ({{.ValuesPlaceholders}})
returning
{{.FieldsWithPKeys | AddPadding}}
`

var RepoUpdateQueryTemplate = `
update {{.SchemaName}}.{{.TableName}}
set 
{{.FieldsNoPKeysWithPlaceholders | AddPadding}}
where {{.PkeyFieldName}} = $1
returning 
{{.FieldsWithPKeys | AddPadding}}
`

var RepoDeleteQueryTemplate = `
delete from only {{.SchemaName}}.{{.TableName}}
where {{.PkeyFieldName}} = $1
`

var RepoGetByIdQueryTemplate = `
select
{{.FieldsWithPKeys | AddPadding}}
from {{.SchemaName}}.{{.TableName}}
where {{.PkeyFieldName}} = $1
order by {{.PkeyFieldName}}
`

var RepoGetAllQueryTemplate = `
select
{{.FieldsWithPKeys | AddPadding}}
from {{.SchemaName}}.{{.TableName}}
order by {{.PkeyFieldName}}
`

var RepoCountQueryTemplate = `select count({{.PkeyFieldName}}) from {{.SchemaName}}.{{.TableName}}`

var RepoGetAllPaginatedQueryTemplate = `
select
{{.FieldsWithPKeys | AddPadding}}
from {{.SchemaName}}.{{.TableName}}
order by {{.PkeyFieldName}}
offset $1 limit $2
`
