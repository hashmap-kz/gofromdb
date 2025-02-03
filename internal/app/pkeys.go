package app

import (
	"fmt"
	"strings"
)

func genUrlPathValuesByPkeys(pkeys []TableToStructFieldInfo) string {
	// /api/v1/categories/{id}/{code}
	clauses := []string{}
	for _, pk := range pkeys {
		clause := fmt.Sprintf("{%s}", pk.DbFieldName)
		clauses = append(clauses, clause)
	}
	return strings.Join(clauses, "/")
}

func genParametersByPkeys(s TableToStructInfo) string {
	// pkCode int, pkID int
	clauses := []string{}
	for _, pk := range s.PrimaryKeys {
		clause := fmt.Sprintf("pk%s %s", pk.FieldName, pk.FieldType)
		clauses = append(clauses, clause)
	}
	return strings.Join(clauses, ", ")
}

func genArgumentsByPkeys(s TableToStructInfo) string {
	// pkCode, pkID
	clauses := []string{}
	for _, pk := range s.PrimaryKeys {
		clause := fmt.Sprintf("pk%s", pk.FieldName)
		clauses = append(clauses, clause)
	}
	return strings.Join(clauses, ", ")
}

func genWhereClauseByPkeys(s TableToStructInfo) string {
	// where pk1 = $1 and pk2 = $2 ...
	clauses := []string{}
	for i, pk := range s.PrimaryKeys {
		clause := fmt.Sprintf("%s = $%d", pk.DbFieldName, i+1)
		clauses = append(clauses, clause)
	}
	return strings.Join(clauses, " and ")
}

func genOrderByClauseByPkeys(s TableToStructInfo) string {
	// order by pk1, pk2, ...
	clauses := []string{}
	for _, pk := range s.PrimaryKeys {
		clauses = append(clauses, pk.DbFieldName)
	}
	return strings.Join(clauses, ", ")
}
