package order_items

import "time"

// OrderItems represents items in a sales, including quantity and price.
type OrderItems struct {
	// Primary key.
	RecordID int `json:"record_id" db:"record_id"`

	// Foreign key referencing the associated order.
	OrderID int `json:"order_id" db:"order_id"`

	// Foreign key referencing the product.
	ProductID int `json:"product_id" db:"product_id"`

	// Number of units of the product.
	Quantity string `json:"quantity" db:"quantity"`

	// Price per unit of the product at the time of ordering.
	Price string `json:"price" db:"price"`

	// Internal field, creation TS
	CreatedAt time.Time `json:"created_at" db:"created_at"`

	// Internal field, last updated TS
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`

	// Internal field, UUID of the row
	GUID string `json:"guid" db:"guid"`
}
