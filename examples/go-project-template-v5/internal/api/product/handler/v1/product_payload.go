package v1

import (
	"time"

	"go-project-template-v5/pkg/pageable"
)

type productCreateRequest struct {
	CategoryID  int     `json:"category_id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type productResponse struct {
	RecordID    int       `json:"record_id"`
	CategoryID  int       `json:"category_id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Guid        string    `json:"guid"`
}

type productResponseList struct {
	Page pageable.Page
	Data []productResponse `json:"data"`
}

type productUpdateRequest struct {
	RecordID    int     `json:"record_id"`
	CategoryID  int     `json:"category_id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}
