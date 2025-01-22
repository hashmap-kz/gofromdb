package main

import (
	"fmt"
	"genpg-v5/internal/genpg"
	"strings"
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

func getSchemaTable(relPath string) (string, string) {
	r := strings.Split(relPath, ".")
	return r[0], r[1]
}

func p(what string) string {
	return fmt.Sprintf("%s int\n", what)
}

func makeOneStruct(relPath string, cols []genpg.ColumnInfo) string {
	template := strings.TrimSpace(`
type %s struct { 
%s 
}
`)
	sbCols := strings.Builder{}

	_, table := getSchemaTable(relPath)
	structName := makeName(table)

	for _, c := range cols {
		fieldName := makeName(c.AttName)
		sbCols.WriteString(p(fieldName))
	}

	return fmt.Sprintf(template, structName, strings.TrimSpace(sbCols.String()))
}

func main() {
	dbInfo := genpg.GetDBInfo()
	for t := range dbInfo {
		columnInfos := dbInfo[t]
		oneStruct := makeOneStruct(t, columnInfos)
		fmt.Println(oneStruct)
	}
}
