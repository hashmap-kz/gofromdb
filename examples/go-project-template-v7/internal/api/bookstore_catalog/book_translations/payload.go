package book_translations

import (
	"go-project-template-v7/pkg/pageable"
	"time"
)

// bookTranslationsCreateRequest
// Natural composite key: one translation per book and language.
type bookTranslationsCreateRequest struct {
	BookID          int64      `json:"book_id"`
	LanguageCode    string     `json:"language_code"`
	TranslatedTitle string     `json:"translated_title"`
	TranslatedBy    *string    `json:"translated_by"`
	ReleasedOn      *time.Time `json:"released_on"`
}

// bookTranslationsUpdateRequest
// Natural composite key: one translation per book and language.
type bookTranslationsUpdateRequest struct {
	TranslatedTitle string     `json:"translated_title"`
	TranslatedBy    *string    `json:"translated_by"`
	ReleasedOn      *time.Time `json:"released_on"`
}

// bookTranslationsResponse
// Natural composite key: one translation per book and language.
type bookTranslationsResponse struct {
	BookID          int64      `json:"book_id"`
	LanguageCode    string     `json:"language_code"`
	TranslatedTitle string     `json:"translated_title"`
	TranslatedBy    *string    `json:"translated_by"`
	ReleasedOn      *time.Time `json:"released_on"`
}

// bookTranslationsResponseList response list
type bookTranslationsResponseList struct {
	// Page information (if present)
	Page *pageable.Page `json:"page,omitempty"`

	// Payload
	Data []bookTranslationsResponse `json:"data"`
}
