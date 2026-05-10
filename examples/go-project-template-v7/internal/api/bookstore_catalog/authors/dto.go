package authors

import "time"

type Dto struct {
	AuthorID    string
	DisplayName string
	LegalName   *string
	Biography   *string
	Metadata    string
	Active      bool
	BornOn      *time.Time
}

type CreateDto struct {
	AuthorID    string
	DisplayName string
	LegalName   *string
	Biography   *string
	Metadata    string
	Active      bool
	BornOn      *time.Time
}

type UpdateDto struct {
	DisplayName *string
	LegalName   *string
	Biography   *string
	Metadata    *string
	Active      *bool
	BornOn      *time.Time
}
