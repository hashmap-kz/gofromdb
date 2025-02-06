package v1

import (
	"time"

	"go-project-template-v5/pkg/pageable"
)

type jobTitlesCreateRequest struct {
	TitleName string `json:"title_name"`
}
type jobTitlesUpdateRequest struct {
	TitleName string `json:"title_name"`
}
type jobTitlesResponse struct {
	RecordID  int    `json:"record_id"`
	TitleName string `json:"title_name"`
	// Internal field, creation TS
	CreatedAt time.Time `json:"created_at"`

	// Internal field, last updated TS
	UpdatedAt time.Time `json:"updated_at"`

	// Internal field, UUID of the row
	Guid string `json:"guid"`
}

// jobTitlesResponseList response list
type jobTitlesResponseList struct {
	// Page information (if present)
	Page *pageable.Page `json:"page,omitempty"`

	// Payload
	Data []jobTitlesResponse `json:"data"`
}
