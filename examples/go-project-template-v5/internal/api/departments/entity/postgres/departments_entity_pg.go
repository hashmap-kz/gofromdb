package postgres

import "time"

type Departments struct {
	RecordID       int    `json:"record_id" db:"record_id"`
	DepartmentName string `json:"department_name" db:"department_name"`
	// Internal field, creation TS
	CreatedAt time.Time `json:"created_at" db:"created_at"`

	// Internal field, last updated TS
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`

	// Internal field, UUID of the row
	Guid string `json:"guid" db:"guid"`
}
