package dto

import "time"

type BuyItemDto struct {
	RecordID  int
	BuyID     int
	ProductID int
	Quantity  int
	Price     string
	CreatedAt time.Time
	UpdatedAt time.Time
	Guid      string
}

type BuyItemCreateDto struct {
	BuyID     int
	ProductID int
	Quantity  int
	Price     string
}

type BuyItemUpdateDto struct {
	BuyID     int
	ProductID int
	Quantity  int
	Price     string
}
