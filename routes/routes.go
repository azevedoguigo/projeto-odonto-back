package routes

import (
	"time"

	"github.com/azevedoguigo/projeto-odonto-back.git/controllers"
	"github.com/azevedoguigo/projeto-odonto-back.git/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.POST("/login", controllers.Login)
	router.POST("/users", controllers.CreateUser)

	protected := router.Group("/", middleware.AuthMiddleware())
	{
		protected.GET("/users/:id", controllers.GetUserById)

		protected.POST("/patients", controllers.CreatePatient)
		protected.GET("/patients/:id", controllers.GetPatientById)

		protected.POST("/dentists", controllers.CreateDentist)
		protected.GET("/dentists/:id", controllers.GetDentistById)
	}

	return router
}
