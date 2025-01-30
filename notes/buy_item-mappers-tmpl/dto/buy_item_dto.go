package dto

import (
	"time"

	"go-project-template-v5/pkg/pageable"
)

// buyItemCreateRequest represents items in a purchase, including quantity and price.
type BuyItemCreateRequest struct {
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
type BuyItemUpdateRequest struct {
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
type BuyItemResponse struct {
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
type BuyItemResponseList struct {
	// Page information (if present)
	Page *pageable.Page `json:"page,omitempty"`

	// Payload
	Data []BuyItemResponse `json:"data"`
}

// Service layer

type BuyItemDto struct {
	RecordID  int
	BuyID     int
	ProductID int
	Quantity  int
	Price     string
	CreatedAt time.Time
	UpdatedAt time.Time
	Guid      string
}

type BuyItemCreateDto struct {
	BuyID     int
	ProductID int
	Quantity  int
	Price     string
}

type BuyItemUpdateDto struct {
	BuyID     int
	ProductID int
	Quantity  int
	Price     string
}
