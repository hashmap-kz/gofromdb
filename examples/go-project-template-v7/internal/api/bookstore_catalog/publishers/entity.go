package publishers

import "time"

// Publishers
// Book publishers. Uses a natural text primary key.
type Publishers struct {
	// Natural publisher code, for example no_starch or manning.
	Code string `json:"code" db:"code"`

	Name        string     `json:"name" db:"name"`
	CountryCode string     `json:"country_code" db:"country_code"`
	Website     *string    `json:"website" db:"website"`
	FoundedOn   *time.Time `json:"founded_on" db:"founded_on"`
	Active      bool       `json:"active" db:"active"`
}
