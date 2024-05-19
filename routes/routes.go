package routes

import (
	"github.com/azevedoguigo/projeto-odonto-back.git/controllers"
	"github.com/azevedoguigo/projeto-odonto-back.git/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/login", controllers.Login)

	protected := router.Group("/", middleware.AuthMiddleware())
	{
		protected.POST("/users", controllers.CreateUser)
		protected.GET("/users/:id", controllers.GetUserById)
	}

	return router
}
