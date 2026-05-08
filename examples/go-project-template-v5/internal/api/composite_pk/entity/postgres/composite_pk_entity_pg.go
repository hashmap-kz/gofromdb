package postgres

type CompositePk struct {
	TenantID int64  `json:"tenant_id" db:"tenant_id"`
	Code     string `json:"code" db:"code"`
	Name     string `json:"name" db:"name"`
}
