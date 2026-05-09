package app

import (
	"fmt"
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
	WithoutPrimaryKeys bool
}

type TableToStructFieldInfo struct {
	FieldComment   string
	FieldName      string
	FieldType      string
	DbFieldName    string
	DbIsNotNull    bool
	DbNullifRhs    string
	DbIsInsertable bool
	DbIsPrimaryKey bool
}

type TableToStructInfo struct {
	StructName                  string
	StructNameLowerFirstLetter  string
	StructNamePluralRequestPath string
	PkeysURLPath                string
	StructComment               string
	DbSchemaName                string
	DbTableName                 string
	Fields                      []TableToStructFieldInfo
	PrimaryKeys                 []TableToStructFieldInfo
	HasPrimaryKey               bool
	HasUpdateFields             bool
}

func GenStructs(connString string) ([]TableToStructInfo, error) {
	dbInfo, err := genpg.GetDBInfo(connString)
	if err != nil {
		return nil, fmt.Errorf("get db info: %w", err)
	}
	var structs []TableToStructInfo

	for t := range dbInfo {
		oneStruct, err := makeOneStruct(t, dbInfo[t])
		if err != nil {
			return nil, fmt.Errorf("make struct for %s: %w", t, err)
		}
		structs = append(structs, oneStruct)
	}

	sort.Slice(structs, func(i, j int) bool {
		return structs[i].StructName < structs[j].StructName
	})

	return structs, nil
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
		if filters.WithoutPrimaryKeys && f.DbIsPrimaryKey {
			continue
		}
		result = append(result, f)
	}
	return result
}

func (s *TableToStructInfo) GetDbFieldsAsString(filters Filters) []string {
	return dbFieldNames(s.GetStructFields(filters))
}

func (s *TableToStructInfo) FullFields() []TableToStructFieldInfo {
	return s.GetStructFields(Filters{
		WithInsertableOnly: false,
		WithInternals:      true,
	})
}

func (s *TableToStructInfo) InsertFields() []TableToStructFieldInfo {
	return s.GetStructFields(Filters{
		WithInsertableOnly: true,
		WithInternals:      false,
	})
}

func (s *TableToStructInfo) UpdateFields() []TableToStructFieldInfo {
	return s.GetStructFields(Filters{
		WithInsertableOnly: true,
		WithInternals:      false,
		WithoutPrimaryKeys: true,
	})
}

func (s *TableToStructInfo) ScanFields() []TableToStructFieldInfo {
	return s.FullFields()
}

func (s *TableToStructInfo) InsertDBFields() []string {
	return dbFieldNames(s.InsertFields())
}

func (s *TableToStructInfo) ScanDBFields() []string {
	return dbFieldNames(s.ScanFields())
}

func dbFieldNames(fields []TableToStructFieldInfo) []string {
	result := make([]string, 0, len(fields))
	for _, f := range fields {
		result = append(result, f.DbFieldName)
	}
	return result
}

func makeOneStruct(relPath string, cols []genpg.ColumnInfo) (TableToStructInfo, error) {
	fields := []TableToStructFieldInfo{}

	schema, table := getSchemaTable(relPath)
	for _, c := range cols {
		if c.GoType == "" {
			return TableToStructInfo{}, fmt.Errorf(
				"no Go type mapping for pg type %q, column %s.%s",
				c.AttType2,
				c.RelPath,
				c.AttName,
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

	structComment := ""
	primaryKeys := []TableToStructFieldInfo{}
	if len(cols) > 0 {
		structComment = cols[0].TabDesc
		primaryKeys = handlePkeys(fields, cols[0].PrimaryKeys)
	}

	pkNames := map[string]struct{}{}
	for _, pk := range primaryKeys {
		pkNames[pk.DbFieldName] = struct{}{}
	}
	for i := range fields {
		_, fields[i].DbIsPrimaryKey = pkNames[fields[i].DbFieldName]
	}
	primaryKeys = handlePkeys(fields, cols[0].PrimaryKeys)

	pkView, err := NewPrimaryKeyView(primaryKeys)
	if err != nil {
		return TableToStructInfo{}, fmt.Errorf("primary key view for %s: %w", table, err)
	}

	structName := makeName(table)
	info := TableToStructInfo{
		StructName:                  structName,
		StructNameLowerFirstLetter:  LowerFirstLetter(structName),
		StructNamePluralRequestPath: makeDnsPathPluralFromDbTable(table),
		PkeysURLPath:                pkView.URLPath,
		StructComment:               structComment,
		DbSchemaName:                schema,
		DbTableName:                 table,
		Fields:                      fields,
		PrimaryKeys:                 primaryKeys,
		HasPrimaryKey:               len(primaryKeys) > 0,
	}
	info.HasUpdateFields = info.HasPrimaryKey && len(info.UpdateFields()) > 0
	return info, nil
}

func handlePkeys(fields []TableToStructFieldInfo, keys []string) []TableToStructFieldInfo {
	r := []TableToStructFieldInfo{}
	for _, pkName := range keys {
		for _, field := range fields {
			if pkName == field.DbFieldName {
				r = append(r, field)
			}
		}
	}
	return r
}
