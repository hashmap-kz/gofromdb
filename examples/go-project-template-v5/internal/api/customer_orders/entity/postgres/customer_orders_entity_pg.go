package postgres

import "time"

// CustomerOrders represents purchases made by clients.
type CustomerOrders struct {
	// Primary key for the customer_orders table.
	RecordID int `json:"record_id" db:"record_id"`

	// Foreign key referencing the client who made the order.
	ClientID int `json:"client_id" db:"client_id"`

	// Optional description or additional details of the order.
	Description *string `json:"description" db:"description"`

	// Internal field, creation TS
	CreatedAt time.Time `json:"created_at" db:"created_at"`

	// Internal field, last updated TS
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`

	// Internal field, UUID of the row
	GUID string `json:"guid" db:"guid"`
}
