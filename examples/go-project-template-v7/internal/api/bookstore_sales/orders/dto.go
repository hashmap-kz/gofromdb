package orders

import "time"

type Dto struct {
	OrderID     int64
	CustomerID  string
	Status      string
	PlacedAt    *time.Time
	PaidAt      *time.Time
	CancelledAt *time.Time
	Comment     *string
}

type CreateDto struct {
	CustomerID  string
	Status      string
	PlacedAt    *time.Time
	PaidAt      *time.Time
	CancelledAt *time.Time
	Comment     *string
}

type UpdateDto struct {
	CustomerID  string
	Status      string
	PlacedAt    *time.Time
	PaidAt      *time.Time
	CancelledAt *time.Time
	Comment     *string
}
