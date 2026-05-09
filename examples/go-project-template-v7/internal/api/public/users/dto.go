package users

import "time"

type Dto struct {
	RecordID  int
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	GUID      string
}

type CreateDto struct {
	Email string
}

type UpdateDto struct {
	Email string
}
