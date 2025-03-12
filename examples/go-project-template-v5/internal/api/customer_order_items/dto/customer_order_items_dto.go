package dto

import "time"

type CustomerOrderItemsDto struct {
	RecordID        int
	CustomerOrderID int
	ProductID       int
	Quantity        string
	Price           string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	GUID            string
}

type CustomerOrderItemsCreateDto struct {
	CustomerOrderID int
	ProductID       int
	Quantity        string
	Price           string
}

type CustomerOrderItemsUpdateDto struct {
	CustomerOrderID int
	ProductID       int
	Quantity        string
	Price           string
}
