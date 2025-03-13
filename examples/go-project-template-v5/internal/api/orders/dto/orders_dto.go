package dto

import "time"

type OrdersDto struct {
	RecordID    int
	UserID      int
	Description *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	GUID        string
}

type OrdersCreateDto struct {
	UserID      int
	Description *string
}

type OrdersUpdateDto struct {
	UserID      int
	Description *string
}
