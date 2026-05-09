package import_batches

import "time"

// ImportBatches
// Import batch metadata. Composite natural primary key.
type ImportBatches struct {
	SourceName string     `json:"source_name" db:"source_name"`
	BatchNo    int        `json:"batch_no" db:"batch_no"`
	StartedAt  time.Time  `json:"started_at" db:"started_at"`
	FinishedAt *time.Time `json:"finished_at" db:"finished_at"`
	FileName   string     `json:"file_name" db:"file_name"`
	RowCount   int        `json:"row_count" db:"row_count"`
	Metadata   string     `json:"metadata" db:"metadata"`
}
