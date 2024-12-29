package controllers

import (
	"net/http"
	"spac-REST/api/schemas"
	"spac-REST/api/usecases"
	"spac-REST/api/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

// GetUserMetadataBulk godoc
// @Summary Get users metadata bulk
// @Description Get metadata for multiple users
// @Tags users
// @Produce json
// @Security BearerAuth
// @Param ids query string true "User IDs array format: [uuid1,uuid2,uuid3]"
// @Success 200 {object} schemas.GetUserMetadataBulkResponse
// @Failure 400 {object} utils.ErrorStruct "Bad Request"
// @Failure 401 {object} utils.ErrorStruct "Unauthorized"
// @Failure 500 {object} utils.ErrorStruct "Internal Server Error"
// @Router /users/metadata/bulk [get]
func (ctrl *UserController) GetUserMetadataBulk(c *gin.Context) {
	idsStr := c.Query("ids")
	// Remove brackets
	idsStr = strings.Trim(idsStr, "[]")

	// Split by comma
	userIDs := strings.Split(idsStr, ",")

	// Validate each ID
	for _, id := range userIDs {
		// Check if ID is valid UUID
		if _, err := uuid.Parse(strings.TrimSpace(id)); err != nil {
			c.Error(utils.NewErrorStruct(400, "Invalid ID format. All IDs must be valid UUIDs"))
			return
		}
	}

	response, err := ctrl.userUseCase.GetUserMetadataBulk(c.Request.Context(), userIDs)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response)
}
