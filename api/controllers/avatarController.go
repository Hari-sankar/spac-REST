package controllers

import (
	"net/http"
	"spac-REST/api/schemas"
	"spac-REST/api/usecases"
	"spac-REST/api/utils"

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

// CreateAvatar godoc
// @Summary Create a new avatar
// @Description Create an avatar with specified image URL and name
// @Tags avatars
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body schemas.CreateAvatarRequest true "Avatar creation request"
// @Success 200 {object} schemas.CreateAvatarResponse
// @Failure 400 {object} utils.ErrorStruct "Invalid request"
// @Failure 401 {object} utils.ErrorStruct "Unauthorized"
// @Failure 500 {object} utils.ErrorStruct "Internal Server Error"
// @Router /avatar [post]
func (ctrl *AvatarController) CreateAvatar(c *gin.Context) {
	var req schemas.CreateAvatarRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(utils.NewErrorStruct(400, err.Error()))
		return
	}

	response, err := ctrl.avatarUseCase.CreateAvatar(c.Request.Context(), &req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response)
}
