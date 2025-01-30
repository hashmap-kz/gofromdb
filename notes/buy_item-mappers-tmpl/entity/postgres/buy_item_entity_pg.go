package postgres

import "time"

// BuyItem represents items in a purchase, including quantity and price.
type BuyItem struct {
	// Primary key for the buy_item table.
	RecordID int `json:"record_id" db:"record_id"`

	// Foreign key referencing the associated purchase.
	BuyID int `json:"buy_id" db:"buy_id"`

	// Foreign key referencing the purchased product.
	ProductID int `json:"product_id" db:"product_id"`

	// Number of units of the product in the purchase.
	Quantity int `json:"quantity" db:"quantity"`

	// Price per unit of the product at the time of purchase.
	Price string `json:"price" db:"price"`

	// Internal field, creation TS
	CreatedAt time.Time `json:"created_at" db:"created_at"`

	// Internal field, last updated TS
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`

	// Internal field, UUID of the row
	Guid string `json:"guid" db:"guid"`
}
