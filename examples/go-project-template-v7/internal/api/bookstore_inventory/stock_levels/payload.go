package stock_levels

import (
	"go-project-template-v7/pkg/pageable"
	"time"
)

// stockLevelsCreateRequest
// Current stock per warehouse and book. Natural + surrogate composite foreign
// key primary key.
type stockLevelsCreateRequest struct {
	WarehouseCode    string     `json:"warehouse_code"`
	BookID           int64      `json:"book_id"`
	AvailableQty     int        `json:"available_qty"`
	ReservedQty      int        `json:"reserved_qty"`
	ReorderThreshold int        `json:"reorder_threshold"`
	LastCountedAt    *time.Time `json:"last_counted_at"`
}

// stockLevelsUpdateRequest
// Current stock per warehouse and book. Natural + surrogate composite foreign
// key primary key.
type stockLevelsUpdateRequest struct {
	AvailableQty     *int       `json:"available_qty"`
	ReservedQty      *int       `json:"reserved_qty"`
	ReorderThreshold *int       `json:"reorder_threshold"`
	LastCountedAt    *time.Time `json:"last_counted_at"`
}

// stockLevelsResponse
// Current stock per warehouse and book. Natural + surrogate composite foreign
// key primary key.
type stockLevelsResponse struct {
	WarehouseCode    string     `json:"warehouse_code"`
	BookID           int64      `json:"book_id"`
	AvailableQty     int        `json:"available_qty"`
	ReservedQty      int        `json:"reserved_qty"`
	ReorderThreshold int        `json:"reorder_threshold"`
	LastCountedAt    *time.Time `json:"last_counted_at"`
}

// stockLevelsResponseList response list
type stockLevelsResponseList struct {
	// Page information (if present)
	Page *pageable.Page `json:"page,omitempty"`

	// Payload
	Data []stockLevelsResponse `json:"data"`
}
