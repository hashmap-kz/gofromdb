package v1

import (
	"time"

	"go-project-template-v5/pkg/pageable"

	"github.com/jackc/pgx/v5/pgtype"
)

// purchaseWorkflowCreateRequest buy steps tracking
type purchaseWorkflowCreateRequest struct {
	// Period, open means that the step is in progress
	ValidPeriod pgtype.Range[time.Time] `json:"valid_period"`

	// Buy-order ID
	BuyID int `json:"buy_id"`

	// Step ID
	PurchaseStepID int `json:"purchase_step_id"`
}

// purchaseWorkflowUpdateRequest buy steps tracking
type purchaseWorkflowUpdateRequest struct {
	// Period, open means that the step is in progress
	ValidPeriod pgtype.Range[time.Time] `json:"valid_period"`

	// Buy-order ID
	BuyID int `json:"buy_id"`

	// Step ID
	PurchaseStepID int `json:"purchase_step_id"`
}

// purchaseWorkflowResponse buy steps tracking
type purchaseWorkflowResponse struct {
	// PK
	RecordID int `json:"record_id"`

	// Period, open means that the step is in progress
	ValidPeriod pgtype.Range[time.Time] `json:"valid_period"`

	// Buy-order ID
	BuyID int `json:"buy_id"`

	// Step ID
	PurchaseStepID int `json:"purchase_step_id"`

	// Internal field, creation TS
	CreatedAt time.Time `json:"created_at"`

	// Internal field, last updated TS
	UpdatedAt time.Time `json:"updated_at"`

	// Internal field, UUID of the row
	Guid string `json:"guid"`
}

// purchaseWorkflowResponseList response list
type purchaseWorkflowResponseList struct {
	// Page information (if present)
	Page *pageable.Page `json:"page,omitempty"`

	// Payload
	Data []purchaseWorkflowResponse `json:"data"`
}
