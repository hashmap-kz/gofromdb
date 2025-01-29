package postgres

import "time"

type ProductDto struct {
	RecordID    int
	CategoryID  int
	Name        string
	Description *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Guid        string
}

type ProductCreateDto struct {
	CategoryID  int
	Name        string
	Description *string
}

type ProductUpdateDto struct {
	RecordID    int
	CategoryID  int
	Name        string
	Description *string
}
