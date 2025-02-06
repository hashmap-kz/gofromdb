package postgres

import "time"

type JobTitles struct {
	RecordID  int    `json:"record_id" db:"record_id"`
	TitleName string `json:"title_name" db:"title_name"`
	// Internal field, creation TS
	CreatedAt time.Time `json:"created_at" db:"created_at"`

	// Internal field, last updated TS
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`

	// Internal field, UUID of the row
	Guid string `json:"guid" db:"guid"`
}
