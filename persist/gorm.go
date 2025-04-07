package persist

import "gorm.io/gorm"

type DBOption func(*gorm.DB) *gorm.DB

func WithName(name string) DBOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("name LIKE ?", "%"+name+"%")
	}
}
func WithId(id uint) DBOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", id)
	}
}
func WithLimit(limit int) DBOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Limit(limit)
	}
}
func WithOffset(offset int) DBOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset)
	}
}
func WithOrder(order string, desc bool) DBOption {
	return func(db *gorm.DB) *gorm.DB {
		if desc {
			order = order + " DESC"
		}
		return db.Order(order)
	}
}
func ApplyOptions(db *gorm.DB, options ...DBOption) *gorm.DB {
	for _, option := range options {
		db = option(db)
	}
	return db
}
