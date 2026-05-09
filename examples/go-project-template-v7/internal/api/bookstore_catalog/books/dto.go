package books

import "time"

type Dto struct {
	BookID        int64
	PublisherCode string
	Isbn13        string
	Title         string
	Subtitle      *string
	Description   *string
	Price         string
	WeightGrams   *int
	Rating        *string
	PublishedOn   *time.Time
	Tags          []string
	Attrs         string
	CoverImage    []byte
	ArchivedAt    *time.Time
	TitleSearch   *string
}

type CreateDto struct {
	PublisherCode string
	Isbn13        string
	Title         string
	Subtitle      *string
	Description   *string
	Price         string
	WeightGrams   *int
	Rating        *string
	PublishedOn   *time.Time
	Tags          []string
	Attrs         string
	CoverImage    []byte
	ArchivedAt    *time.Time
}

type UpdateDto struct {
	PublisherCode string
	Isbn13        string
	Title         string
	Subtitle      *string
	Description   *string
	Price         string
	WeightGrams   *int
	Rating        *string
	PublishedOn   *time.Time
	Tags          []string
	Attrs         string
	CoverImage    []byte
	ArchivedAt    *time.Time
}
