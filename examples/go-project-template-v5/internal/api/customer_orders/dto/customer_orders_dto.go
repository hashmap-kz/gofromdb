package dto

import "time"

type CustomerOrdersDto struct {
	RecordID    int
	ClientID    int
	Description *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Guid        string
}

type CustomerOrdersCreateDto struct {
	ClientID    int
	Description *string
}

type CustomerOrdersUpdateDto struct {
	ClientID    int
	Description *string
}
