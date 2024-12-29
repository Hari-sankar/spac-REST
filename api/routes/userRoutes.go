package routes

import (
	"spac-REST/api/controllers"
	middleware "spac-REST/api/middlewares"
	"spac-REST/api/usecases"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine, userUseCase usecases.UserUseCase) {
	userController := controllers.NewUserController(&userUseCase)

	userRoutes := router.Group("/users").Use(middleware.AuthenticateRequest)
	{
		userRoutes.POST("/metadata", userController.UpdateMetadata)
		userRoutes.GET("/metadata/bulk", userController.GetUserMetadataBulk)
	}
}
