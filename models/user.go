package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&User{})
}
