package stock_levels

import "time"

// StockLevels
// Current stock per warehouse and book. Natural + surrogate composite foreign
// key primary key.
type StockLevels struct {
	WarehouseCode    string     `json:"warehouse_code" db:"warehouse_code"`
	BookID           int64      `json:"book_id" db:"book_id"`
	AvailableQty     int        `json:"available_qty" db:"available_qty"`
	ReservedQty      int        `json:"reserved_qty" db:"reserved_qty"`
	ReorderThreshold int        `json:"reorder_threshold" db:"reorder_threshold"`
	LastCountedAt    *time.Time `json:"last_counted_at" db:"last_counted_at"`
}
