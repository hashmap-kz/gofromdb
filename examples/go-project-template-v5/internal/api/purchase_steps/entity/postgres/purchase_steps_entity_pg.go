package postgres

import "time"

// PurchaseSteps purchase steps enum: ordered, delivered, etc...
type PurchaseSteps struct {
	RecordID int `json:"record_id" db:"record_id"`
	// Step name, unique
	StepName string `json:"step_name" db:"step_name"`

	// Internal field, creation TS
	CreatedAt time.Time `json:"created_at" db:"created_at"`

	// Internal field, last updated TS
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`

	// Internal field, UUID of the row
	Guid string `json:"guid" db:"guid"`
}
