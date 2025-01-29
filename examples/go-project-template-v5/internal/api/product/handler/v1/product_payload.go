package v1

import (
	"go-project-template-v5/pkg/pageable"
	"time"
)

// productCreateRequest stores products with a reference to their category.
type productCreateRequest struct {
	// Foreign key referencing the category to which the product belongs.
	CategoryID int `json:"category_id"`

	// Name of the product.
	Name string `json:"name"`

	// Detailed description of the product.
	Description *string `json:"description"`
}

// productUpdateRequest stores products with a reference to their category.
type productUpdateRequest struct {
	// Primary key for the product table.
	RecordID int `json:"record_id"`

	// Foreign key referencing the category to which the product belongs.
	CategoryID int `json:"category_id"`

	// Name of the product.
	Name string `json:"name"`

	// Detailed description of the product.
	Description *string `json:"description"`
}

// productResponse stores products with a reference to their category.
type productResponse struct {
	// Primary key for the product table.
	RecordID int `json:"record_id"`

	// Foreign key referencing the category to which the product belongs.
	CategoryID int `json:"category_id"`

	// Name of the product.
	Name string `json:"name"`

	// Detailed description of the product.
	Description *string `json:"description"`

	// Internal field, creation TS
	CreatedAt time.Time `json:"created_at"`

	// Internal field, last updated TS
	UpdatedAt time.Time `json:"updated_at"`

	// Internal field, UUID of the row
	Guid string `json:"guid"`
}

// productResponseList response list
type productResponseList struct {
	// Page information (if present)
	Page pageable.Page `json:"page,omitempty"`

	// Payload
	Data []productResponse `json:"data"`
}
