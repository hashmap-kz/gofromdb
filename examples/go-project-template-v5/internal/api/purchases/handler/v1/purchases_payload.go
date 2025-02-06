package v1

import (
	"time"

	"go-project-template-v5/pkg/pageable"
)

// purchasesCreateRequest represents purchases made by clients.
type purchasesCreateRequest struct {
	// Foreign key referencing the customer who made the purchase.
	CustomerID int `json:"customer_id"`

	// Optional description or additional details of the purchase.
	Description *string `json:"description"`
}

// purchasesUpdateRequest represents purchases made by clients.
type purchasesUpdateRequest struct {
	// Foreign key referencing the customer who made the purchase.
	CustomerID int `json:"customer_id"`

	// Optional description or additional details of the purchase.
	Description *string `json:"description"`
}

// purchasesResponse represents purchases made by clients.
type purchasesResponse struct {
	// Primary key for the buy table.
	RecordID int `json:"record_id"`

	// Foreign key referencing the customer who made the purchase.
	CustomerID int `json:"customer_id"`

	// Optional description or additional details of the purchase.
	Description *string `json:"description"`

	// Internal field, creation TS
	CreatedAt time.Time `json:"created_at"`

	// Internal field, last updated TS
	UpdatedAt time.Time `json:"updated_at"`

	// Internal field, UUID of the row
	Guid string `json:"guid"`
}

// purchasesResponseList response list
type purchasesResponseList struct {
	// Page information (if present)
	Page *pageable.Page `json:"page,omitempty"`

	// Payload
	Data []purchasesResponse `json:"data"`
}
