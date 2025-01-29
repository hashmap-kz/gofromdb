package dto

import "time"

type BuyDto struct {
	RecordID    int
	ClientID    int
	Description *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Guid        string
}

type BuyCreateDto struct {
	ClientID    int
	Description *string
}

type BuyUpdateDto struct {
	RecordID    int
	ClientID    int
	Description *string
}
