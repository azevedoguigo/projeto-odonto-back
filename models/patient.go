package models

import (
	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model
	Name            string    `json:"name"`
	BirthDate       string    `json:"birth_date"`
	PhoneNumber     string    `json:"phone_number"`
	HealthInsurance string    `json:"healt_insurance"`
	Expiration      string    `json:"expiration"`
	CPF             string    `json:"cpf"`
	Address         string    `json:"addres"`
	Services        []Service `json:"services"`
}
