package v1

import (
	"time"

	"go-project-template-v5/pkg/pageable"
)

// customersCreateRequest stores customers information, identified by a unique email.
type customersCreateRequest struct {
	// Unique email address of the customer.
	Email string `json:"email"`
}

// customersUpdateRequest stores customers information, identified by a unique email.
type customersUpdateRequest struct {
	// Unique email address of the customer.
	Email string `json:"email"`
}

// customersResponse stores customers information, identified by a unique email.
type customersResponse struct {
	// Primary key for the customers table.
	RecordID int `json:"record_id"`

	// Unique email address of the customer.
	Email string `json:"email"`

	// Internal field, creation TS
	CreatedAt time.Time `json:"created_at"`

	// Internal field, last updated TS
	UpdatedAt time.Time `json:"updated_at"`

	// Internal field, UUID of the row
	Guid string `json:"guid"`
}

// customersResponseList response list
type customersResponseList struct {
	// Page information (if present)
	Page *pageable.Page `json:"page,omitempty"`

	// Payload
	Data []customersResponse `json:"data"`
}
