package routes

import (
	"github.com/azevedoguigo/projeto-odonto-back.git/controllers"
	"github.com/azevedoguigo/projeto-odonto-back.git/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/login", controllers.Login)
	router.POST("/users", controllers.CreateUser)

	protected := router.Group("/", middleware.AuthMiddleware())
	{
		protected.GET("/users/:id", controllers.GetUserById)
		protected.POST("/patients", controllers.CreatePatient)
	}

	return router
}
