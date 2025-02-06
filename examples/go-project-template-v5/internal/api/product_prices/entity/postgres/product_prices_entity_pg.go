package postgres

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

// ProductPrices prices of goods and services
type ProductPrices struct {
	// PK
	RecordID int `json:"record_id" db:"record_id"`

	// Effective date range for the price
	ValidPeriod pgtype.Range[time.Time] `json:"valid_period" db:"valid_period"`

	// References to products
	ProductID int `json:"product_id" db:"product_id"`

	// Actual price
	ProductPrice string `json:"product_price" db:"product_price"`

	// Internal field, creation TS
	CreatedAt time.Time `json:"created_at" db:"created_at"`

	// Internal field, last updated TS
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`

	// Internal field, UUID of the row
	Guid string `json:"guid" db:"guid"`
}
