package dto

import "time"

type NullableTypesDto struct {
	ID        int64
	Name      *string
	Amount    *string
	Payload   *string
	Tags      []string
	Active    *bool
	CreatedAt *time.Time
}

type NullableTypesCreateDto struct {
	Name    *string
	Amount  *string
	Payload *string
	Tags    []string
	Active  *bool
}

type NullableTypesUpdateDto struct {
	Name    *string
	Amount  *string
	Payload *string
	Tags    []string
	Active  *bool
}
