package postgres

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

// ProductCategories represents product categories, supporting hierarchical relationships.
type ProductCategories struct {
	// Primary key for the category table.
	RecordID int `json:"record_id" db:"record_id"`

	// Name of the category.
	Name string `json:"name" db:"name"`

	// Reference to the parent category. NULL if it is a root category.
	ParentID *int `json:"parent_id" db:"parent_id"`

	// Validity period of this category naming.
	ValidPeriod pgtype.Range[time.Time] `json:"valid_period" db:"valid_period"`

	// Whether this category is the last actual.
	IsCurrent *bool `json:"is_current" db:"is_current"`

	// Internal field, creation TS
	CreatedAt time.Time `json:"created_at" db:"created_at"`

	// Internal field, last updated TS
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`

	// Internal field, UUID of the row
	Guid string `json:"guid" db:"guid"`
}
