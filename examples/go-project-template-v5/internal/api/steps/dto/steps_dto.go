package dto

import "time"

type StepsDto struct {
	RecordID  int
	StepName  string
	CreatedAt time.Time
	UpdatedAt time.Time
	Guid      string
}

type StepsCreateDto struct {
	StepName string
}

type StepsUpdateDto struct {
	StepName string
}
