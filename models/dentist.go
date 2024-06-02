package models

import "gorm.io/gorm"

type Dentist struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
	CPF   string `json:"cpf"`
}

func MigrateDentistTable(db *gorm.DB) {
	db.AutoMigrate(&Dentist{})
}
