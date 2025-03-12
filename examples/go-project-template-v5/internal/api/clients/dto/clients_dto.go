package dto

import "time"

type ClientsDto struct {
	RecordID  int
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	GUID      string
}

type ClientsCreateDto struct {
	Email string
}

type ClientsUpdateDto struct {
	Email string
}
