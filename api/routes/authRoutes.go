package routes

import (
	"spac-REST/api/controllers"
	"spac-REST/api/usecases"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.Engine, authUseCases usecases.AuthUseCase) {
	authController := controllers.NewAuthController(authUseCases)

	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/signup", authController.SignUp)
		userRoutes.GET("/signin", authController.SignIn)

	}
}
