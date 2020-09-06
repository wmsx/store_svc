package models

import "gorm.io/gorm"



type Store struct {
	Model
	Bulk       string `gorm:"not null;type:varchar(32);comment:桶名"`
	ObjectName string `gorm:"not null;type:varchar(256);comment:对象名"`
	Filename   string `gorm:"not null;type:varchar(256);comment:文件名"`
	Size       int64  `gorm:"not null;type:int;comment:大小"`
}

func AddStore(object *Store) error {
	return db.Create(object).Error
}

func GetStoresById(ids []int64) ([]*Store, error) {
	var objects []*Store
	err := db.Where("id in (?)", ids).Find(&objects).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return objects, nil
}
