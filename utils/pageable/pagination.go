package pageable

type Pagination struct {
	Page       int `json:"page"`
	Size       int `json:"size"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

func NewPagination(total int64, page Pageable) *Pagination {
	size := int64(page.Size())
	mod := total % size
	totalPages := total / size
	if totalPages == 0 {
		totalPages = 1
	}
	if mod > 0 && total > size {
		totalPages += 1
	}
	return &Pagination{Total: int(total), TotalPages: int(totalPages), Page: page.Page(), Size: page.Size()}
}
