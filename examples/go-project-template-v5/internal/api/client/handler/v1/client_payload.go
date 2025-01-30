package v1

import (
	"time"

	"go-project-template-v5/pkg/pageable"
)

// clientCreateRequest stores client information, identified by a unique email.
type clientCreateRequest struct {
	// Unique email address of the client.
	Email string `json:"email"`
}

// clientUpdateRequest stores client information, identified by a unique email.
type clientUpdateRequest struct {
	// Unique email address of the client.
	Email string `json:"email"`
}

// clientResponse stores client information, identified by a unique email.
type clientResponse struct {
	// Primary key for the client table.
	RecordID int `json:"record_id"`

	// Unique email address of the client.
	Email string `json:"email"`

	// Internal field, creation TS
	CreatedAt time.Time `json:"created_at"`

	// Internal field, last updated TS
	UpdatedAt time.Time `json:"updated_at"`

	// Internal field, UUID of the row
	Guid string `json:"guid"`
}

// clientResponseList response list
type clientResponseList struct {
	// Page information (if present)
	Page *pageable.Page `json:"page,omitempty"`

	// Payload
	Data []clientResponse `json:"data"`
}
