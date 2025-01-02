package routes

import (
	"spac-REST/api/controllers"
	middleware "spac-REST/api/middlewares"
	"spac-REST/api/usecases"

	"github.com/gin-gonic/gin"
)

func RegisterSpaceRoutes(router *gin.Engine, spaceUseCase usecases.SpaceUseCase) {
	spaceController := controllers.NewSpaceController(&spaceUseCase)

	spaceRoutes := router.Group("/space").Use(middleware.AuthenticateRequest)
	{
		spaceRoutes.POST("", spaceController.CreateSpace)
		spaceRoutes.DELETE("/:spaceId", spaceController.DeleteSpace)
		spaceRoutes.GET("/all", spaceController.GetAllSpaces)
	}
}
