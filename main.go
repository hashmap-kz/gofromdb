package main

import (
	"bytes"
	"fmt"
	"go/format"
	"log"
	"strings"
	"text/template"

	"genpg-v5/internal/genpg"
	"genpg-v5/internal/tmplts"
)

const (
	// used for update/delete handling
	pkeyDatabaseFieldName = "record_id"
	pkeyStructFieldName   = "RecordID"
)

func capitalizeFirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}

func lowerFirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToLower(string(s[0])) + s[1:]
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
	DbIsPk       bool
	DbHasDefault bool
}

type TableToStructInfo struct {
	StructName    string
	StructComment string
	DbTableName   string
	Fields        []TableToStructFieldInfo
}

func (s *TableToStructInfo) getDbFieldsAsString(withPkeys, skipIfHasDefault bool) []string {
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

func (s *TableToStructInfo) getStructFieldsAsString(withPkeys, skipIfHasDefault bool) []string {
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

func createPlaceholders(cnt int) []string {
	pl := []string{}
	for i := 0; i < cnt; i++ {
		pl = append(pl, fmt.Sprintf("$%d", i+1))
	}
	return pl
}

// Function to add padding (tabs) to each line
func addPadding(input string) string {
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		lines[i] = "\t" + line // Add a tab before each line
	}
	return strings.Join(lines, "\n")
}

// Function to add padding (tabs) to each line
func addPadding2(input string) string {
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		lines[i] = "\t\t" + line // Add a tab before each line
	}
	return strings.Join(lines, "\n")
}

func execTemplate(name, t string, data map[string]any, funcMap map[string]any) string {
	var result bytes.Buffer
	tmpl, err := template.New(name).Funcs(funcMap).Parse(t)
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(&result, data)
	if err != nil {
		log.Fatal(err)
	}
	return result.String()
}

func formatGoCode(input string) (string, error) {
	// Convert the string to a []byte
	source := []byte(input)

	// Format the code using go/format
	formattedSource, err := format.Source(source)
	if err != nil {
		return "", fmt.Errorf("failed to format code: %w", err)
	}

	// Convert the formatted code back to a string
	return string(formattedSource), nil
}

func printFormatted(input string) string {
	code, err := formatGoCode(input)
	if err != nil {
		return input
	}
	return code
}

func genUpdateSets(from []string) []string {
	result := []string{}
	for i := 0; i < len(from); i++ {
		indexOf := i + 2 // because $1 is a first parameter (starts with 1, not 0), and also $1 is reserved for ID
		result = append(result, from[i]+" = $"+fmt.Sprintf("%d", indexOf))
	}
	return result
}

func main() {
	dbInfo := genpg.GetDBInfo()
	structs := []TableToStructInfo{}

	for t := range dbInfo {
		columnInfos := dbInfo[t]
		oneStruct := makeOneStruct(t, columnInfos)
		structs = append(structs, oneStruct)
	}

	funcMap := template.FuncMap{
		"AddPadding":  addPadding,
		"AddPadding2": addPadding2,
		"ToLower":     strings.ToLower,
		"ToCamel":     lowerFirstLetter,
	}

	for _, s := range structs {

		// Entity template
		entityTemplateResult := execTemplate("entity", tmplts.EntityTemplate,
			map[string]any{
				"StructName":    s.StructName,
				"StructComment": s.StructComment,
				"Columns":       s.Fields,
			}, funcMap)
		fmt.Println(printFormatted(entityTemplateResult))

		// Insert template

		fieldsWithoutPkeysAndDefaults := s.getDbFieldsAsString(false, true)
		fieldsWithPkeysAndDefaults := s.getDbFieldsAsString(true, false)
		updateSets := genUpdateSets(fieldsWithoutPkeysAndDefaults)

		repoSaveQueryResult := execTemplate("query-save", tmplts.RepoSaveQueryTemplate,
			map[string]any{
				"SchemaName":         "public",
				"TableName":          s.DbTableName,
				"FieldsNoPKeys":      strings.Join(fieldsWithoutPkeysAndDefaults, ",\n"),
				"FieldsWithPKeys":    strings.Join(fieldsWithPkeysAndDefaults, ",\n"),
				"ValuesPlaceholders": strings.Join(createPlaceholders(len(fieldsWithoutPkeysAndDefaults)), ", "),
			}, funcMap)

		repoUpdateQueryResult := execTemplate("query-update", tmplts.RepoUpdateQueryTemplate,
			map[string]any{
				"SchemaName":                    "public",
				"TableName":                     s.DbTableName,
				"FieldsNoPKeysWithPlaceholders": strings.Join(updateSets, ",\n"),
				"FieldsWithPKeys":               strings.Join(fieldsWithPkeysAndDefaults, ",\n"),
				"PkeyFieldName":                 pkeyDatabaseFieldName,
			}, funcMap)

		repoDeleteQueryResult := execTemplate("query-delete", tmplts.RepoDeleteQueryTemplate,
			map[string]any{
				"SchemaName":    "public",
				"TableName":     s.DbTableName,
				"PkeyFieldName": pkeyDatabaseFieldName,
			}, funcMap)

		// Function template

		structFieldsWithoutPkeysAndDefaults := s.getStructFieldsAsString(false, true)
		structFieldsWithPkeysAndDefaults := s.getStructFieldsAsString(true, false)

		structFieldsUpdate := []string{pkeyStructFieldName}
		structFieldsUpdate = append(structFieldsUpdate, structFieldsWithoutPkeysAndDefaults...)

		result := execTemplate("funcs", tmplts.RepoImplTemplate,
			map[string]any{
				"RepoSaveQuery":         repoSaveQueryResult,
				"RepoUpdateQuery":       repoUpdateQueryResult,
				"RepoDeleteQuery":       repoDeleteQueryResult,
				"StructName":            s.StructName,
				"PackageName":           strings.ToLower(s.DbTableName),
				"InterfaceName":         s.StructName + "Repository",
				"ImplName":              lowerFirstLetter(s.StructName) + "Repository",
				"StructFieldsUpdate":    structFieldsUpdate,
				"StructFieldsNoPKeys":   structFieldsWithoutPkeysAndDefaults,
				"StructFieldsWithPKeys": structFieldsWithPkeysAndDefaults,
			}, funcMap)

		fmt.Println(printFormatted(result))
	}
}
