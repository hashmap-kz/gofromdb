package dto

import "time"

type CustomerOrderItemsDto struct {
	RecordID        int
	CustomerOrderID int
	ProductID       int
	Quantity        int
	Price           string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Guid            string
}

type CustomerOrderItemsCreateDto struct {
	CustomerOrderID int
	ProductID       int
	Quantity        int
	Price           string
}

type CustomerOrderItemsUpdateDto struct {
	CustomerOrderID int
	ProductID       int
	Quantity        int
	Price           string
}
