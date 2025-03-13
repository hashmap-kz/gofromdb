package dto

import "time"

type UsersDto struct {
	RecordID  int
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	GUID      string
}

type UsersCreateDto struct {
	Email string
}

type UsersUpdateDto struct {
	Email string
}
