package v1

import (
	"time"

	"go-project-template-v5/pkg/pageable"
)

type nullableTypesCreateRequest struct {
	Name    *string  `json:"name"`
	Amount  *string  `json:"amount"`
	Payload *string  `json:"payload"`
	Tags    []string `json:"tags"`
	Active  *bool    `json:"active"`
}
type nullableTypesUpdateRequest struct {
	Name    *string  `json:"name"`
	Amount  *string  `json:"amount"`
	Payload *string  `json:"payload"`
	Tags    []string `json:"tags"`
	Active  *bool    `json:"active"`
}
type nullableTypesResponse struct {
	ID        int64      `json:"id"`
	Name      *string    `json:"name"`
	Amount    *string    `json:"amount"`
	Payload   *string    `json:"payload"`
	Tags      []string   `json:"tags"`
	Active    *bool      `json:"active"`
	CreatedAt *time.Time `json:"created_at"`
}

// nullableTypesResponseList response list
type nullableTypesResponseList struct {
	// Page information (if present)
	Page *pageable.Page `json:"page,omitempty"`

	// Payload
	Data []nullableTypesResponse `json:"data"`
}
