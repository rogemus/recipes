package models

import "math"

type Pagination struct {
	Page       int   `json:"page"`
	PageSize   int   `json:"page_size"`
	Pages      []int `json:"pages"`
	PagesCount int   `json:"pages_count"`
	ItemsCount int   `json:"items_count"`
}

func (p *Pagination) makePages() {
	pages := make([]int, p.PagesCount)

	for i := 0; i < p.PagesCount; i++ {
		pages[i] = 1 + i
	}

	p.Pages = pages
}

func (p *Pagination) CountPages(pageSize int) {
	p.PagesCount = int(math.Ceil(float64(p.ItemsCount) / float64(pageSize)))
	p.makePages()
}
