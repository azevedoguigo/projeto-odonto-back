package main

import (
	"github.com/azevedoguigo/projeto-odonto-back.git/config"
	"github.com/azevedoguigo/projeto-odonto-back.git/models"
	"github.com/azevedoguigo/projeto-odonto-back.git/routes"
)

func main() {
	config.InitDB()
	models.MigrateUserTable(config.DB)
	models.MigratePatientTable(config.DB)
	models.MigrateDentistTable(config.DB)

	router := routes.SetupRouter()
	router.Run()
}
