package v1

import (
	"time"

	"go-project-template-v5/pkg/pageable"
)

type categoryCreateRequest struct {
	Name     string `json:"name"`
	ParentID *int   `json:"parent_id"`
}

type categoryResponse struct {
	RecordID  int       `json:"record_id"`
	Name      string    `json:"name"`
	ParentID  *int      `json:"parent_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Guid      string    `json:"guid"`
}

type categoryResponseList struct {
	Page pageable.Page
	Data []categoryResponse `json:"data"`
}

type categoryUpdateRequest struct {
	RecordID int    `json:"record_id"`
	Name     string `json:"name"`
	ParentID *int   `json:"parent_id"`
}
