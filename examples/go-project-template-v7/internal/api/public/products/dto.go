package products

import "time"

type Dto struct {
	RecordID    int
	CategoryID  int
	Name        string
	Description *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	GUID        string
}

type CreateDto struct {
	CategoryID  int
	Name        string
	Description *string
}

type UpdateDto struct {
	CategoryID  *int
	Name        *string
	Description *string
}
