package order_items

import (
	"go-project-template-v7/pkg/pageable"
	"time"
)

// orderItemsCreateRequest
// Represents items in a sales, including quantity and price.
type orderItemsCreateRequest struct {
	// Foreign key referencing the associated order.
	OrderID int `json:"order_id"`

	// Foreign key referencing the product.
	ProductID int `json:"product_id"`

	// Number of units of the product.
	Quantity string `json:"quantity"`

	// Price per unit of the product at the time of ordering.
	Price string `json:"price"`
}

// orderItemsUpdateRequest
// Represents items in a sales, including quantity and price.
type orderItemsUpdateRequest struct {
	// Foreign key referencing the associated order.
	OrderID *int `json:"order_id"`

	// Foreign key referencing the product.
	ProductID *int `json:"product_id"`

	// Number of units of the product.
	Quantity *string `json:"quantity"`

	// Price per unit of the product at the time of ordering.
	Price *string `json:"price"`
}

// orderItemsResponse
// Represents items in a sales, including quantity and price.
type orderItemsResponse struct {
	// Primary key.
	RecordID int `json:"record_id"`

	// Foreign key referencing the associated order.
	OrderID int `json:"order_id"`

	// Foreign key referencing the product.
	ProductID int `json:"product_id"`

	// Number of units of the product.
	Quantity string `json:"quantity"`

	// Price per unit of the product at the time of ordering.
	Price string `json:"price"`

	// Internal field, creation TS
	CreatedAt time.Time `json:"created_at"`

	// Internal field, last updated TS
	UpdatedAt time.Time `json:"updated_at"`

	// Internal field, UUID of the row
	GUID string `json:"guid"`
}

// orderItemsResponseList response list
type orderItemsResponseList struct {
	// Page information (if present)
	Page *pageable.Page `json:"page,omitempty"`

	// Payload
	Data []orderItemsResponse `json:"data"`
}
