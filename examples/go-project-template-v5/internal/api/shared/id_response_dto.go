package shared

// IdResponse - generic response for newly created objects
type IdResponse struct {
	// Returned ID of newly created object
	ID any `json:"id"`
}
