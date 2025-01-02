package routes

import (
	"spac-REST/api/controllers"
	middleware "spac-REST/api/middlewares"
	"spac-REST/api/usecases"

	"github.com/gin-gonic/gin"
)

func RegisterMapRoutes(router *gin.Engine, mapUseCase usecases.MapUseCase) {
	mapController := controllers.NewMapController(&mapUseCase)

	mapRoutes := router.Group("/map").Use(middleware.AuthenticateRequest)
	{
		mapRoutes.POST("/", mapController.CreateMap)
		mapRoutes.POST("/:mapId/elements", mapController.AddMapElements)
		mapRoutes.GET("/all", mapController.GetAllMaps)
	}
}
