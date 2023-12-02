package helpers

import (
	"math"

	"gorm.io/gorm"
)

type Pagination struct {
	Page         uint   `json:"page,omitempty"`
	Limit        uint   `json:"limit,omitempty"`
	TotalPage    uint   `json:"totalPage,omitempty"`
	TotalData    uint   `json:"totalRows,omitempty"`
	NextPage     string `json:"nextPage,omitempty"`
	PreviousPage string `json:"previousPage,omitempty"`
}

func (p *Pagination) GetPage() uint {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) GetOffset() uint {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetLimit() uint {
	if p.Limit == 0 {
		p.Limit = 5
	}
	return p.Limit
}

func Paginate(data interface{}, pagination *Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var TotalData int64

	db.Model(data).Count(&TotalData)
	pagination.TotalData = uint(TotalData)

	pagination.TotalPage = uint(math.Ceil(float64(TotalData) / float64(pagination.GetLimit())))

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(int(pagination.GetOffset())).Limit(int(pagination.GetLimit()))
	}
}
