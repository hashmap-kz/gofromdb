package postgres

import "time"

type Currencies struct {
	RecordID      int    `json:"record_id" db:"record_id"`
	CurrencyCode  string `json:"currency_code" db:"currency_code"`
	CurrencyValue string `json:"currency_value" db:"currency_value"`
	// Internal field, creation TS
	CreatedAt time.Time `json:"created_at" db:"created_at"`

	// Internal field, last updated TS
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`

	// Internal field, UUID of the row
	Guid string `json:"guid" db:"guid"`
}
