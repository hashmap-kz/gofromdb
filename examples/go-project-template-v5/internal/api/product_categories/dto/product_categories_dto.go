package dto

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type ProductCategoriesDto struct {
	RecordID    int
	Name        string
	ParentID    *int
	ValidPeriod pgtype.Range[time.Time]
	IsCurrent   *bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Guid        string
}

type ProductCategoriesCreateDto struct {
	Name        string
	ParentID    *int
	ValidPeriod pgtype.Range[time.Time]
}

type ProductCategoriesUpdateDto struct {
	Name        string
	ParentID    *int
	ValidPeriod pgtype.Range[time.Time]
}
