package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `gorm:"index,type:VARCHAR(255)"`
	Age uint `gorm:"int(3)"`
}
