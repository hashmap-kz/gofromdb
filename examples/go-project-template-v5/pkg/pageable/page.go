package pageable

type Page struct {
	// Current page size
	Size int `json:"size,omitempty"`
	// Total elements in all pages
	TotalElements int `json:"totalElements,omitempty"`
	// Total pages count
	TotalPages int `json:"totalPages,omitempty"`
	// Current page number
	Number int `json:"number,omitempty"`
}

func CreatePage(pq *PaginationQuery, totalCount int) Page {
	return Page{
		Size:          pq.GetSize(),
		TotalElements: totalCount,
		TotalPages:    GetTotalPages(totalCount, pq.GetSize()),
		Number:        pq.GetNumber(),
	}
}
