package postgres

import "time"

// Products stores products with a reference to their category.
type Products struct {
	// Primary key for the product table.
	RecordID int `json:"record_id" db:"record_id"`

	// Foreign key referencing the category to which the product belongs.
	CategoryID int `json:"category_id" db:"category_id"`

	// Name of the product.
	Name string `json:"name" db:"name"`

	// Detailed description of the product.
	Description *string `json:"description" db:"description"`

	// Internal field, creation TS
	CreatedAt time.Time `json:"created_at" db:"created_at"`

	// Internal field, last updated TS
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`

	// Internal field, UUID of the row
	GUID string `json:"guid" db:"guid"`
}
