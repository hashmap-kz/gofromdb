package orders

import "time"

type Dto struct {
	RecordID    int
	UserID      int
	Description *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	GUID        string
}

type CreateDto struct {
	UserID      int
	Description *string
}

type UpdateDto struct {
	UserID      int
	Description *string
}
