package routes

import (
	"spac-REST/api/controllers"
	middleware "spac-REST/api/middlewares"
	"spac-REST/api/usecases"

	"github.com/gin-gonic/gin"
)

func RegisterElementRoutes(router *gin.Engine, elementUseCase usecases.ElementUseCase) {
	elementController := controllers.NewElementController(&elementUseCase)

	elementRoutes := router.Group("/element").Use(middleware.AuthenticateRequest)
	{
		elementRoutes.POST("/", elementController.CreateElement)
		elementRoutes.PUT("/:elementId", elementController.UpdateElement)
		elementRoutes.GET("/", elementController.GetAllElements)

	}
}
