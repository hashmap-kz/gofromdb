package v1

import (
	"time"

	"go-project-template-v5/pkg/pageable"
)

type noPkCreateRequest struct {
	EventTime time.Time `json:"event_time"`
	Message   string    `json:"message"`
}
type noPkUpdateRequest struct {
	EventTime time.Time `json:"event_time"`
	Message   string    `json:"message"`
}
type noPkResponse struct {
	EventTime time.Time `json:"event_time"`
	Message   string    `json:"message"`
}

// noPkResponseList response list
type noPkResponseList struct {
	// Page information (if present)
	Page *pageable.Page `json:"page,omitempty"`

	// Payload
	Data []noPkResponse `json:"data"`
}
