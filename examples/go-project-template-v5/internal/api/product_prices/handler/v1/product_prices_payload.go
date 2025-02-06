package v1

import (
	"time"

	"go-project-template-v5/pkg/pageable"

	"github.com/jackc/pgx/v5/pgtype"
)

// productPricesCreateRequest prices of goods and services
type productPricesCreateRequest struct {
	// Effective date range for the price
	ProductPricePeriod pgtype.Range[time.Time] `json:"product_price_period"`

	// References to products
	ProductID int `json:"product_id"`

	// Actual price
	ProductPrice string `json:"product_price"`
}

// productPricesUpdateRequest prices of goods and services
type productPricesUpdateRequest struct {
	// Effective date range for the price
	ProductPricePeriod pgtype.Range[time.Time] `json:"product_price_period"`

	// References to products
	ProductID int `json:"product_id"`

	// Actual price
	ProductPrice string `json:"product_price"`
}

// productPricesResponse prices of goods and services
type productPricesResponse struct {
	// PK
	RecordID int `json:"record_id"`

	// Effective date range for the price
	ProductPricePeriod pgtype.Range[time.Time] `json:"product_price_period"`

	// References to products
	ProductID int `json:"product_id"`

	// Actual price
	ProductPrice string `json:"product_price"`

	// Internal field, creation TS
	CreatedAt time.Time `json:"created_at"`

	// Internal field, last updated TS
	UpdatedAt time.Time `json:"updated_at"`

	// Internal field, UUID of the row
	Guid string `json:"guid"`
}

// productPricesResponseList response list
type productPricesResponseList struct {
	// Page information (if present)
	Page *pageable.Page `json:"page,omitempty"`

	// Payload
	Data []productPricesResponse `json:"data"`
}
