package authors

import (
	"go-project-template-v7/pkg/pageable"
	"time"
)

// authorsCreateRequest
// Authors. Uses a UUID primary key with a database default.
type authorsCreateRequest struct {
	AuthorID    string     `json:"author_id"`
	DisplayName string     `json:"display_name"`
	LegalName   *string    `json:"legal_name"`
	Biography   *string    `json:"biography"`
	Metadata    string     `json:"metadata"`
	Active      bool       `json:"active"`
	BornOn      *time.Time `json:"born_on"`
}

// authorsUpdateRequest
// Authors. Uses a UUID primary key with a database default.
type authorsUpdateRequest struct {
	DisplayName *string    `json:"display_name"`
	LegalName   *string    `json:"legal_name"`
	Biography   *string    `json:"biography"`
	Metadata    *string    `json:"metadata"`
	Active      *bool      `json:"active"`
	BornOn      *time.Time `json:"born_on"`
}

// authorsResponse
// Authors. Uses a UUID primary key with a database default.
type authorsResponse struct {
	AuthorID    string     `json:"author_id"`
	DisplayName string     `json:"display_name"`
	LegalName   *string    `json:"legal_name"`
	Biography   *string    `json:"biography"`
	Metadata    string     `json:"metadata"`
	Active      bool       `json:"active"`
	BornOn      *time.Time `json:"born_on"`
}

// authorsResponseList response list
type authorsResponseList struct {
	// Page information (if present)
	Page *pageable.Page `json:"page,omitempty"`

	// Payload
	Data []authorsResponse `json:"data"`
}
