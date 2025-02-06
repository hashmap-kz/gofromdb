package dto

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type PurchaseStepsDto struct {
	RecordID    int
	ValidPeriod pgtype.Range[time.Time]
	BuyID       int
	StepID      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Guid        string
}

type PurchaseStepsCreateDto struct {
	ValidPeriod pgtype.Range[time.Time]
	BuyID       int
	StepID      int
}

type PurchaseStepsUpdateDto struct {
	ValidPeriod pgtype.Range[time.Time]
	BuyID       int
	StepID      int
}
