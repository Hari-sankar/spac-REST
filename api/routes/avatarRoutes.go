package routes

import (
	"spac-REST/api/controllers"
	middleware "spac-REST/api/middlewares"
	"spac-REST/api/usecases"

	"github.com/gin-gonic/gin"
)

func RegisterAvatarRoutes(router *gin.Engine, avatarUseCase usecases.AvatarUseCase) {
	avatarController := controllers.NewAvatarController(&avatarUseCase)

	avatarRoutes := router.Group("/avatars").Use(middleware.AuthenticateRequest)
	{
		avatarRoutes.GET("", avatarController.GetAvailableAvatars)

	}
}
