package postgres

import "time"

type NoPk struct {
	EventTime time.Time `json:"event_time" db:"event_time"`
	Message   string    `json:"message" db:"message"`
}
