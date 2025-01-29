package v1

import (
	"time"

	"go-project-template-v5/pkg/pageable"
)

type clientCreateRequest struct {
	Email string `json:"email"`
}

type clientResponse struct {
	RecordID  int       `json:"record_id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Guid      string    `json:"guid"`
}

type clientResponseList struct {
	Page pageable.Page
	Data []clientResponse `json:"data"`
}

type clientUpdateRequest struct {
	RecordID int    `json:"record_id"`
	Email    string `json:"email"`
}
