package core

import (
	"fmt"
	"strings"
)

func CreatePlaceholders(cnt int) []string {
	pl := []string{}
	for i := 0; i < cnt; i++ {
		pl = append(pl, fmt.Sprintf("$%d", i+1))
	}
	return pl
}

func maxFieldNameLen(from []TableToStructFieldInfo) int {
	m := 0
	for _, f := range from {
		if len(f.DbFieldName) > m {
			m = len(f.DbFieldName)
		}
	}
	return m
}

func GenUpdateSets(from []TableToStructFieldInfo, pkeysCnt int) []string {
	//nolint:prealloc
	result := []string{}
	maxNameLen := maxFieldNameLen(from)
	for i, field := range from {
		// because $1 is a first parameter (starts with 1, not 0), and also $1 is reserved for ID
		indexOf := i + pkeysCnt + 1

		// v1 --- pattern: `<FIELD_NAME> = coalesce(nullif($1, <RHS_EMPTY_VALUE>::typename), <FIELD_NAME>)`
		// v2 --- nil pointer -> keep existing value; non-nil -> update with the provided value
		updatePattern := fmt.Sprintf("%s %s= coalesce($%d, %s)",
			field.DbFieldName,
			strings.Repeat(" ", maxNameLen-len(field.DbFieldName)), // just a padding
			indexOf,
			field.DbFieldName,
		)

		result = append(result, updatePattern)
	}
	return result
}
