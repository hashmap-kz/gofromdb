package v1

import (
	"time"

	"go-project-template-v5/pkg/pageable"
)

// customerOrderItemsCreateRequest represents items in a sales, including quantity and price.
type customerOrderItemsCreateRequest struct {
	// Foreign key referencing the associated customer-order.
	CustomerOrderID int `json:"customer_order_id"`

	// Foreign key referencing the product.
	ProductID int `json:"product_id"`

	// Number of units of the product.
	Quantity int `json:"quantity"`

	// Price per unit of the product at the time of ordering.
	Price string `json:"price"`
}

// customerOrderItemsUpdateRequest represents items in a sales, including quantity and price.
type customerOrderItemsUpdateRequest struct {
	// Foreign key referencing the associated customer-order.
	CustomerOrderID int `json:"customer_order_id"`

	// Foreign key referencing the product.
	ProductID int `json:"product_id"`

	// Number of units of the product.
	Quantity int `json:"quantity"`

	// Price per unit of the product at the time of ordering.
	Price string `json:"price"`
}

// customerOrderItemsResponse represents items in a sales, including quantity and price.
type customerOrderItemsResponse struct {
	// Primary key for the sales_items table.
	RecordID int `json:"record_id"`

	// Foreign key referencing the associated customer-order.
	CustomerOrderID int `json:"customer_order_id"`

	// Foreign key referencing the product.
	ProductID int `json:"product_id"`

	// Number of units of the product.
	Quantity int `json:"quantity"`

	// Price per unit of the product at the time of ordering.
	Price string `json:"price"`

	// Internal field, creation TS
	CreatedAt time.Time `json:"created_at"`

	// Internal field, last updated TS
	UpdatedAt time.Time `json:"updated_at"`

	// Internal field, UUID of the row
	Guid string `json:"guid"`
}

// customerOrderItemsResponseList response list
type customerOrderItemsResponseList struct {
	// Page information (if present)
	Page *pageable.Page `json:"page,omitempty"`

	// Payload
	Data []customerOrderItemsResponse `json:"data"`
}
