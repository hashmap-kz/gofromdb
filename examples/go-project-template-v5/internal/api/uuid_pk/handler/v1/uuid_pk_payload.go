package v1

import (
	"go-project-template-v5/pkg/pageable"
)

type uUIDPkCreateRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type uUIDPkUpdateRequest struct {
	Name string `json:"name"`
}
type uUIDPkResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// uUIDPkResponseList response list
type uUIDPkResponseList struct {
	// Page information (if present)
	Page *pageable.Page `json:"page,omitempty"`

	// Payload
	Data []uUIDPkResponse `json:"data"`
}
