package stock_events

import "time"

// StockEvents
// Append-only stock event stream. Intentionally has no primary key.
type StockEvents struct {
	HappenedAt    time.Time `json:"happened_at" db:"happened_at"`
	WarehouseCode string    `json:"warehouse_code" db:"warehouse_code"`
	BookID        int64     `json:"book_id" db:"book_id"`
	DeltaQty      int       `json:"delta_qty" db:"delta_qty"`
	Reason        string    `json:"reason" db:"reason"`
	Payload       string    `json:"payload" db:"payload"`
}
