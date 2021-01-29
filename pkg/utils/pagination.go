package utils

import (
	"go.mongodb.org/mongo-driver/mongo/options"
	"math"
)

type Pagination struct {
	defaultSize int64
	page int64
	size int64
	total int64
	pages int64
}

type PaginationData struct {
	Page int64 `json:"page"`
	Size int64 `json:"size"`
	Total int64 `json:"total"`
	Pages int64 `json:"pages"`
}

func NewPagination(defaultSize int64) *Pagination {
	if defaultSize <= 0 {
		defaultSize = 10
	}
	return &Pagination{
		defaultSize:defaultSize,
		page:      1,
		size:  defaultSize,
		total:     0,
		pages: 		0,
	}
}


func (p *Pagination)SetTotal(total int64) {
	p.total = total
}

func (p *Pagination)SetPage(page int64) {
	if page <= 0 {
		page = 1
	}
	p.page = page
}

func (p *Pagination)SetSize(size int64) {
	if size <= 0 {
		size = p.defaultSize
	}

	p.size = size
}


func (p *Pagination)GetPaginationData() PaginationData{
	if p.size == 0 {
		p.size = p.defaultSize
	}
	if p.total != 0 {
		p.pages = int64(math.Ceil(float64(p.total) / float64(p.size)))
	}

	return PaginationData{
		Page:      p.page,
		Size:      p.size,
		Total:     p.total,
		Pages: 	   p.pages,
	}
}

func (p *Pagination)GetFindOptions() *options.FindOptions{
	findOptions := options.Find()

	if p.size > 0 {
		findOptions.SetLimit(p.size)
		findOptions.SetSkip((p.page - 1) * p.size)
	}

	return findOptions
}


