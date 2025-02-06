package postgres

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

// PurchaseSteps buy steps tracking
type PurchaseSteps struct {
	// PK
	RecordID int `json:"record_id" db:"record_id"`

	// Period, open means that the step is in progress
	ValidPeriod pgtype.Range[time.Time] `json:"valid_period" db:"valid_period"`

	// Buy-order ID
	BuyID int `json:"buy_id" db:"buy_id"`

	// Step ID
	StepID int `json:"step_id" db:"step_id"`

	// Internal field, creation TS
	CreatedAt time.Time `json:"created_at" db:"created_at"`

	// Internal field, last updated TS
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`

	// Internal field, UUID of the row
	Guid string `json:"guid" db:"guid"`
}
