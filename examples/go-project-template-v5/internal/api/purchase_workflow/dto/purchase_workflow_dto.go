package dto

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type PurchaseWorkflowDto struct {
	RecordID       int
	ValidPeriod    pgtype.Range[time.Time]
	BuyID          int
	PurchaseStepID int
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Guid           string
}

type PurchaseWorkflowCreateDto struct {
	ValidPeriod    pgtype.Range[time.Time]
	BuyID          int
	PurchaseStepID int
}

type PurchaseWorkflowUpdateDto struct {
	ValidPeriod    pgtype.Range[time.Time]
	BuyID          int
	PurchaseStepID int
}
