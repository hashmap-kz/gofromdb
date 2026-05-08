package dto

type CompositePkDto struct {
	TenantID int64
	Code     string
	Name     string
}

type CompositePkCreateDto struct {
	TenantID int64
	Code     string
	Name     string
}

type CompositePkUpdateDto struct {
	Name string
}
