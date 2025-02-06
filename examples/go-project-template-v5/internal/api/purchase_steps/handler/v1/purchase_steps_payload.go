package v1

import (
	"time"

	"go-project-template-v5/pkg/pageable"

	"github.com/jackc/pgx/v5/pgtype"
)

// purchaseStepsCreateRequest buy steps tracking
type purchaseStepsCreateRequest struct {
	// Period, open means that the step is in progress
	ValidPeriod pgtype.Range[time.Time] `json:"valid_period"`

	// Buy-order ID
	BuyID int `json:"buy_id"`

	// Step ID
	StepID int `json:"step_id"`
}

// purchaseStepsUpdateRequest buy steps tracking
type purchaseStepsUpdateRequest struct {
	// Period, open means that the step is in progress
	ValidPeriod pgtype.Range[time.Time] `json:"valid_period"`

	// Buy-order ID
	BuyID int `json:"buy_id"`

	// Step ID
	StepID int `json:"step_id"`
}

// purchaseStepsResponse buy steps tracking
type purchaseStepsResponse struct {
	// PK
	RecordID int `json:"record_id"`

	// Period, open means that the step is in progress
	ValidPeriod pgtype.Range[time.Time] `json:"valid_period"`

	// Buy-order ID
	BuyID int `json:"buy_id"`

	// Step ID
	StepID int `json:"step_id"`

	// Internal field, creation TS
	CreatedAt time.Time `json:"created_at"`

	// Internal field, last updated TS
	UpdatedAt time.Time `json:"updated_at"`

	// Internal field, UUID of the row
	Guid string `json:"guid"`
}

// purchaseStepsResponseList response list
type purchaseStepsResponseList struct {
	// Page information (if present)
	Page *pageable.Page `json:"page,omitempty"`

	// Payload
	Data []purchaseStepsResponse `json:"data"`
}
