package dto

import "time"

type OrderItemsDto struct {
	RecordID  int
	OrderID   int
	ProductID int
	Quantity  string
	Price     string
	CreatedAt time.Time
	UpdatedAt time.Time
	GUID      string
}

type OrderItemsCreateDto struct {
	OrderID   int
	ProductID int
	Quantity  string
	Price     string
}

type OrderItemsUpdateDto struct {
	OrderID   int
	ProductID int
	Quantity  string
	Price     string
}
