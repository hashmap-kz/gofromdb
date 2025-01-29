package v1

import "go-project-template-v5/pkg/pageable"

// clientCreateRequest - handler layer
type clientCreateRequest struct {
	Email string `json:"email" validate:"required"`
}

// clientResponse - single client payload
type clientResponse struct {
	// Identifier
	ID int `json:"id"`

	// Client email
	Email string `json:"email"`
}

// clientResponseList - paginated response
type clientResponseList struct {
	// Pagination
	Page pageable.Page `json:"page"`

	// Files list
	Data []clientResponse `json:"data,omitempty"`
}

// clientUpdateRequest - request for update available client fields
type clientUpdateRequest struct {
	// Client email
	Email string
}
