package routes

import (
	"github.com/azevedoguigo/projeto-odonto-back.git/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/users", controllers.CreateUser)

	return router
}
