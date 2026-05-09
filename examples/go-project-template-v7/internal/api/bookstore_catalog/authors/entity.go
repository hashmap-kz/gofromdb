package authors

import "time"

// Authors
// Authors. Uses a UUID primary key with a database default.
type Authors struct {
	AuthorID    string     `json:"author_id" db:"author_id"`
	DisplayName string     `json:"display_name" db:"display_name"`
	LegalName   *string    `json:"legal_name" db:"legal_name"`
	Biography   *string    `json:"biography" db:"biography"`
	Metadata    string     `json:"metadata" db:"metadata"`
	Active      bool       `json:"active" db:"active"`
	BornOn      *time.Time `json:"born_on" db:"born_on"`
}
