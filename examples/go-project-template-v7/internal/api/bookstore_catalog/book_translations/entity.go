package book_translations

import "time"

// BookTranslations
// Natural composite key: one translation per book and language.
type BookTranslations struct {
	BookID          int64      `json:"book_id" db:"book_id"`
	LanguageCode    string     `json:"language_code" db:"language_code"`
	TranslatedTitle string     `json:"translated_title" db:"translated_title"`
	TranslatedBy    *string    `json:"translated_by" db:"translated_by"`
	ReleasedOn      *time.Time `json:"released_on" db:"released_on"`
}
