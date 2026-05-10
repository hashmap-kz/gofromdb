package import_batches

import "time"

type Dto struct {
	SourceName string
	BatchNo    int
	StartedAt  time.Time
	FinishedAt *time.Time
	FileName   string
	RowCount   int
	Metadata   string
}

type CreateDto struct {
	SourceName string
	BatchNo    int
	StartedAt  time.Time
	FinishedAt *time.Time
	FileName   string
	RowCount   int
	Metadata   string
}

type UpdateDto struct {
	StartedAt  *time.Time
	FinishedAt *time.Time
	FileName   *string
	RowCount   *int
	Metadata   *string
}
