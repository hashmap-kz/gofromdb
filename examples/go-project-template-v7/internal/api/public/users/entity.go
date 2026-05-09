package users

import "time"

// Users
// Stores users information, identified by a unique email.
type Users struct {
	// Primary key for the users table.
	RecordID int `json:"record_id" db:"record_id"`

	// Unique email address of the user.
	Email string `json:"email" db:"email"`

	// Internal field, creation TS
	CreatedAt time.Time `json:"created_at" db:"created_at"`

	// Internal field, last updated TS
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`

	// Internal field, UUID of the row
	GUID string `json:"guid" db:"guid"`
}
