package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"unique_index" json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}
