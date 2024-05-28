package models

import (
	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model
	Name            string `json:"name"`
	BirthDate       string `json:"birth-date"`
	PhoneNumber     string `json:"phone-number"`
	HealthInsurance string `json:"healt-insurance"`
	Expiration      string `json:"expiration"`
	CPF             string `json:"cpf"`
	Address         string `json:"addres"`
}

func MigratePatientTable(db *gorm.DB) {
	db.AutoMigrate(&Patient{})
}
