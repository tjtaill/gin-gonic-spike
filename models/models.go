package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func Register() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "spike.db")
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&User{})
	return db, nil
}
