package dto

import "time"

type PurchaseStepsDto struct {
	RecordID  int
	StepName  string
	CreatedAt time.Time
	UpdatedAt time.Time
	Guid      string
}

type PurchaseStepsCreateDto struct {
	StepName string
}

type PurchaseStepsUpdateDto struct {
	StepName string
}
