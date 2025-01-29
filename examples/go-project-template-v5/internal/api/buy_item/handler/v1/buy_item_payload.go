package v1

import (
	"go-project-template-v5/pkg/pageable"
	"time"
)

// buyItemCreateRequest represents items in a purchase, including quantity and price.
type buyItemCreateRequest struct {
	// Foreign key referencing the associated purchase.
	BuyID int `json:"buy_id"`

	// Foreign key referencing the purchased product.
	ProductID int `json:"product_id"`

	// Number of units of the product in the purchase.
	Quantity int `json:"quantity"`

	// Price per unit of the product at the time of purchase.
	Price string `json:"price"`
}

// buyItemUpdateRequest represents items in a purchase, including quantity and price.
type buyItemUpdateRequest struct {
	// Primary key for the buy_item table.
	RecordID int `json:"record_id"`

	// Foreign key referencing the associated purchase.
	BuyID int `json:"buy_id"`

	// Foreign key referencing the purchased product.
	ProductID int `json:"product_id"`

	// Number of units of the product in the purchase.
	Quantity int `json:"quantity"`

	// Price per unit of the product at the time of purchase.
	Price string `json:"price"`
}

// buyItemResponse represents items in a purchase, including quantity and price.
type buyItemResponse struct {
	// Primary key for the buy_item table.
	RecordID int `json:"record_id"`

	// Foreign key referencing the associated purchase.
	BuyID int `json:"buy_id"`

	// Foreign key referencing the purchased product.
	ProductID int `json:"product_id"`

	// Number of units of the product in the purchase.
	Quantity int `json:"quantity"`

	// Price per unit of the product at the time of purchase.
	Price string `json:"price"`

	// Internal field, creation TS
	CreatedAt time.Time `json:"created_at"`

	// Internal field, last updated TS
	UpdatedAt time.Time `json:"updated_at"`

	// Internal field, UUID of the row
	Guid string `json:"guid"`
}

// buyItemResponseList response list
type buyItemResponseList struct {
	// Page information (if present)
	Page pageable.Page `json:"page,omitempty"`

	// Payload
	Data []buyItemResponse `json:"data"`
}
