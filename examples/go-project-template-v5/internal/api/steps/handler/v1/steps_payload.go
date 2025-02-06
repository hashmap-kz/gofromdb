package v1

import (
	"time"

	"go-project-template-v5/pkg/pageable"
)

// stepsCreateRequest purchase steps enum: ordered, delivered, etc...
type stepsCreateRequest struct {
	// Step name, unique
	StepName string `json:"step_name"`
}

// stepsUpdateRequest purchase steps enum: ordered, delivered, etc...
type stepsUpdateRequest struct {
	// Step name, unique
	StepName string `json:"step_name"`
}

// stepsResponse purchase steps enum: ordered, delivered, etc...
type stepsResponse struct {
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

// stepsResponseList response list
type stepsResponseList struct {
	// Page information (if present)
	Page *pageable.Page `json:"page,omitempty"`

	// Payload
	Data []stepsResponse `json:"data"`
}
