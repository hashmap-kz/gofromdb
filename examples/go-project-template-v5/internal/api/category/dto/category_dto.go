package dto

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type CategoryDto struct {
	RecordID    int
	Name        string
	ParentID    *int
	ValidPeriod pgtype.Range[time.Time]
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Guid        string
}

type CategoryCreateDto struct {
	Name        string
	ParentID    *int
	ValidPeriod pgtype.Range[time.Time]
}

type CategoryUpdateDto struct {
	Name        string
	ParentID    *int
	ValidPeriod pgtype.Range[time.Time]
}
