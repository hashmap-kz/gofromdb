package postgres

import "time"

type NullableTypes struct {
	ID        int64      `json:"id" db:"id"`
	Name      *string    `json:"name" db:"name"`
	Amount    *string    `json:"amount" db:"amount"`
	Payload   *string    `json:"payload" db:"payload"`
	Tags      []string   `json:"tags" db:"tags"`
	Active    *bool      `json:"active" db:"active"`
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
}
