package import_errors

import (
	"go-project-template-v7/pkg/pageable"
	"time"
)

// importErrorsCreateRequest
// Import validation errors. Intentionally has no primary key.
type importErrorsCreateRequest struct {
	SourceName string  `json:"source_name"`
	BatchNo    int     `json:"batch_no"`
	RowNo      int     `json:"row_no"`
	ColumnName *string `json:"column_name"`
	Message    string  `json:"message"`
	RawPayload string  `json:"raw_payload"`
}

// importErrorsUpdateRequest
// Import validation errors. Intentionally has no primary key.
type importErrorsUpdateRequest struct {
	SourceName string  `json:"source_name"`
	BatchNo    int     `json:"batch_no"`
	RowNo      int     `json:"row_no"`
	ColumnName *string `json:"column_name"`
	Message    string  `json:"message"`
	RawPayload string  `json:"raw_payload"`
}

// importErrorsResponse
// Import validation errors. Intentionally has no primary key.
type importErrorsResponse struct {
	SourceName string    `json:"source_name"`
	BatchNo    int       `json:"batch_no"`
	RowNo      int       `json:"row_no"`
	ColumnName *string   `json:"column_name"`
	Message    string    `json:"message"`
	RawPayload string    `json:"raw_payload"`
	CreatedAt  time.Time `json:"created_at"`
}

// importErrorsResponseList response list
type importErrorsResponseList struct {
	// Page information (if present)
	Page *pageable.Page `json:"page,omitempty"`

	// Payload
	Data []importErrorsResponse `json:"data"`
}
