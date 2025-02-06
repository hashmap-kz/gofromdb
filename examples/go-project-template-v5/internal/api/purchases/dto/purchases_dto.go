package dto

import "time"

type PurchasesDto struct {
	RecordID    int
	CustomerID  int
	Description *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Guid        string
}

type PurchasesCreateDto struct {
	CustomerID  int
	Description *string
}

type PurchasesUpdateDto struct {
	CustomerID  int
	Description *string
}
