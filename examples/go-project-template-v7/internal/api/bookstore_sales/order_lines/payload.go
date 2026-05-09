package order_lines

import (
	"go-project-template-v7/pkg/pageable"
)

// orderLinesCreateRequest
// Order lines. Tests composite primary key order: order_id, line_no.
type orderLinesCreateRequest struct {
	OrderID        int64   `json:"order_id"`
	LineNo         int     `json:"line_no"`
	BookID         int64   `json:"book_id"`
	Quantity       int16   `json:"quantity"`
	UnitPrice      string  `json:"unit_price"`
	DiscountAmount string  `json:"discount_amount"`
	Note           *string `json:"note"`
}

// orderLinesUpdateRequest
// Order lines. Tests composite primary key order: order_id, line_no.
type orderLinesUpdateRequest struct {
	BookID         int64   `json:"book_id"`
	Quantity       int16   `json:"quantity"`
	UnitPrice      string  `json:"unit_price"`
	DiscountAmount string  `json:"discount_amount"`
	Note           *string `json:"note"`
}

// orderLinesResponse
// Order lines. Tests composite primary key order: order_id, line_no.
type orderLinesResponse struct {
	OrderID        int64   `json:"order_id"`
	LineNo         int     `json:"line_no"`
	BookID         int64   `json:"book_id"`
	Quantity       int16   `json:"quantity"`
	UnitPrice      string  `json:"unit_price"`
	DiscountAmount string  `json:"discount_amount"`
	Note           *string `json:"note"`
}

// orderLinesResponseList response list
type orderLinesResponseList struct {
	// Page information (if present)
	Page *pageable.Page `json:"page,omitempty"`

	// Payload
	Data []orderLinesResponse `json:"data"`
}
