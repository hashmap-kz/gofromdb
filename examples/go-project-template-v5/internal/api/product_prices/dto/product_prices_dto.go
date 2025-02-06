package dto

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type ProductPricesDto struct {
	RecordID     int
	ValidPeriod  pgtype.Range[time.Time]
	ProductID    int
	ProductPrice string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Guid         string
}

type ProductPricesCreateDto struct {
	ValidPeriod  pgtype.Range[time.Time]
	ProductID    int
	ProductPrice string
}

type ProductPricesUpdateDto struct {
	ValidPeriod  pgtype.Range[time.Time]
	ProductID    int
	ProductPrice string
}
