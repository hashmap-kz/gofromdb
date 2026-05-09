package customers

import (
	"go-project-template-v7/pkg/pageable"
	"time"
)

// customersCreateRequest
// Store customers. UUID primary key and simple unique field.
type customersCreateRequest struct {
	CustomerID     string    `json:"customer_id"`
	Email          string    `json:"email"`
	FullName       string    `json:"full_name"`
	Phone          *string   `json:"phone"`
	MarketingOptIn bool      `json:"marketing_opt_in"`
	RegisteredAt   time.Time `json:"registered_at"`
}

// customersUpdateRequest
// Store customers. UUID primary key and simple unique field.
type customersUpdateRequest struct {
	Email          string    `json:"email"`
	FullName       string    `json:"full_name"`
	Phone          *string   `json:"phone"`
	MarketingOptIn bool      `json:"marketing_opt_in"`
	RegisteredAt   time.Time `json:"registered_at"`
}

// customersResponse
// Store customers. UUID primary key and simple unique field.
type customersResponse struct {
	CustomerID     string    `json:"customer_id"`
	Email          string    `json:"email"`
	FullName       string    `json:"full_name"`
	Phone          *string   `json:"phone"`
	MarketingOptIn bool      `json:"marketing_opt_in"`
	RegisteredAt   time.Time `json:"registered_at"`
}

// customersResponseList response list
type customersResponseList struct {
	// Page information (if present)
	Page *pageable.Page `json:"page,omitempty"`

	// Payload
	Data []customersResponse `json:"data"`
}
