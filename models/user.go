package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"unique_index" json:"name"`
	Password string `json:"password"`
}
