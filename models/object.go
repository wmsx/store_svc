package models

import "gorm.io/gorm"

type Object struct {
	Model
	Bulk       string
	ObjectName string
	Filename   string
	Size       int64
	Status     byte
}

func BatchAddObject(objects []*Object) error {
	if len(objects) == 0 {
		return nil
	}
	return db.Create(&objects).Error
}

func GetObjectsById(ids []int64) ([]*Object, error) {
	var objects []*Object
	err := db.Where("id in (?)", ids).Find(&objects).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return objects, nil

}
