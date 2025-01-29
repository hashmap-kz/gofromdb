package app

import (
	"fmt"
	"log"

	"genpg-v5/internal/genpg"
)

type TableToStructFieldInfo struct {
	FieldComment string
	FieldName    string
	FieldType    string
	DbFieldName  string
	DbIsNotNull  bool
	DbIsPk       bool
	DbHasDefault bool
}

type TableToStructInfo struct {
	StructName    string
	StructComment string
	DbTableName   string
	Fields        []TableToStructFieldInfo
}

func GenStructs() []TableToStructInfo {
	dbInfo := genpg.GetDBInfo()
	var structs []TableToStructInfo

	for t := range dbInfo {
		columnInfos := dbInfo[t]
		oneStruct := makeOneStruct(t, columnInfos)
		structs = append(structs, oneStruct)
	}

	return structs
}

func (s *TableToStructInfo) GetDbFieldsAsString(withPkeys, skipIfHasDefault bool) []string {
	r := []string{}
	for _, f := range s.Fields {
		if !withPkeys && f.DbIsPk {
			continue
		}
		if skipIfHasDefault && f.DbHasDefault {
			continue
		}
		r = append(r, f.DbFieldName)
	}
	return r
}

func (s *TableToStructInfo) GetStructFieldsAsString(withPkeys, skipIfHasDefault bool) []string {
	r := []string{}
	for _, f := range s.Fields {
		if !withPkeys && f.DbIsPk {
			continue
		}
		if skipIfHasDefault && f.DbHasDefault {
			continue
		}
		r = append(r, f.FieldName)
	}
	return r
}

func makeOneStruct(relPath string, cols []genpg.ColumnInfo) TableToStructInfo {
	fields := []TableToStructFieldInfo{}

	_, table := getSchemaTable(relPath)
	for _, c := range cols {
		if c.GoType == "" {
			log.Fatalf("cannot find type mapping for pg-type: `%s`, column: `%s`",
				c.AttType,
				fmt.Sprintf("%s.%s", c.RelPath, c.AttName),
			)
		}

		fields = append(fields, TableToStructFieldInfo{
			FieldComment: c.ColDesc,
			FieldName:    makeName(c.AttName),
			FieldType:    c.GoType,
			DbFieldName:  c.AttName,
			DbIsNotNull:  c.AttNotNull,
			DbIsPk:       c.IsPK,
			DbHasDefault: c.Def != nil,
		})
	}

	structComment := ""
	if len(cols) > 0 {
		structComment = cols[0].TabDesc
	}

	return TableToStructInfo{
		StructName:    makeName(table),
		StructComment: structComment,
		DbTableName:   table,
		Fields:        fields,
	}
}
