package postgres

import "time"

// Customers stores customers information, identified by a unique email.
type Customers struct {
	// Primary key for the customers table.
	RecordID int `json:"record_id" db:"record_id"`

	// Unique email address of the customer.
	Email string `json:"email" db:"email"`

	// Internal field, creation TS
	CreatedAt time.Time `json:"created_at" db:"created_at"`

	// Internal field, last updated TS
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`

	// Internal field, UUID of the row
	Guid string `json:"guid" db:"guid"`
}
