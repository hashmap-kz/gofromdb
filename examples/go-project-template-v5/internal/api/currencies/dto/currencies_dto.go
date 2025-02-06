package dto

import "time"

type CurrenciesDto struct {
	RecordID      int
	CurrencyCode  string
	CurrencyValue string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Guid          string
}

type CurrenciesCreateDto struct {
	CurrencyCode  string
	CurrencyValue string
}

type CurrenciesUpdateDto struct {
	CurrencyCode  string
	CurrencyValue string
}
