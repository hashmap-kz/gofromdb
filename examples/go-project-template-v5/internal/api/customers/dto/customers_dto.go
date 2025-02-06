package dto

import "time"

type CustomersDto struct {
	RecordID  int
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	Guid      string
}

type CustomersCreateDto struct {
	Email string
}

type CustomersUpdateDto struct {
	Email string
}
