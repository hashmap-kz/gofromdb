package v1

import (
	"time"

	"go-project-template-v5/pkg/pageable"
)

// purchaseStepsCreateRequest purchase steps enum: ordered, delivered, etc...
type purchaseStepsCreateRequest struct {
	// Step name, unique
	StepName string `json:"step_name"`
}

// purchaseStepsUpdateRequest purchase steps enum: ordered, delivered, etc...
type purchaseStepsUpdateRequest struct {
	// Step name, unique
	StepName string `json:"step_name"`
}

// purchaseStepsResponse purchase steps enum: ordered, delivered, etc...
type purchaseStepsResponse struct {
	RecordID int `json:"record_id"`
	// Step name, unique
	StepName string `json:"step_name"`

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
