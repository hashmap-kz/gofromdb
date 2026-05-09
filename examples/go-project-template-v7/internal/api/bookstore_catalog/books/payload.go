package books

import (
	"go-project-template-v7/pkg/pageable"
	"time"
)

// booksCreateRequest
// Books available in the store. Covers serial PK, arrays, jsonb, bytea,
// nullable fields, and generated columns.
type booksCreateRequest struct {
	PublisherCode string     `json:"publisher_code"`
	Isbn13        string     `json:"isbn13"`
	Title         string     `json:"title"`
	Subtitle      *string    `json:"subtitle"`
	Description   *string    `json:"description"`
	Price         string     `json:"price"`
	WeightGrams   *int       `json:"weight_grams"`
	Rating        *string    `json:"rating"`
	PublishedOn   *time.Time `json:"published_on"`
	Tags          []string   `json:"tags"`
	Attrs         string     `json:"attrs"`
	CoverImage    []byte     `json:"cover_image"`
	ArchivedAt    *time.Time `json:"archived_at"`
}

// booksUpdateRequest
// Books available in the store. Covers serial PK, arrays, jsonb, bytea,
// nullable fields, and generated columns.
type booksUpdateRequest struct {
	PublisherCode string     `json:"publisher_code"`
	Isbn13        string     `json:"isbn13"`
	Title         string     `json:"title"`
	Subtitle      *string    `json:"subtitle"`
	Description   *string    `json:"description"`
	Price         string     `json:"price"`
	WeightGrams   *int       `json:"weight_grams"`
	Rating        *string    `json:"rating"`
	PublishedOn   *time.Time `json:"published_on"`
	Tags          []string   `json:"tags"`
	Attrs         string     `json:"attrs"`
	CoverImage    []byte     `json:"cover_image"`
	ArchivedAt    *time.Time `json:"archived_at"`
}

// booksResponse
// Books available in the store. Covers serial PK, arrays, jsonb, bytea,
// nullable fields, and generated columns.
type booksResponse struct {
	BookID        int64      `json:"book_id"`
	PublisherCode string     `json:"publisher_code"`
	Isbn13        string     `json:"isbn13"`
	Title         string     `json:"title"`
	Subtitle      *string    `json:"subtitle"`
	Description   *string    `json:"description"`
	Price         string     `json:"price"`
	WeightGrams   *int       `json:"weight_grams"`
	Rating        *string    `json:"rating"`
	PublishedOn   *time.Time `json:"published_on"`
	Tags          []string   `json:"tags"`
	Attrs         string     `json:"attrs"`
	CoverImage    []byte     `json:"cover_image"`
	ArchivedAt    *time.Time `json:"archived_at"`
	TitleSearch   *string    `json:"title_search"`
}

// booksResponseList response list
type booksResponseList struct {
	// Page information (if present)
	Page *pageable.Page `json:"page,omitempty"`

	// Payload
	Data []booksResponse `json:"data"`
}
