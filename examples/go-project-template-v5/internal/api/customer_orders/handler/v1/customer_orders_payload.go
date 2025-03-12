package v1

import (
	"time"

	"go-project-template-v5/pkg/pageable"
)

// customerOrdersCreateRequest represents purchases made by clients.
type customerOrdersCreateRequest struct {
	// Foreign key referencing the client who made the order.
	ClientID int `json:"client_id"`

	// Optional description or additional details of the order.
	Description *string `json:"description"`
}

// customerOrdersUpdateRequest represents purchases made by clients.
type customerOrdersUpdateRequest struct {
	// Foreign key referencing the client who made the order.
	ClientID int `json:"client_id"`

	// Optional description or additional details of the order.
	Description *string `json:"description"`
}

// customerOrdersResponse represents purchases made by clients.
type customerOrdersResponse struct {
	// Primary key for the customer_orders table.
	RecordID int `json:"record_id"`

	// Foreign key referencing the client who made the order.
	ClientID int `json:"client_id"`

	// Optional description or additional details of the order.
	Description *string `json:"description"`

	// Internal field, creation TS
	CreatedAt time.Time `json:"created_at"`

	// Internal field, last updated TS
	UpdatedAt time.Time `json:"updated_at"`

	// Internal field, UUID of the row
	GUID string `json:"guid"`
}

// customerOrdersResponseList response list
type customerOrdersResponseList struct {
	// Page information (if present)
	Page *pageable.Page `json:"page,omitempty"`

	// Payload
	Data []customerOrdersResponse `json:"data"`
}
