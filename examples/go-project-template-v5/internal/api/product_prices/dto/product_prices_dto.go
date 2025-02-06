package dto

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type ProductPricesDto struct {
	RecordID           int
	ProductPricePeriod pgtype.Range[time.Time]
	ProductID          int
	ProductPrice       string
	CreatedAt          time.Time
	UpdatedAt          time.Time
	Guid               string
}

type ProductPricesCreateDto struct {
	ProductPricePeriod pgtype.Range[time.Time]
	ProductID          int
	ProductPrice       string
}

type ProductPricesUpdateDto struct {
	ProductPricePeriod pgtype.Range[time.Time]
	ProductID          int
	ProductPrice       string
}
