package models

import "github.com/jinzhu/gorm"

type Object struct {
	gorm.Model
	Bulk       string
	ObjectName string
	Filename   string
	Size       int64
	Status     byte
}

func BatchAddObject(objects []Object) error  {
	if len(objects) == 0  {
		return nil
	}
	return db.Create(&objects).Error
}
