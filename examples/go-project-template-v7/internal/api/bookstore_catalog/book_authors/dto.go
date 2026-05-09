package book_authors

type Dto struct {
	BookID            int64
	AuthorID          string
	ContributionOrder int16
	Role              string
	Notes             *string
}

type CreateDto struct {
	BookID            int64
	AuthorID          string
	ContributionOrder int16
	Role              string
	Notes             *string
}

type UpdateDto struct {
	ContributionOrder int16
	Role              string
	Notes             *string
}
