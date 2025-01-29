package dto

import "time"

type CategoryDto struct {
	RecordID  int
	Name      string
	ParentID  *int
	CreatedAt time.Time
	UpdatedAt time.Time
	Guid      string
}

type CategoryCreateDto struct {
	Name     string
	ParentID *int
}

type CategoryUpdateDto struct {
	RecordID int
	Name     string
	ParentID *int
}
