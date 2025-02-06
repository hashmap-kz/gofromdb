package dto

import "time"

type JobTitlesDto struct {
	RecordID  int
	TitleName string
	CreatedAt time.Time
	UpdatedAt time.Time
	Guid      string
}

type JobTitlesCreateDto struct {
	TitleName string
}

type JobTitlesUpdateDto struct {
	TitleName string
}
