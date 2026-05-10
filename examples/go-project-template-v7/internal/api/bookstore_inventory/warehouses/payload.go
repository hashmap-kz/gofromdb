package warehouses

import (
	"go-project-template-v7/pkg/pageable"
)

// warehousesCreateRequest
// Warehouses. Natural text primary key.
type warehousesCreateRequest struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Timezone string `json:"timezone"`
	Active   bool   `json:"active"`
}

// warehousesUpdateRequest
// Warehouses. Natural text primary key.
type warehousesUpdateRequest struct {
	Name     *string `json:"name"`
	Address  *string `json:"address"`
	Timezone *string `json:"timezone"`
	Active   *bool   `json:"active"`
}

// warehousesResponse
// Warehouses. Natural text primary key.
type warehousesResponse struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Timezone string `json:"timezone"`
	Active   bool   `json:"active"`
}

// warehousesResponseList response list
type warehousesResponseList struct {
	// Page information (if present)
	Page *pageable.Page `json:"page,omitempty"`

	// Payload
	Data []warehousesResponse `json:"data"`
}
