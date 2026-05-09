package categories

import (
	"go-project-template-v5/pkg/pageable"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

// categoriesCreateRequest represents product categories, supporting hierarchical relationships.
type categoriesCreateRequest struct {
	// Name of the category.
	Name string `json:"name"`

	// Reference to the parent category. NULL if it is a root category.
	ParentID *int `json:"parent_id"`

	// Validity period of this category naming.
	ValidPeriod pgtype.Range[time.Time] `json:"valid_period"`
}

// categoriesUpdateRequest represents product categories, supporting hierarchical relationships.
type categoriesUpdateRequest struct {
	// Name of the category.
	Name string `json:"name"`

	// Reference to the parent category. NULL if it is a root category.
	ParentID *int `json:"parent_id"`

	// Validity period of this category naming.
	ValidPeriod pgtype.Range[time.Time] `json:"valid_period"`
}

// categoriesResponse represents product categories, supporting hierarchical relationships.
type categoriesResponse struct {
	// Primary key.
	RecordID int `json:"record_id"`

	// Name of the category.
	Name string `json:"name"`

	// Reference to the parent category. NULL if it is a root category.
	ParentID *int `json:"parent_id"`

	// Validity period of this category naming.
	ValidPeriod pgtype.Range[time.Time] `json:"valid_period"`

	// Whether this category is the last actual.
	IsCurrent *bool `json:"is_current"`

	// Internal field, creation TS
	CreatedAt time.Time `json:"created_at"`

	// Internal field, last updated TS
	UpdatedAt time.Time `json:"updated_at"`

	// Internal field, UUID of the row
	GUID string `json:"guid"`
}

// categoriesResponseList response list
type categoriesResponseList struct {
	// Page information (if present)
	Page *pageable.Page `json:"page,omitempty"`

	// Payload
	Data []categoriesResponse `json:"data"`
}
