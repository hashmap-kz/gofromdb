package v1

import (
	"time"

	"go-project-template-v5/pkg/pageable"

	"github.com/jackc/pgx/v5/pgtype"
)

// productCategoriesCreateRequest represents product categories, supporting hierarchical relationships.
type productCategoriesCreateRequest struct {
	// Name of the category.
	Name string `json:"name"`

	// Reference to the parent category. NULL if it is a root category.
	ParentID *int `json:"parent_id"`

	// Validity period of this category naming.
	ValidPeriod pgtype.Range[time.Time] `json:"valid_period"`
}

// productCategoriesUpdateRequest represents product categories, supporting hierarchical relationships.
type productCategoriesUpdateRequest struct {
	// Name of the category.
	Name string `json:"name"`

	// Reference to the parent category. NULL if it is a root category.
	ParentID *int `json:"parent_id"`

	// Validity period of this category naming.
	ValidPeriod pgtype.Range[time.Time] `json:"valid_period"`
}

// productCategoriesResponse represents product categories, supporting hierarchical relationships.
type productCategoriesResponse struct {
	// Primary key for the category table.
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
	Guid string `json:"guid"`
}

// productCategoriesResponseList response list
type productCategoriesResponseList struct {
	// Page information (if present)
	Page *pageable.Page `json:"page,omitempty"`

	// Payload
	Data []productCategoriesResponse `json:"data"`
}
