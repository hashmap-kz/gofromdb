package book_translations

import "time"

type Dto struct {
	BookID          int64
	LanguageCode    string
	TranslatedTitle string
	TranslatedBy    *string
	ReleasedOn      *time.Time
}

type CreateDto struct {
	BookID          int64
	LanguageCode    string
	TranslatedTitle string
	TranslatedBy    *string
	ReleasedOn      *time.Time
}

type UpdateDto struct {
	TranslatedTitle *string
	TranslatedBy    *string
	ReleasedOn      *time.Time
}
