package v1

import (
	"time"

	"go-project-template-v5/pkg/pageable"
)

type buyItemCreateRequest struct {
	BuyID     int    `json:"buy_id"`
	ProductID int    `json:"product_id"`
	Quantity  int    `json:"quantity"`
	Price     string `json:"price"`
}

type buyItemResponse struct {
	RecordID  int       `json:"record_id"`
	BuyID     int       `json:"buy_id"`
	ProductID int       `json:"product_id"`
	Quantity  int       `json:"quantity"`
	Price     string    `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Guid      string    `json:"guid"`
}

type buyItemResponseList struct {
	Page pageable.Page
	Data []buyItemResponse `json:"data"`
}

type buyItemUpdateRequest struct {
	RecordID  int    `json:"record_id"`
	BuyID     int    `json:"buy_id"`
	ProductID int    `json:"product_id"`
	Quantity  int    `json:"quantity"`
	Price     string `json:"price"`
}
