package controllers

import (
	"net/http"
	"spac-REST/api/usecases"

	"github.com/gin-gonic/gin"
)

type AvatarController struct {
	avatarUseCase *usecases.AvatarUseCase
}

func NewAvatarController(avatarUseCase *usecases.AvatarUseCase) *AvatarController {
	return &AvatarController{avatarUseCase: avatarUseCase}
}

// GetAvailableAvatars godoc
// @Summary Get available avatars
// @Description Get list of all available avatars
// @Tags avatars
// @Produce json
// @Security BearerAuth
// @Success 200 {object} schemas.GetAvatarsResponse
// @Failure 401 {object} utils.ErrorStruct "Unauthorized"
// @Failure 500 {object} utils.ErrorStruct "Internal Server Error"
// @Router /avatars [get]
func (ctrl *AvatarController) GetAvailableAvatars(c *gin.Context) {
	response, err := ctrl.avatarUseCase.GetAvailableAvatars(c.Request.Context())
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, response)
}
