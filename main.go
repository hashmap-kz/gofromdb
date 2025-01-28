package main

import (
	"fmt"
	"log"
	"strings"

	"genpg-v5/internal/genpg"
)

func capitalizeFirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}

func makeName(from string) string {
	r := strings.Split(from, "_")
	sb := strings.Builder{}
	for _, elem := range r {
		if strings.ToLower(elem) == "id" {
			elem = "ID"
		}
		sb.WriteString(capitalizeFirstLetter(elem))
	}
	return sb.String()
}

func getColType(t string) string {
	// TODO: range types (pgtype.Range[T])

	postgresToGo := map[string]string{
		// Numeric types
		"int2":        "int16",
		"int4":        "int",
		"int8":        "int64",
		"numeric":     "string", // use string if you don't planning calculations. or external libs if you do.
		"decimal":     "string", // use string if you don't planning calculations. or external libs if you do.
		"serial":      "int",
		"bigserial":   "int64",
		"smallserial": "int16",
		"float4":      "float32",
		"float8":      "float64",

		// Character types
		"varchar": "string",
		"char":    "string",
		"text":    "string",
		"bpchar":  "string", // Blank-padded char
		"name":    "string",

		// Boolean
		"bool": "bool",

		// Date/Time types
		"date":        "time.Time",
		"timestamp":   "time.Time",
		"timestamptz": "time.Time", // Timestamp with time zone
		"time":        "time.Time",
		"timetz":      "time.Time", // Time with time zone
		"interval":    "string",

		// JSON types
		"json":  "string",
		"jsonb": "string",

		// UUID
		"uuid": "string",

		// Binary data
		"bytea": "[]byte",

		// Arrays
		"_int2":    "[]int16",
		"_int4":    "[]int",
		"_int8":    "[]int64",
		"_numeric": "[]string", // use string if you don't planning calculations. or external libs if you do.
		"_text":    "[]string",
		"_uuid":    "[]string",
		"_bool":    "[]bool",
		"_varchar": "[]string",
		"_date":    "[]time.Time",
		"_float4":  "[]float32",
		"_float8":  "[]float64",

		// Other types
		"xml":      "string",
		"tsvector": "string",
		"tsquery":  "string",
		"oid":      "uint32",
		"xid":      "uint32",
		"cid":      "uint32",
		"regclass": "string",
		"regproc":  "string",
		"regtype":  "string",
		"pg_lsn":   "string", // Log sequence number
		"record":   "interface{}",
		"void":     "struct{}",
		"unknown":  "string",
	}
	if typ, ok := postgresToGo[t]; ok {
		return typ
	}
	log.Fatalf("cannot get type for column: %s", t)
	return ""
}

func getSchemaTable(relPath string) (string, string) {
	r := strings.Split(relPath, ".")
	return r[0], r[1]
}

func p(what string) string {
	return fmt.Sprintf("\t%s\n", what)
}

type TableToStructFieldInfo struct {
	FieldComment string
	FieldName    string
	FieldType    string
	DbFieldName  string
	DbIsNotNull  bool
}

type TableToStructInfo struct {
	StructName string
	Fields     []TableToStructFieldInfo
}

func (s *TableToStructInfo) getDbFieldsAsString() []string {
	r := []string{}
	for _, f := range s.Fields {
		r = append(r, f.DbFieldName)
	}
	return r
}

func makeOneStruct(relPath string, cols []genpg.ColumnInfo) TableToStructInfo {
	fields := []TableToStructFieldInfo{}

	_, table := getSchemaTable(relPath)
	for _, c := range cols {
		fields = append(fields, TableToStructFieldInfo{
			FieldComment: c.ColDesc,
			FieldName:    makeName(c.AttName),
			FieldType:    getColType(c.AttType2),
			DbFieldName:  c.AttName,
			DbIsNotNull:  c.AttNotNull,
		})
	}

	return TableToStructInfo{
		StructName: makeName(table),
		Fields:     fields,
	}
}

func main() {
	dbInfo := genpg.GetDBInfo()
	structs := []TableToStructInfo{}
	for t := range dbInfo {
		columnInfos := dbInfo[t]
		oneStruct := makeOneStruct(t, columnInfos)
		structs = append(structs, oneStruct)
	}
	for _, s := range structs {
		dbFieldsAsString := s.getDbFieldsAsString()
		fmt.Println(strings.Join(dbFieldsAsString, ","))
	}
}
