package products

import (
	"go-project-template-v7/pkg/pageable"
	"time"
)

// productsCreateRequest stores products with a reference to their category.
type productsCreateRequest struct {
	// Foreign key referencing the category to which the product belongs.
	CategoryID int `json:"category_id"`

	// Name of the product.
	Name string `json:"name"`

	// Detailed description of the product.
	Description *string `json:"description"`
}

// productsUpdateRequest stores products with a reference to their category.
type productsUpdateRequest struct {
	// Foreign key referencing the category to which the product belongs.
	CategoryID int `json:"category_id"`

	// Name of the product.
	Name string `json:"name"`

	// Detailed description of the product.
	Description *string `json:"description"`
}

// productsResponse stores products with a reference to their category.
type productsResponse struct {
	// Primary key.
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
	GUID string `json:"guid"`
}

// productsResponseList response list
type productsResponseList struct {
	// Page information (if present)
	Page *pageable.Page `json:"page,omitempty"`

	// Payload
	Data []productsResponse `json:"data"`
}
