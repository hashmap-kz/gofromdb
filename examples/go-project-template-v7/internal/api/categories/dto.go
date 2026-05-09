package categories

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Dto struct {
	RecordID    int
	Name        string
	ParentID    *int
	ValidPeriod pgtype.Range[time.Time]
	IsCurrent   *bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	GUID        string
}

type CreateDto struct {
	Name        string
	ParentID    *int
	ValidPeriod pgtype.Range[time.Time]
}

type UpdateDto struct {
	Name        string
	ParentID    *int
	ValidPeriod pgtype.Range[time.Time]
}
