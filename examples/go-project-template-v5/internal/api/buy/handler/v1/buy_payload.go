package v1

import (
	"time"

	"go-project-template-v5/pkg/pageable"
)

type buyCreateRequest struct {
	ClientID    int     `json:"client_id"`
	Description *string `json:"description"`
}

type buyResponse struct {
	RecordID    int       `json:"record_id"`
	ClientID    int       `json:"client_id"`
	Description *string   `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Guid        string    `json:"guid"`
}

type buyResponseList struct {
	Page pageable.Page
	Data []buyResponse `json:"data"`
}

type buyUpdateRequest struct {
	RecordID    int     `json:"record_id"`
	ClientID    int     `json:"client_id"`
	Description *string `json:"description"`
}
