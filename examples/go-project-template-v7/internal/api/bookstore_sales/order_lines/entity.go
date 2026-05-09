package order_lines

// OrderLines
// Order lines. Tests composite primary key order: order_id, line_no.
type OrderLines struct {
	OrderID        int64   `json:"order_id" db:"order_id"`
	LineNo         int     `json:"line_no" db:"line_no"`
	BookID         int64   `json:"book_id" db:"book_id"`
	Quantity       int16   `json:"quantity" db:"quantity"`
	UnitPrice      string  `json:"unit_price" db:"unit_price"`
	DiscountAmount string  `json:"discount_amount" db:"discount_amount"`
	Note           *string `json:"note" db:"note"`
}
