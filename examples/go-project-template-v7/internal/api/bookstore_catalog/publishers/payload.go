package publishers

import (
	"go-project-template-v7/pkg/pageable"
	"time"
)

// publishersCreateRequest
// Book publishers. Uses a natural text primary key.
type publishersCreateRequest struct {
	// Natural publisher code, for example no_starch or manning.
	Code string `json:"code"`

	Name        string     `json:"name"`
	CountryCode string     `json:"country_code"`
	Website     *string    `json:"website"`
	FoundedOn   *time.Time `json:"founded_on"`
	Active      bool       `json:"active"`
}

// publishersUpdateRequest
// Book publishers. Uses a natural text primary key.
type publishersUpdateRequest struct {
	Name        *string    `json:"name"`
	CountryCode *string    `json:"country_code"`
	Website     *string    `json:"website"`
	FoundedOn   *time.Time `json:"founded_on"`
	Active      *bool      `json:"active"`
}

// publishersResponse
// Book publishers. Uses a natural text primary key.
type publishersResponse struct {
	// Natural publisher code, for example no_starch or manning.
	Code string `json:"code"`

	Name        string     `json:"name"`
	CountryCode string     `json:"country_code"`
	Website     *string    `json:"website"`
	FoundedOn   *time.Time `json:"founded_on"`
	Active      bool       `json:"active"`
}

// publishersResponseList response list
type publishersResponseList struct {
	// Page information (if present)
	Page *pageable.Page `json:"page,omitempty"`

	// Payload
	Data []publishersResponse `json:"data"`
}
