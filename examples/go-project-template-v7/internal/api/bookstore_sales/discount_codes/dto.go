package discount_codes

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Dto struct {
	Code        string
	Description *string
	PercentOff  string
	ValidPeriod pgtype.Range[time.Time]
	MaxUses     *int
	Active      bool
}

type CreateDto struct {
	Code        string
	Description *string
	PercentOff  string
	ValidPeriod pgtype.Range[time.Time]
	MaxUses     *int
	Active      bool
}

type UpdateDto struct {
	Description *string
	PercentOff  *string
	ValidPeriod *pgtype.Range[time.Time]
	MaxUses     *int
	Active      *bool
}
