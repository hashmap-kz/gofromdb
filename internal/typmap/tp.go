package typmap

import "log"

func GetColTypeForNullable(t string) string {
	postgresToGoNullable := map[string]string{
		// Numeric types
		"int2":        "sql.NullInt16",
		"int4":        "sql.NullInt32",
		"int8":        "sql.NullInt64",
		"numeric":     "sql.NullString", // use sql.NullString if you don't planning calculations. or external libs if you do.
		"decimal":     "sql.NullString", // use sql.NullString if you don't planning calculations. or external libs if you do.
		"serial":      "sql.NullInt32",
		"bigserial":   "sql.NullInt64",
		"smallserial": "sql.NullInt16",
		"float4":      "sql.NullFloat64", // NOTE
		"float8":      "sql.NullFloat64",

		// Character types
		"varchar": "sql.NullString",
		"char":    "sql.NullString",
		"text":    "sql.NullString",
		"bpchar":  "sql.NullString", // Blank-padded char
		"name":    "sql.NullString",

		// Boolean
		"bool": "sql.NullBool",

		// Date/Time types
		"date":        "sql.NullTime",
		"timestamp":   "sql.NullTime",
		"timestamptz": "sql.NullTime", // Timestamp with time zone
		"time":        "sql.NullTime",
		"timetz":      "sql.NullTime", // Time with time zone
		"interval":    "sql.NullString",

		// JSON types
		"json":  "sql.NullString",
		"jsonb": "sql.NullString",

		// UUID
		"uuid": "sql.NullString",

		// Binary data
		"bytea": "[]byte",

		// Arrays
		"_int2":    "[]sql.NullInt16",
		"_int4":    "[]sql.NullInt32",
		"_int8":    "[]sql.NullInt64",
		"_numeric": "[]sql.NullString", // use sql.NullString if you don't planning calculations. or external libs if you do.
		"_text":    "[]sql.NullString",
		"_uuid":    "[]sql.NullString",
		"_bool":    "[]sql.NullBool",
		"_varchar": "[]sql.NullString",
		"_date":    "[]sql.NullTime",
		"_float4":  "[]sql.NullFloat64", // NOTE
		"_float8":  "[]sql.NullFloat64",

		// Other types
		"xml":      "sql.NullString",
		"tsvector": "sql.NullString",
		"tsquery":  "sql.NullString",
		"oid":      "uint32",
		"xid":      "uint32",
		"cid":      "uint32",
		"regclass": "sql.NullString",
		"regproc":  "sql.NullString",
		"regtype":  "sql.NullString",
		"pg_lsn":   "sql.NullString", // Log sequence number
		"record":   "interface{}",
		"void":     "struct{}",
		"unknown":  "sql.NullString",
	}

	if typ, ok := postgresToGoNullable[t]; ok {
		return typ
	}
	log.Fatalf("cannot get type for column: %s", t)
	return ""
}

func GetColType(t string) string {
	// TODO: range types (pgtype.Range[T])

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

		// Binary data
		"bytea": "[]byte",

		// Arrays
		"_int2":    "[]int16",
		"_int4":    "[]int",
		"_int8":    "[]int64",
		"_numeric": "[]string", // use string if you don't planning calculations. or external libs if you do.
		"_text":    "[]string",
		"_uuid":    "[]string",
		"_bool":    "[]bool",
		"_varchar": "[]string",
		"_date":    "[]time.Time",
		"_float4":  "[]float32",
		"_float8":  "[]float64",

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
		"record":   "interface{}",
		"void":     "struct{}",
		"unknown":  "string",
	}
	if typ, ok := postgresToGo[t]; ok {
		return typ
	}
	log.Fatalf("cannot get type for column: %s", t)
	return ""
}
