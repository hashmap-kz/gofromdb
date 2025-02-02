package app

import (
	"fmt"
	"log"
	"sort"

	"genpg-v5/internal/genpg"
)

var filters = map[string]struct{}{
	"created_at": {},
	"updated_at": {},
	"guid":       {},
}

type Filters struct {
	WithInsertableOnly bool
	WithInternals      bool
}

type TableToStructFieldInfo struct {
	FieldComment   string
	FieldName      string
	FieldType      string
	DbFieldName    string
	DbIsNotNull    bool
	DbNullifRhs    string
	DbIsInsertable bool
}

type TableToStructInfo struct {
	StructName                  string
	StructNameLowerFirstLetter  string
	StructNamePluralRequestPath string
	StructComment               string
	DbTableName                 string
	Fields                      []TableToStructFieldInfo
	PrimaryKeys                 []string
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

func isInternalFieldToSkip(name string) bool {
	if _, ok := filters[name]; ok {
		return true
	}
	return false
}

func (s *TableToStructInfo) GetStructFields(filters Filters) []TableToStructFieldInfo {
	var result []TableToStructFieldInfo
	for _, f := range s.Fields {
		if filters.WithInsertableOnly && !f.DbIsInsertable {
			continue
		}
		if !filters.WithInternals && isInternalFieldToSkip(f.DbFieldName) {
			continue
		}
		result = append(result, f)
	}
	return result
}

func (s *TableToStructInfo) GetDbFieldsAsString(filters Filters) []string {
	var result []string
	for _, f := range s.GetStructFields(filters) {
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
				c.AttType2,
				fmt.Sprintf("%s.%s", c.RelPath, c.AttName),
			)
		}

		fields = append(fields, TableToStructFieldInfo{
			FieldComment:   c.ColDesc,
			FieldName:      makeName(c.AttName),
			FieldType:      c.GoType,
			DbFieldName:    c.AttName,
			DbIsNotNull:    c.AttNotNull,
			DbNullifRhs:    c.NullifRhs,
			DbIsInsertable: c.IsInsertable,
		})
	}

	// TODO: simplify by doing 2 queries? first one to get all tables, and the second one to get all columns?
	// perhaps, but later
	structComment := ""
	primaryKeys := []string{}
	if len(cols) > 0 {
		structComment = cols[0].TabDesc
		primaryKeys = cols[0].PrimaryKeys
	}

	structName := makeName(table)

	return TableToStructInfo{
		StructName:                  structName,
		StructNameLowerFirstLetter:  LowerFirstLetter(structName),
		StructNamePluralRequestPath: makeDnsPathPluralFromDbTable(table),
		StructComment:               structComment,
		DbTableName:                 table,
		Fields:                      fields,
		PrimaryKeys:                 primaryKeys,
	}
}
