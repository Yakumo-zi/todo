package persist

import (
	"todo/model"

	"gorm.io/gorm"
)

func WithStatus(status ...model.Status) DBOption {
	return func(db *gorm.DB) *gorm.DB {
		if len(status) > 1 {
			return db.Where("status IN ?", status)
		}
		if len(status) == 0 {
			return db
		}
		return db.Where("status = ?", status[0])
	}
}
func WithComments(comments string) DBOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("comments LIKE ?", "%"+comments+"%")
	}
}
