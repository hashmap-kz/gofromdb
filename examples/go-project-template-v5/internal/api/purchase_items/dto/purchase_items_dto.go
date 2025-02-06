package dto

import "time"

type PurchaseItemsDto struct {
	RecordID   int
	PurchaseID int
	ProductID  int
	Quantity   int
	Price      string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Guid       string
}

type PurchaseItemsCreateDto struct {
	PurchaseID int
	ProductID  int
	Quantity   int
	Price      string
}

type PurchaseItemsUpdateDto struct {
	PurchaseID int
	ProductID  int
	Quantity   int
	Price      string
}
