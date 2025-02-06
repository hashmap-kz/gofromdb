package dto

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type CategoriesDto struct {
	RecordID    int
	Name        string
	ParentID    *int
	ValidPeriod pgtype.Range[time.Time]
	IsCurrent   *bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Guid        string
}

type CategoriesCreateDto struct {
	Name        string
	ParentID    *int
	ValidPeriod pgtype.Range[time.Time]
}

type CategoriesUpdateDto struct {
	Name        string
	ParentID    *int
	ValidPeriod pgtype.Range[time.Time]
}
