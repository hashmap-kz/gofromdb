package v1

import (
	"time"

	"go-project-template-v5/pkg/pageable"
)

// clientsCreateRequest stores users information, identified by a unique email.
type clientsCreateRequest struct {
	// Unique email address of the user.
	Email string `json:"email"`
}

// clientsUpdateRequest stores users information, identified by a unique email.
type clientsUpdateRequest struct {
	// Unique email address of the user.
	Email string `json:"email"`
}

// clientsResponse stores users information, identified by a unique email.
type clientsResponse struct {
	// Primary key for the users table.
	RecordID int `json:"record_id"`

	// Unique email address of the user.
	Email string `json:"email"`

	// Internal field, creation TS
	CreatedAt time.Time `json:"created_at"`

	// Internal field, last updated TS
	UpdatedAt time.Time `json:"updated_at"`

	// Internal field, UUID of the row
	Guid string `json:"guid"`
}

// clientsResponseList response list
type clientsResponseList struct {
	// Page information (if present)
	Page *pageable.Page `json:"page,omitempty"`

	// Payload
	Data []clientsResponse `json:"data"`
}
