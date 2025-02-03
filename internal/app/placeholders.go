package app

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
	result := []string{}
	maxNameLen := maxFieldNameLen(from)
	for i := 0; i < len(from); i++ {
		structFieldInfo := from[i]

		// because $1 is a first parameter (starts with 1, not 0), and also $1 is reserved for ID
		indexOf := i + pkeysCnt + 1

		// pattern: `<FIELD_NAME> = coalesce(nullif($1, <RHS_EMPTY_VALUE>::typename), <FIELD_NAME>)`
		updatePattern := fmt.Sprintf("%s %s= coalesce(nullif($%d, %s), %s)",
			structFieldInfo.DbFieldName,
			strings.Repeat(" ", maxNameLen-len(structFieldInfo.DbFieldName)), // just a padding
			indexOf,
			structFieldInfo.DbNullifRhs,
			structFieldInfo.DbFieldName,
		)

		result = append(result, updatePattern)
	}
	return result
}
