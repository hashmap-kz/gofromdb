package orders

import "time"

// Orders represents purchases made by users.
type Orders struct {
	// Primary key.
	RecordID int `json:"record_id" db:"record_id"`

	// Foreign key referencing the user who made the order.
	UserID int `json:"user_id" db:"user_id"`

	// Optional description or additional details of the order.
	Description *string `json:"description" db:"description"`

	// Internal field, creation TS
	CreatedAt time.Time `json:"created_at" db:"created_at"`

	// Internal field, last updated TS
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`

	// Internal field, UUID of the row
	GUID string `json:"guid" db:"guid"`
}
