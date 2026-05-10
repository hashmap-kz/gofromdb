package book_authors

import (
	"go-project-template-v7/pkg/pageable"
)

// bookAuthorsCreateRequest
// Many-to-many link. Tests composite primary key order: book_id, author_id.
type bookAuthorsCreateRequest struct {
	BookID            int64   `json:"book_id"`
	AuthorID          string  `json:"author_id"`
	ContributionOrder int16   `json:"contribution_order"`
	Role              string  `json:"role"`
	Notes             *string `json:"notes"`
}

// bookAuthorsUpdateRequest
// Many-to-many link. Tests composite primary key order: book_id, author_id.
type bookAuthorsUpdateRequest struct {
	ContributionOrder *int16  `json:"contribution_order"`
	Role              *string `json:"role"`
	Notes             *string `json:"notes"`
}

// bookAuthorsResponse
// Many-to-many link. Tests composite primary key order: book_id, author_id.
type bookAuthorsResponse struct {
	BookID            int64   `json:"book_id"`
	AuthorID          string  `json:"author_id"`
	ContributionOrder int16   `json:"contribution_order"`
	Role              string  `json:"role"`
	Notes             *string `json:"notes"`
}

// bookAuthorsResponseList response list
type bookAuthorsResponseList struct {
	// Page information (if present)
	Page *pageable.Page `json:"page,omitempty"`

	// Payload
	Data []bookAuthorsResponse `json:"data"`
}
