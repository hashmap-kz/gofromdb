package orders

import (
	"go-project-template-v7/pkg/pageable"
	"time"
)

// bookstoreSalesOrdersCreateRequest
// Sales orders. Duplicate table name with public.orders to test schema-aware
// generation.
type bookstoreSalesOrdersCreateRequest struct {
	CustomerID  string     `json:"customer_id"`
	Status      string     `json:"status"`
	PlacedAt    *time.Time `json:"placed_at"`
	PaidAt      *time.Time `json:"paid_at"`
	CancelledAt *time.Time `json:"cancelled_at"`
	Comment     *string    `json:"comment"`
}

// bookstoreSalesOrdersUpdateRequest
// Sales orders. Duplicate table name with public.orders to test schema-aware
// generation.
type bookstoreSalesOrdersUpdateRequest struct {
	CustomerID  string     `json:"customer_id"`
	Status      string     `json:"status"`
	PlacedAt    *time.Time `json:"placed_at"`
	PaidAt      *time.Time `json:"paid_at"`
	CancelledAt *time.Time `json:"cancelled_at"`
	Comment     *string    `json:"comment"`
}

// bookstoreSalesOrdersResponse
// Sales orders. Duplicate table name with public.orders to test schema-aware
// generation.
type bookstoreSalesOrdersResponse struct {
	OrderID     int64      `json:"order_id"`
	CustomerID  string     `json:"customer_id"`
	Status      string     `json:"status"`
	PlacedAt    *time.Time `json:"placed_at"`
	PaidAt      *time.Time `json:"paid_at"`
	CancelledAt *time.Time `json:"cancelled_at"`
	Comment     *string    `json:"comment"`
}

// bookstoreSalesOrdersResponseList response list
type bookstoreSalesOrdersResponseList struct {
	// Page information (if present)
	Page *pageable.Page `json:"page,omitempty"`

	// Payload
	Data []bookstoreSalesOrdersResponse `json:"data"`
}
