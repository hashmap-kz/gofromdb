package dto

type SerialPkDto struct {
	ID   int64
	Name string
}

type SerialPkCreateDto struct {
	Name string
}

type SerialPkUpdateDto struct {
	Name string
}
