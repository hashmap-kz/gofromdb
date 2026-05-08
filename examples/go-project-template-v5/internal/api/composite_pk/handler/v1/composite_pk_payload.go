package v1

import (
	"go-project-template-v5/pkg/pageable"
)

type compositePkCreateRequest struct {
	TenantID int64  `json:"tenant_id"`
	Code     string `json:"code"`
	Name     string `json:"name"`
}
type compositePkUpdateRequest struct {
	Name string `json:"name"`
}
type compositePkResponse struct {
	TenantID int64  `json:"tenant_id"`
	Code     string `json:"code"`
	Name     string `json:"name"`
}

// compositePkResponseList response list
type compositePkResponseList struct {
	// Page information (if present)
	Page *pageable.Page `json:"page,omitempty"`

	// Payload
	Data []compositePkResponse `json:"data"`
}
