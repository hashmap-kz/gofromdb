package dto

import "time"

type NoPkDto struct {
	EventTime time.Time
	Message   string
}

type NoPkCreateDto struct {
	EventTime time.Time
	Message   string
}

type NoPkUpdateDto struct {
	EventTime time.Time
	Message   string
}
