package postgres

import "time"

// PurchaseItems represents items in a purchase, including quantity and price.
type PurchaseItems struct {
	// Primary key for the buy_item table.
	RecordID int `json:"record_id" db:"record_id"`

	// Foreign key referencing the associated purchase.
	PurchaseID int `json:"purchase_id" db:"purchase_id"`

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
