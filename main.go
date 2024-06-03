package main

import (
	"github.com/azevedoguigo/projeto-odonto-back.git/config"
	"github.com/azevedoguigo/projeto-odonto-back.git/routes"
)

func main() {
	config.InitDB()

	router := routes.SetupRouter()
	router.Run()
}
