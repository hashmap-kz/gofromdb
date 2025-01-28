package typmap

import (
	"log"
	"strings"
)

func GetColType(t string, nullable bool) string {
	// TODO: range types (pgtype.Range[T])

	// [] + baseType
	arrayTypes := map[string]string{
		// Special handling for bytea
		"bytea": "byte",

		// Arrays
		"_int2":    "int16",
		"_int4":    "int",
		"_int8":    "int64",
		"_numeric": "string", // use string if you don't planning calculations. or external libs if you do.
		"_text":    "string",
		"_uuid":    "string",
		"_bool":    "bool",
		"_varchar": "string",
		"_date":    "time.Time",
		"_float4":  "float32",
		"_float8":  "float64",
	}

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
	}

	// handle array types
	if strings.HasPrefix(t, "_") || t == "bytea" {
		if baseType, ok := arrayTypes[t]; ok {
			if nullable {
				return "[]*" + baseType
			}
			return "[]" + baseType
		}
	}

	// handle base types
	if typ, ok := postgresToGo[t]; ok {
		if nullable {
			return "*" + typ
		}
		return typ
	}

	log.Fatalf("cannot get type for column: %s", t)
	return ""
}
