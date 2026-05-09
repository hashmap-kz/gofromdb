package book_authors

// BookAuthors
// Many-to-many link. Tests composite primary key order: book_id, author_id.
type BookAuthors struct {
	BookID            int64   `json:"book_id" db:"book_id"`
	AuthorID          string  `json:"author_id" db:"author_id"`
	ContributionOrder int16   `json:"contribution_order" db:"contribution_order"`
	Role              string  `json:"role" db:"role"`
	Notes             *string `json:"notes" db:"notes"`
}
