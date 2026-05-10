package stock_levels

import "time"

type Dto struct {
	WarehouseCode    string
	BookID           int64
	AvailableQty     int
	ReservedQty      int
	ReorderThreshold int
	LastCountedAt    *time.Time
}

type CreateDto struct {
	WarehouseCode    string
	BookID           int64
	AvailableQty     int
	ReservedQty      int
	ReorderThreshold int
	LastCountedAt    *time.Time
}

type UpdateDto struct {
	AvailableQty     *int
	ReservedQty      *int
	ReorderThreshold *int
	LastCountedAt    *time.Time
}
