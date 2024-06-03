package models

import "gorm.io/gorm"

type Service struct {
	gorm.Model

	DentistId uint `json:"dentist_id"`
	PatientId uint `json:"patient_id"`
}
