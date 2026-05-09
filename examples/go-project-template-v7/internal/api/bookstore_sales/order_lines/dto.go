package order_lines

type Dto struct {
	OrderID        int64
	LineNo         int
	BookID         int64
	Quantity       int16
	UnitPrice      string
	DiscountAmount string
	Note           *string
}

type CreateDto struct {
	OrderID        int64
	LineNo         int
	BookID         int64
	Quantity       int16
	UnitPrice      string
	DiscountAmount string
	Note           *string
}

type UpdateDto struct {
	BookID         int64
	Quantity       int16
	UnitPrice      string
	DiscountAmount string
	Note           *string
}
