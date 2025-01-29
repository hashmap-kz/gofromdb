package postgres

import "time"

// Buy represents purchases made by clients.
type Buy struct {
	// Primary key for the buy table.
	RecordID int `json:"record_id" db:"record_id"`

	// Foreign key referencing the client who made the purchase.
	ClientID int `json:"client_id" db:"client_id"`

	// Optional description or additional details of the purchase.
	Description *string `json:"description" db:"description"`

	// Internal field, creation TS
	CreatedAt time.Time `json:"created_at" db:"created_at"`

	// Internal field, last updated TS
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`

	// Internal field, UUID of the row
	Guid string `json:"guid" db:"guid"`
}
