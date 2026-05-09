package orders

import (
	"go-project-template-v7/pkg/pageable"
	"time"
)

// ordersCreateRequest represents purchases made by users.
type ordersCreateRequest struct {
	// Foreign key referencing the user who made the order.
	UserID int `json:"user_id"`

	// Optional description or additional details of the order.
	Description *string `json:"description"`
}

// ordersUpdateRequest represents purchases made by users.
type ordersUpdateRequest struct {
	// Foreign key referencing the user who made the order.
	UserID int `json:"user_id"`

	// Optional description or additional details of the order.
	Description *string `json:"description"`
}

// ordersResponse represents purchases made by users.
type ordersResponse struct {
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

// ordersResponseList response list
type ordersResponseList struct {
	// Page information (if present)
	Page *pageable.Page `json:"page,omitempty"`

	// Payload
	Data []ordersResponse `json:"data"`
}
