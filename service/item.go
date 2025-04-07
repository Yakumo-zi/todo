package service

import (
	"todo/model"
	"todo/persist"

	"gorm.io/gorm"
)

type ItemService struct {
	db *gorm.DB
}

func NewItemService(db *gorm.DB) *ItemService {
	return &ItemService{
		db: db,
	}
}

func (s *ItemService) Create(item model.Item) error {
	return s.db.Create(&item).Error
}

func (s *ItemService) Get(options ...persist.DBOption) (model.Item, error) {
	var item model.Item
	err := persist.ApplyOptions(s.db, options...).First(&item).Error
	if err != nil {
		return model.Item{}, err
	}
	return item, nil
}

func (s *ItemService) Update(item model.Item, options ...persist.DBOption) error {
	return persist.ApplyOptions(s.db, options...).Save(&item).Error
}

func (s *ItemService) Delete(options ...persist.DBOption) error {
	return persist.ApplyOptions(s.db, options...).Delete(&model.Item{}).Error
}

func (s *ItemService) List(options ...persist.DBOption) ([]model.Item, error) {
	var items []model.Item
	err := persist.ApplyOptions(s.db, options...).Find(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}
