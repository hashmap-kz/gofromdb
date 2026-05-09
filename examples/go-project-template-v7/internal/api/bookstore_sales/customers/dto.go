package customers

import "time"

type Dto struct {
	CustomerID     string
	Email          string
	FullName       string
	Phone          *string
	MarketingOptIn bool
	RegisteredAt   time.Time
}

type CreateDto struct {
	CustomerID     string
	Email          string
	FullName       string
	Phone          *string
	MarketingOptIn bool
	RegisteredAt   time.Time
}

type UpdateDto struct {
	Email          string
	FullName       string
	Phone          *string
	MarketingOptIn bool
	RegisteredAt   time.Time
}
