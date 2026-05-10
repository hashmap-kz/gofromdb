package orders

import (
	"go-project-template-v7/pkg/pageable"
	"time"
)

// publicOrdersCreateRequest
// Represents purchases made by users.
type publicOrdersCreateRequest struct {
	// Foreign key referencing the user who made the order.
	UserID int `json:"user_id"`

	// Optional description or additional details of the order.
	Description *string `json:"description"`
}

// publicOrdersUpdateRequest
// Represents purchases made by users.
type publicOrdersUpdateRequest struct {
	// Foreign key referencing the user who made the order.
	UserID *int `json:"user_id"`

	// Optional description or additional details of the order.
	Description *string `json:"description"`
}

// publicOrdersResponse
// Represents purchases made by users.
type publicOrdersResponse struct {
	// Primary key.
	RecordID int `json:"record_id"`

	// Foreign key referencing the user who made the order.
	UserID int `json:"user_id"`

	// Optional description or additional details of the order.
	Description *string `json:"description"`

	// Internal field, creation TS
	CreatedAt time.Time `json:"created_at"`

	// Internal field, last updated TS
	UpdatedAt time.Time `json:"updated_at"`

	// Internal field, UUID of the row
	GUID string `json:"guid"`
}

// publicOrdersResponseList response list
type publicOrdersResponseList struct {
	// Page information (if present)
	Page *pageable.Page `json:"page,omitempty"`

	// Payload
	Data []publicOrdersResponse `json:"data"`
}
