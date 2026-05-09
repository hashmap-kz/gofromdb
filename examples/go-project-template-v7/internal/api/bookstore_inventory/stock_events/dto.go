package stock_events

import "time"

type Dto struct {
	HappenedAt    time.Time
	WarehouseCode string
	BookID        int64
	DeltaQty      int
	Reason        string
	Payload       string
}

type CreateDto struct {
	HappenedAt    time.Time
	WarehouseCode string
	BookID        int64
	DeltaQty      int
	Reason        string
	Payload       string
}

type UpdateDto struct {
	HappenedAt    time.Time
	WarehouseCode string
	BookID        int64
	DeltaQty      int
	Reason        string
	Payload       string
}
