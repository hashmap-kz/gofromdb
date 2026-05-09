package books

import "time"

// Books
// Books available in the store. Covers serial PK, arrays, jsonb, bytea,
// nullable fields, and generated columns.
type Books struct {
	BookID        int64      `json:"book_id" db:"book_id"`
	PublisherCode string     `json:"publisher_code" db:"publisher_code"`
	Isbn13        string     `json:"isbn13" db:"isbn13"`
	Title         string     `json:"title" db:"title"`
	Subtitle      *string    `json:"subtitle" db:"subtitle"`
	Description   *string    `json:"description" db:"description"`
	Price         string     `json:"price" db:"price"`
	WeightGrams   *int       `json:"weight_grams" db:"weight_grams"`
	Rating        *string    `json:"rating" db:"rating"`
	PublishedOn   *time.Time `json:"published_on" db:"published_on"`
	Tags          []string   `json:"tags" db:"tags"`
	Attrs         string     `json:"attrs" db:"attrs"`
	CoverImage    []byte     `json:"cover_image" db:"cover_image"`
	ArchivedAt    *time.Time `json:"archived_at" db:"archived_at"`
	TitleSearch   *string    `json:"title_search" db:"title_search"`
}
