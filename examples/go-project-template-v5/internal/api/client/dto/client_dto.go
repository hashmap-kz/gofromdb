package dto

import "time"

type ClientDto struct {
	RecordID  int
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	Guid      string
}

type ClientCreateDto struct {
	Email string
}

type ClientUpdateDto struct {
	RecordID int
	Email    string
}
