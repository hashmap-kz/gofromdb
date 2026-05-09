package app

import (
	"fmt"
	"strings"
)

type PrimaryKeyView struct {
	Fields []TableToStructFieldInfo

	Params      string
	Args        string
	WhereClause string
	OrderClause string
	OrderSQL    string
	URLPath     string
	PathRead    string
}

func NewPrimaryKeyView(fields []TableToStructFieldInfo) (PrimaryKeyView, error) {
	pk := PrimaryKeyView{Fields: fields}
	pk.Params = pkParams(fields)
	pk.Args = pkArgs(fields)
	pk.WhereClause = pkWhereClause(fields)
	pk.OrderClause = pkOrderClause(fields)
	pk.OrderSQL = pkOrderSQL(pk.OrderClause)
	pk.URLPath = pkURLPath(fields)
	pathRead, err := pkPathRead(fields)
	if err != nil {
		return PrimaryKeyView{}, err
	}
	pk.PathRead = pathRead
	return pk, nil
}

func pkURLPath(fields []TableToStructFieldInfo) string {
	clauses := make([]string, 0, len(fields))
	for _, f := range fields {
		clauses = append(clauses, fmt.Sprintf("{%s}", f.DbFieldName))
	}
	return strings.Join(clauses, "/")
}

func pkParams(fields []TableToStructFieldInfo) string {
	clauses := make([]string, 0, len(fields))
	for _, f := range fields {
		clauses = append(clauses, fmt.Sprintf("pk%s %s", f.FieldName, f.FieldType))
	}
	return strings.Join(clauses, ", ")
}

func pkArgs(fields []TableToStructFieldInfo) string {
	clauses := make([]string, 0, len(fields))
	for _, f := range fields {
		clauses = append(clauses, fmt.Sprintf("pk%s", f.FieldName))
	}
	return strings.Join(clauses, ", ")
}

func pkWhereClause(fields []TableToStructFieldInfo) string {
	clauses := make([]string, 0, len(fields))
	for i, f := range fields {
		clauses = append(clauses, fmt.Sprintf("%s = $%d", f.DbFieldName, i+1))
	}
	return strings.Join(clauses, " and ")
}

func pkOrderClause(fields []TableToStructFieldInfo) string {
	clauses := make([]string, 0, len(fields))
	for _, f := range fields {
		clauses = append(clauses, f.DbFieldName)
	}
	return strings.Join(clauses, ", ")
}

func pkOrderSQL(orderClause string) string {
	if orderClause == "" {
		return ""
	}
	return "order by " + orderClause
}

func pkPathRead(fields []TableToStructFieldInfo) (string, error) {
	tmpl := `
	pk%s, err := %s(r, "%s")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}`

	var sb strings.Builder
	for _, f := range fields {
		parser, ok := pathValueParser(f.FieldType)
		if !ok {
			return "", fmt.Errorf(
				"no path parser for primary key %q with Go type %q",
				f.DbFieldName,
				f.FieldType,
			)
		}
		fmt.Fprintf(&sb, tmpl, f.FieldName, parser, f.DbFieldName)
	}
	return sb.String(), nil
}

func pathValueParser(fieldType string) (string, bool) {
	switch fieldType {
	case "int", "int32":
		return "httputils.PathValueI32", true
	case "int16":
		return "httputils.PathValueI16", true
	case "int64":
		return "httputils.PathValueI64", true
	case "uint32":
		return "httputils.PathValueU32", true
	case "string":
		return "httputils.PathValueString", true
	default:
		return "", false
	}
}
