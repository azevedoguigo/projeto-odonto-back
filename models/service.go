package models

import "gorm.io/gorm"

type Service struct {
	gorm.Model

	DentistId uint `json:"dentist-id"`
	PatientId uint `json:"patient-id"`
}
