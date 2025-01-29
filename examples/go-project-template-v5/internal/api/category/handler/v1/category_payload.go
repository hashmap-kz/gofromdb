package v1

import (
	"go-project-template-v5/pkg/pageable"
	"time"
)

// categoryCreateRequest represents product categories, supporting hierarchical relationships.
type categoryCreateRequest struct {
	// Name of the category.
	Name string `json:"name"`

	// Reference to the parent category. NULL if it is a root category.
	ParentID *int `json:"parent_id"`
}

// categoryUpdateRequest represents product categories, supporting hierarchical relationships.
type categoryUpdateRequest struct {
	// Primary key for the category table.
	RecordID int `json:"record_id"`

	// Name of the category.
	Name string `json:"name"`

	// Reference to the parent category. NULL if it is a root category.
	ParentID *int `json:"parent_id"`
}

// categoryResponse represents product categories, supporting hierarchical relationships.
type categoryResponse struct {
	// Primary key for the category table.
	RecordID int `json:"record_id"`

	// Name of the category.
	Name string `json:"name"`

	// Reference to the parent category. NULL if it is a root category.
	ParentID *int `json:"parent_id"`

	// Internal field, creation TS
	CreatedAt time.Time `json:"created_at"`

	// Internal field, last updated TS
	UpdatedAt time.Time `json:"updated_at"`

	// Internal field, UUID of the row
	Guid string `json:"guid"`
}

// categoryResponseList response list
type categoryResponseList struct {
	// Page information (if present)
	Page pageable.Page `json:"page,omitempty"`

	// Payload
	Data []categoryResponse `json:"data"`
}
