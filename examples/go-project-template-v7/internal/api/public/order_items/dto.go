package order_items

import "time"

type Dto struct {
	RecordID  int
	OrderID   int
	ProductID int
	Quantity  string
	Price     string
	CreatedAt time.Time
	UpdatedAt time.Time
	GUID      string
}

type CreateDto struct {
	OrderID   int
	ProductID int
	Quantity  string
	Price     string
}

type UpdateDto struct {
	OrderID   *int
	ProductID *int
	Quantity  *string
	Price     *string
}
