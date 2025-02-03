package app

import "strings"

func getOptionalImports(s TableToStructInfo) string {
	optionalImports := map[string]struct{}{}

	for _, f := range s.Fields {
		if strings.HasPrefix(f.FieldType, "pgtype.") {
			optionalImports["github.com/jackc/pgx/v5/pgtype"] = struct{}{}
		}
		if strings.HasPrefix(f.FieldType, "time.") {
			optionalImports["time"] = struct{}{}
		}
	}

	r := strings.Builder{}
	for k := range optionalImports {
		r.WriteString(k + "\n")
	}
	return r.String()
}
