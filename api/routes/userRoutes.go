package routes

import (
	"spac-REST/api/controllers"
	"spac-REST/api/usecases"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine, userUseCase usecases.UserUseCase) {
	userController := controllers.NewUserController(userUseCase)

	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/", userController.CreateUser)
		userRoutes.GET("/:id", userController.GetUserByID)
		userRoutes.GET("/", userController.GetAllUsers)
		userRoutes.PUT("/:id", userController.UpdateUser)
		userRoutes.DELETE("/:id", userController.DeleteUser)
	}
}
