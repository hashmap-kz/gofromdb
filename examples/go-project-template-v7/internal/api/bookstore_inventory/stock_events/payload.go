package stock_events

import (
	"go-project-template-v7/pkg/pageable"
	"time"
)

// stockEventsCreateRequest
// Append-only stock event stream. Intentionally has no primary key.
type stockEventsCreateRequest struct {
	HappenedAt    time.Time `json:"happened_at"`
	WarehouseCode string    `json:"warehouse_code"`
	BookID        int64     `json:"book_id"`
	DeltaQty      int       `json:"delta_qty"`
	Reason        string    `json:"reason"`
	Payload       string    `json:"payload"`
}

// stockEventsUpdateRequest
// Append-only stock event stream. Intentionally has no primary key.
type stockEventsUpdateRequest struct {
	HappenedAt    time.Time `json:"happened_at"`
	WarehouseCode string    `json:"warehouse_code"`
	BookID        int64     `json:"book_id"`
	DeltaQty      int       `json:"delta_qty"`
	Reason        string    `json:"reason"`
	Payload       string    `json:"payload"`
}

// stockEventsResponse
// Append-only stock event stream. Intentionally has no primary key.
type stockEventsResponse struct {
	HappenedAt    time.Time `json:"happened_at"`
	WarehouseCode string    `json:"warehouse_code"`
	BookID        int64     `json:"book_id"`
	DeltaQty      int       `json:"delta_qty"`
	Reason        string    `json:"reason"`
	Payload       string    `json:"payload"`
}

// stockEventsResponseList response list
type stockEventsResponseList struct {
	// Page information (if present)
	Page *pageable.Page `json:"page,omitempty"`

	// Payload
	Data []stockEventsResponse `json:"data"`
}
