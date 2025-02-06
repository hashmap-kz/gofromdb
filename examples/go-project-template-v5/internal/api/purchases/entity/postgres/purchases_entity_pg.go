package postgres

import "time"

// Purchases represents purchases made by clients.
type Purchases struct {
	// Primary key for the buy table.
	RecordID int `json:"record_id" db:"record_id"`

	// Foreign key referencing the customer who made the purchase.
	CustomerID int `json:"customer_id" db:"customer_id"`

	// Optional description or additional details of the purchase.
	Description *string `json:"description" db:"description"`

	// Internal field, creation TS
	CreatedAt time.Time `json:"created_at" db:"created_at"`

	// Internal field, last updated TS
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`

	// Internal field, UUID of the row
	Guid string `json:"guid" db:"guid"`
}
