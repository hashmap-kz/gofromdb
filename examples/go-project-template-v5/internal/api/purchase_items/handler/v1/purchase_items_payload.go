package v1

import (
	"time"

	"go-project-template-v5/pkg/pageable"
)

// purchaseItemsCreateRequest represents items in a purchase, including quantity and price.
type purchaseItemsCreateRequest struct {
	// Foreign key referencing the associated purchase.
	PurchaseID int `json:"purchase_id"`

	// Foreign key referencing the purchased product.
	ProductID int `json:"product_id"`

	// Number of units of the product in the purchase.
	Quantity int `json:"quantity"`

	// Price per unit of the product at the time of purchase.
	Price string `json:"price"`
}

// purchaseItemsUpdateRequest represents items in a purchase, including quantity and price.
type purchaseItemsUpdateRequest struct {
	// Foreign key referencing the associated purchase.
	PurchaseID int `json:"purchase_id"`

	// Foreign key referencing the purchased product.
	ProductID int `json:"product_id"`

	// Number of units of the product in the purchase.
	Quantity int `json:"quantity"`

	// Price per unit of the product at the time of purchase.
	Price string `json:"price"`
}

// purchaseItemsResponse represents items in a purchase, including quantity and price.
type purchaseItemsResponse struct {
	// Primary key for the buy_item table.
	RecordID int `json:"record_id"`

	// Foreign key referencing the associated purchase.
	PurchaseID int `json:"purchase_id"`

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

// purchaseItemsResponseList response list
type purchaseItemsResponseList struct {
	// Page information (if present)
	Page *pageable.Page `json:"page,omitempty"`

	// Payload
	Data []purchaseItemsResponse `json:"data"`
}
