package controllers

import (
	"net/http"
	"spac-REST/api/schemas"
	"spac-REST/api/usecases"
	"spac-REST/api/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUseCase *usecases.UserUseCase
}

func NewUserController(userUseCase *usecases.UserUseCase) *UserController {
	return &UserController{
		userUseCase: userUseCase,
	}
}

// UpdateMetadata godoc
// @Summary Update user metadata
// @Description Update user's avatar and other metadata
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body schemas.UpdateMetadataRequest true "Update metadata request"
// @Success 200 {object} schemas.UpdateMetadataResponse
// @Failure 400 {object} utils.ErrorStruct "Invalid request"
// @Failure 401 {object} utils.ErrorStruct "Invalid token"
// @Failure 404 {object} utils.ErrorStruct "User not found"
// @Router /users/metadata [post]
func (ctrl *UserController) UpdateMetadata(c *gin.Context) {
	ctx := c.Request.Context()
	var req schemas.UpdateMetadataRequest

	userID := c.GetInt("userID")

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(utils.NewErrorStruct(400, err.Error()))
		return
	}

	response, err := ctrl.userUseCase.UpdateMetadata(ctx, userID, &req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response)
}
