package customers

import "time"

// Customers
// Store customers. UUID primary key and simple unique field.
type Customers struct {
	CustomerID     string    `json:"customer_id" db:"customer_id"`
	Email          string    `json:"email" db:"email"`
	FullName       string    `json:"full_name" db:"full_name"`
	Phone          *string   `json:"phone" db:"phone"`
	MarketingOptIn bool      `json:"marketing_opt_in" db:"marketing_opt_in"`
	RegisteredAt   time.Time `json:"registered_at" db:"registered_at"`
}
