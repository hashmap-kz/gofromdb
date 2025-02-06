package dto

import "time"

type DepartmentsDto struct {
	RecordID       int
	DepartmentName string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Guid           string
}

type DepartmentsCreateDto struct {
	DepartmentName string
}

type DepartmentsUpdateDto struct {
	DepartmentName string
}
