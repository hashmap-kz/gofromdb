package postgres

type NaturalPk struct {
	Code string `json:"code" db:"code"`
	Name string `json:"name" db:"name"`
}
