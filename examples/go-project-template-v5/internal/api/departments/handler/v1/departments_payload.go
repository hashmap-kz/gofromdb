package v1

import (
	"time"

	"go-project-template-v5/pkg/pageable"
)

type departmentsCreateRequest struct {
	DepartmentName string `json:"department_name"`
}
type departmentsUpdateRequest struct {
	DepartmentName string `json:"department_name"`
}
type departmentsResponse struct {
	RecordID       int    `json:"record_id"`
	DepartmentName string `json:"department_name"`
	// Internal field, creation TS
	CreatedAt time.Time `json:"created_at"`

	// Internal field, last updated TS
	UpdatedAt time.Time `json:"updated_at"`

	// Internal field, UUID of the row
	Guid string `json:"guid"`
}

// departmentsResponseList response list
type departmentsResponseList struct {
	// Page information (if present)
	Page *pageable.Page `json:"page,omitempty"`

	// Payload
	Data []departmentsResponse `json:"data"`
}
