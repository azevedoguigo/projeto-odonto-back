package config

import (
	"log"

	"github.com/azevedoguigo/projeto-odonto-back.git/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=postgres password=postgres dbname=projeto_odonto port=5432 sslmode=disable TimeZone=America/Maceio"

	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB.AutoMigrate(&models.User{}, &models.Dentist{}, &models.Patient{}, &models.Service{})
}
