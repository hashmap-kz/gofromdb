package app

import (
	"fmt"
	"log"
	"sort"

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
	DbNullifRhs  string
}

type TableToStructInfo struct {
	StructName                  string
	StructNameLowerFirstLetter  string
	StructNamePluralRequestPath string
	StructComment               string
	DbTableName                 string
	Fields                      []TableToStructFieldInfo
}

func GenStructs(connString string) []TableToStructInfo {
	dbInfo := genpg.GetDBInfo(connString)
	var structs []TableToStructInfo

	for t := range dbInfo {
		columnInfos := dbInfo[t]
		oneStruct := makeOneStruct(t, columnInfos)
		structs = append(structs, oneStruct)
	}

	sort.Slice(structs, func(i, j int) bool {
		return structs[i].StructName < structs[j].StructName
	})

	return structs
}

func (s *TableToStructInfo) GetStructFields(withPkeys, skipIfHasDefault bool) []TableToStructFieldInfo {
	var result []TableToStructFieldInfo
	for _, f := range s.Fields {
		if !withPkeys && f.DbIsPk {
			continue
		}
		if skipIfHasDefault && f.DbHasDefault && !f.DbIsPk {
			continue
		}
		result = append(result, f)
	}
	return result
}

func (s *TableToStructInfo) GetStructFieldsAsString(withPkeys, skipIfHasDefault bool) []string {
	var result []string
	for _, f := range s.GetStructFields(withPkeys, skipIfHasDefault) {
		result = append(result, f.FieldName)
	}
	return result
}

func (s *TableToStructInfo) GetDbFieldsAsString(withPkeys, skipIfHasDefault bool) []string {
	var result []string
	for _, f := range s.GetStructFields(withPkeys, skipIfHasDefault) {
		result = append(result, f.DbFieldName)
	}
	return result
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
			DbNullifRhs:  c.NullifRhs,
		})
	}

	structComment := ""
	if len(cols) > 0 {
		structComment = cols[0].TabDesc
	}

	structName := makeName(table)

	return TableToStructInfo{
		StructName:                  structName,
		StructNameLowerFirstLetter:  LowerFirstLetter(structName),
		StructNamePluralRequestPath: makeDnsPathPluralFromDbTable(table),
		StructComment:               structComment,
		DbTableName:                 table,
		Fields:                      fields,
	}
}
