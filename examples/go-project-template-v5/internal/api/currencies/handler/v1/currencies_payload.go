package v1

import (
	"time"

	"go-project-template-v5/pkg/pageable"
)

type currenciesCreateRequest struct {
	CurrencyCode  string `json:"currency_code"`
	CurrencyValue string `json:"currency_value"`
}
type currenciesUpdateRequest struct {
	CurrencyCode  string `json:"currency_code"`
	CurrencyValue string `json:"currency_value"`
}
type currenciesResponse struct {
	RecordID      int    `json:"record_id"`
	CurrencyCode  string `json:"currency_code"`
	CurrencyValue string `json:"currency_value"`
	// Internal field, creation TS
	CreatedAt time.Time `json:"created_at"`

	// Internal field, last updated TS
	UpdatedAt time.Time `json:"updated_at"`

	// Internal field, UUID of the row
	Guid string `json:"guid"`
}

// currenciesResponseList response list
type currenciesResponseList struct {
	// Page information (if present)
	Page *pageable.Page `json:"page,omitempty"`

	// Payload
	Data []currenciesResponse `json:"data"`
}
