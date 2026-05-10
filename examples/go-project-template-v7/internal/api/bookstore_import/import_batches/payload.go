package import_batches

import (
	"go-project-template-v7/pkg/pageable"
	"time"
)

// importBatchesCreateRequest
// Import batch metadata. Composite natural primary key.
type importBatchesCreateRequest struct {
	SourceName string     `json:"source_name"`
	BatchNo    int        `json:"batch_no"`
	StartedAt  time.Time  `json:"started_at"`
	FinishedAt *time.Time `json:"finished_at"`
	FileName   string     `json:"file_name"`
	RowCount   int        `json:"row_count"`
	Metadata   string     `json:"metadata"`
}

// importBatchesUpdateRequest
// Import batch metadata. Composite natural primary key.
type importBatchesUpdateRequest struct {
	StartedAt  *time.Time `json:"started_at"`
	FinishedAt *time.Time `json:"finished_at"`
	FileName   *string    `json:"file_name"`
	RowCount   *int       `json:"row_count"`
	Metadata   *string    `json:"metadata"`
}

// importBatchesResponse
// Import batch metadata. Composite natural primary key.
type importBatchesResponse struct {
	SourceName string     `json:"source_name"`
	BatchNo    int        `json:"batch_no"`
	StartedAt  time.Time  `json:"started_at"`
	FinishedAt *time.Time `json:"finished_at"`
	FileName   string     `json:"file_name"`
	RowCount   int        `json:"row_count"`
	Metadata   string     `json:"metadata"`
}

// importBatchesResponseList response list
type importBatchesResponseList struct {
	// Page information (if present)
	Page *pageable.Page `json:"page,omitempty"`

	// Payload
	Data []importBatchesResponse `json:"data"`
}
