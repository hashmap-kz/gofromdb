package import_errors

import "time"

type Dto struct {
	SourceName string
	BatchNo    int
	RowNo      int
	ColumnName *string
	Message    string
	RawPayload string
	CreatedAt  time.Time
}

type CreateDto struct {
	SourceName string
	BatchNo    int
	RowNo      int
	ColumnName *string
	Message    string
	RawPayload string
}

type UpdateDto struct {
	SourceName *string
	BatchNo    *int
	RowNo      *int
	ColumnName *string
	Message    *string
	RawPayload *string
}
