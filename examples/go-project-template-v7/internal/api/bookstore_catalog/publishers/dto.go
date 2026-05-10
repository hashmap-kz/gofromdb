package publishers

import "time"

type Dto struct {
	Code        string
	Name        string
	CountryCode string
	Website     *string
	FoundedOn   *time.Time
	Active      bool
}

type CreateDto struct {
	Code        string
	Name        string
	CountryCode string
	Website     *string
	FoundedOn   *time.Time
	Active      bool
}

type UpdateDto struct {
	Name        *string
	CountryCode *string
	Website     *string
	FoundedOn   *time.Time
	Active      *bool
}
