package discount_codes

import (
	"go-project-template-v7/pkg/pageable"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

// discountCodesCreateRequest
// Discount codes. Natural text primary key plus daterange.
type discountCodesCreateRequest struct {
	Code        string                  `json:"code"`
	Description *string                 `json:"description"`
	PercentOff  string                  `json:"percent_off"`
	ValidPeriod pgtype.Range[time.Time] `json:"valid_period"`
	MaxUses     *int                    `json:"max_uses"`
	Active      bool                    `json:"active"`
}

// discountCodesUpdateRequest
// Discount codes. Natural text primary key plus daterange.
type discountCodesUpdateRequest struct {
	Description *string                  `json:"description"`
	PercentOff  *string                  `json:"percent_off"`
	ValidPeriod *pgtype.Range[time.Time] `json:"valid_period"`
	MaxUses     *int                     `json:"max_uses"`
	Active      *bool                    `json:"active"`
}

// discountCodesResponse
// Discount codes. Natural text primary key plus daterange.
type discountCodesResponse struct {
	Code        string                  `json:"code"`
	Description *string                 `json:"description"`
	PercentOff  string                  `json:"percent_off"`
	ValidPeriod pgtype.Range[time.Time] `json:"valid_period"`
	MaxUses     *int                    `json:"max_uses"`
	Active      bool                    `json:"active"`
}

// discountCodesResponseList response list
type discountCodesResponseList struct {
	// Page information (if present)
	Page *pageable.Page `json:"page,omitempty"`

	// Payload
	Data []discountCodesResponse `json:"data"`
}
