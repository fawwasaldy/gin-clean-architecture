package pagination

import (
	"gorm.io/gorm"
	"math"
)

func Paginate(req Request) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (req.Page - 1) * req.PerPage
		return db.Offset(offset).Limit(req.PerPage)
	}
}

func TotalPage(count, perPage int64) int64 {
	totalPage := int64(math.Ceil(float64(count) / float64(perPage)))

	return totalPage
}
