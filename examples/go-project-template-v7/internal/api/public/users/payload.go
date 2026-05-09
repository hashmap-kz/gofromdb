package users

import (
	"go-project-template-v7/pkg/pageable"
	"time"
)

// usersCreateRequest
// Stores users information, identified by a unique email.
type usersCreateRequest struct {
	// Unique email address of the user.
	Email string `json:"email"`
}

// usersUpdateRequest
// Stores users information, identified by a unique email.
type usersUpdateRequest struct {
	// Unique email address of the user.
	Email string `json:"email"`
}

// usersResponse
// Stores users information, identified by a unique email.
type usersResponse struct {
	// Primary key for the users table.
	RecordID int `json:"record_id"`

	// Unique email address of the user.
	Email string `json:"email"`

	// Internal field, creation TS
	CreatedAt time.Time `json:"created_at"`

	// Internal field, last updated TS
	UpdatedAt time.Time `json:"updated_at"`

	// Internal field, UUID of the row
	GUID string `json:"guid"`
}

// usersResponseList response list
type usersResponseList struct {
	// Page information (if present)
	Page *pageable.Page `json:"page,omitempty"`

	// Payload
	Data []usersResponse `json:"data"`
}
