package import_errors

import "time"

// ImportErrors
// Import validation errors. Intentionally has no primary key.
type ImportErrors struct {
	SourceName string    `json:"source_name" db:"source_name"`
	BatchNo    int       `json:"batch_no" db:"batch_no"`
	RowNo      int       `json:"row_no" db:"row_no"`
	ColumnName *string   `json:"column_name" db:"column_name"`
	Message    string    `json:"message" db:"message"`
	RawPayload string    `json:"raw_payload" db:"raw_payload"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}
