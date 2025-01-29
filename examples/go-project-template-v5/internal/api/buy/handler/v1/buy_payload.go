package v1

import (
	"time"

	"go-project-template-v5/pkg/pageable"
)

// buyCreateRequest represents purchases made by clients.
type buyCreateRequest struct {
	// Foreign key referencing the client who made the purchase.
	ClientID int `json:"client_id"`

	// Optional description or additional details of the purchase.
	Description *string `json:"description"`
}

// buyUpdateRequest represents purchases made by clients.
type buyUpdateRequest struct {
	// Primary key for the buy table.
	RecordID int `json:"record_id"`

	// Foreign key referencing the client who made the purchase.
	ClientID int `json:"client_id"`

	// Optional description or additional details of the purchase.
	Description *string `json:"description"`
}

// buyResponse represents purchases made by clients.
type buyResponse struct {
	// Primary key for the buy table.
	RecordID int `json:"record_id"`

	// Foreign key referencing the client who made the purchase.
	ClientID int `json:"client_id"`

	// Optional description or additional details of the purchase.
	Description *string `json:"description"`

	// Internal field, creation TS
	CreatedAt time.Time `json:"created_at"`

	// Internal field, last updated TS
	UpdatedAt time.Time `json:"updated_at"`

	// Internal field, UUID of the row
	Guid string `json:"guid"`
}

// buyResponseList response list
type buyResponseList struct {
	// Page information (if present)
	Page *pageable.Page `json:"page,omitempty"`

	// Payload
	Data []buyResponse `json:"data"`
}
