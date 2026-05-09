package orders

import "time"

// BookstoreSalesOrders
// Sales orders. Duplicate table name with public.orders to test schema-aware
// generation.
type BookstoreSalesOrders struct {
	OrderID     int64      `json:"order_id" db:"order_id"`
	CustomerID  string     `json:"customer_id" db:"customer_id"`
	Status      string     `json:"status" db:"status"`
	PlacedAt    *time.Time `json:"placed_at" db:"placed_at"`
	PaidAt      *time.Time `json:"paid_at" db:"paid_at"`
	CancelledAt *time.Time `json:"cancelled_at" db:"cancelled_at"`
	Comment     *string    `json:"comment" db:"comment"`
}
