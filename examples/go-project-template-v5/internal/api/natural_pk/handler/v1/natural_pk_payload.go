package v1

import (
	"go-project-template-v5/pkg/pageable"
)

type naturalPkCreateRequest struct {
	Code string `json:"code"`
	Name string `json:"name"`
}
type naturalPkUpdateRequest struct {
	Name string `json:"name"`
}
type naturalPkResponse struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// naturalPkResponseList response list
type naturalPkResponseList struct {
	// Page information (if present)
	Page *pageable.Page `json:"page,omitempty"`

	// Payload
	Data []naturalPkResponse `json:"data"`
}
