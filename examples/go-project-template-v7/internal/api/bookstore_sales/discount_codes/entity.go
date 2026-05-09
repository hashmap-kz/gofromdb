package discount_codes

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

// DiscountCodes
// Discount codes. Natural text primary key plus daterange.
type DiscountCodes struct {
	Code        string                  `json:"code" db:"code"`
	Description *string                 `json:"description" db:"description"`
	PercentOff  string                  `json:"percent_off" db:"percent_off"`
	ValidPeriod pgtype.Range[time.Time] `json:"valid_period" db:"valid_period"`
	MaxUses     *int                    `json:"max_uses" db:"max_uses"`
	Active      bool                    `json:"active" db:"active"`
}
