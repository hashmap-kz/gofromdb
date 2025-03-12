package dto

import "time"

type ProductsDto struct {
	RecordID    int
	CategoryID  int
	Name        string
	Description *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	GUID        string
}

type ProductsCreateDto struct {
	CategoryID  int
	Name        string
	Description *string
}

type ProductsUpdateDto struct {
	CategoryID  int
	Name        string
	Description *string
}
