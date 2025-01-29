package app

import (
	"strings"
	"text/template"
)

var FuncMap = template.FuncMap{
	"AddPadding":  addPadding,
	"AddPadding2": addPadding2,
	"ToLower":     strings.ToLower,
	"ToCamel":     LowerFirstLetter,
}

func capitalizeFirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}

func LowerFirstLetter(s string) string {
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
