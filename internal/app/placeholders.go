package app

import "fmt"

func CreatePlaceholders(cnt int) []string {
	pl := []string{}
	for i := 0; i < cnt; i++ {
		pl = append(pl, fmt.Sprintf("$%d", i+1))
	}
	return pl
}

func GenUpdateSets(from []TableToStructFieldInfo) []string {
	result := []string{}
	for i := 0; i < len(from); i++ {
		structFieldInfo := from[i]

		// because $1 is a first parameter (starts with 1, not 0), and also $1 is reserved for ID
		indexOf := i + 2

		// pattern: `<FIELD_NAME> = coalesce(nullif($1, <RHS_EMPTY_VALUE>::typename), <FIELD_NAME>)`
		updatePattern := fmt.Sprintf("%s = coalesce(nullif($%d, %s), %s)",
			structFieldInfo.DbFieldName,
			indexOf,
			structFieldInfo.DbNullifRhs,
			structFieldInfo.DbFieldName,
		)

		result = append(result, updatePattern)
	}
	return result
}
