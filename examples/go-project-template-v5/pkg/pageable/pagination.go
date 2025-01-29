package pageable

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
)

const (
	defaultSize = 10
)

// PaginationQuery params
type PaginationQuery struct {
	Size int `json:"size,omitempty"`
	Page int `json:"page,omitempty"`
}

func (q *PaginationQuery) SetSize(sizeQuery string) error {
	if sizeQuery == "" {
		q.Size = defaultSize
		return nil
	}
	n, err := strconv.Atoi(sizeQuery)
	if err != nil {
		return err
	}
	q.Size = n

	return nil
}

func (q *PaginationQuery) GetSize() int {
	return q.Size
}

func (q *PaginationQuery) SetNumber(pageQuery string) error {
	if pageQuery == "" {
		q.Size = 0
		return nil
	}
	n, err := strconv.Atoi(pageQuery)
	if err != nil {
		return err
	}
	q.Page = n

	return nil
}

func (q *PaginationQuery) GetNumber() int {
	return q.Page
}

func (q *PaginationQuery) GetOffset() int {
	if q.Page == 0 {
		return 0
	}
	return (q.Page - 1) * q.Size
}

func (q *PaginationQuery) GetLimit() int {
	return q.Size
}

func (q *PaginationQuery) GetQueryString() string {
	return fmt.Sprintf("page=%v&size=%v", q.GetNumber(), q.GetSize())
}

// Get pagination query struct from request params
func GetPaginationFromCtx(r *http.Request) (*PaginationQuery, error) {
	q := &PaginationQuery{}
	if err := q.SetNumber(r.URL.Query().Get("page")); err != nil {
		return nil, err
	}
	if err := q.SetSize(r.URL.Query().Get("size")); err != nil {
		return nil, err
	}

	return q, nil
}

func GetTotalPages(totalCount, pageSize int) int {
	d := float64(totalCount) / float64(pageSize)
	return int(math.Ceil(d))
}
