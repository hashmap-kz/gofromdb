package v1

import (
	"go-project-template-v5/pkg/pageable"
)

type serialPkCreateRequest struct {
	Name string `json:"name"`
}
type serialPkUpdateRequest struct {
	Name string `json:"name"`
}
type serialPkResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// serialPkResponseList response list
type serialPkResponseList struct {
	// Page information (if present)
	Page *pageable.Page `json:"page,omitempty"`

	// Payload
	Data []serialPkResponse `json:"data"`
}
