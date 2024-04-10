package dto

type Pagination struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
}

// GetPageNum offset的偏移量要-1
func (p *Pagination) GetPageNum() int {
	if p.PageNum <= 0 {
		p.PageNum = 0
	}
	return p.PageNum - 1
}

func (p *Pagination) GetPageSize() int {
	if p.PageSize <= 0 {
		p.PageSize = 1
	}
	return p.PageSize
}
