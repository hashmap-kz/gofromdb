package warehouses

// Warehouses
// Warehouses. Natural text primary key.
type Warehouses struct {
	Code     string `json:"code" db:"code"`
	Name     string `json:"name" db:"name"`
	Address  string `json:"address" db:"address"`
	Timezone string `json:"timezone" db:"timezone"`
	Active   bool   `json:"active" db:"active"`
}
