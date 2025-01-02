package controllers

import (
	"net/http"
	"spac-REST/api/schemas"
	"spac-REST/api/usecases"
	"spac-REST/api/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SpaceController struct {
	spaceUseCase *usecases.SpaceUseCase
}

func NewSpaceController(spaceUseCase *usecases.SpaceUseCase) *SpaceController {
	return &SpaceController{
		spaceUseCase: spaceUseCase,
	}
}

// CreateSpace godoc
// @Summary Create a new space
// @Description Create a space with specified name, dimensions and map ID
// @Tags spaces
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body schemas.CreateSpaceRequest true "Space creation request"
// @Success 200 {object} schemas.CreateSpaceResponse
// @Failure 400 {object} utils.ErrorStruct "Invalid request"
// @Failure 401 {object} utils.ErrorStruct "Unauthorized"
// @Failure 500 {object} utils.ErrorStruct "Internal Server Error"
// @Example {object} Request-Example
//
//	{
//	  "name": "My Space",
//	  "dimensions": "100x200",
//	  "mapId": "123e4567-e89b-12d3-a456-426614174000"
//	}
//
// @Example {object} Response-Example
//
//	{
//	  "spaceId": "123e4567-e89b-12d3-a456-426614174001"
//	}
//
// @Router /space [post]
func (ctrl *SpaceController) CreateSpace(c *gin.Context) {
	var req schemas.CreateSpaceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(utils.NewErrorStruct(400, err.Error()))
		return
	}

	userID, err := uuid.Parse(c.GetString("userID"))
	if err != nil {
		c.Error(utils.NewErrorStruct(401, "Invalid user ID"))
		return
	}

	response, err := ctrl.spaceUseCase.CreateSpace(c.Request.Context(), &req, userID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteSpace godoc
// @Summary Delete a space
// @Description Delete a space by its ID
// @Tags spaces
// @Produce json
// @Security BearerAuth
// @Param spaceId path string true "Space ID" Example(123e4567-e89b-12d3-a456-426614174000)
// @Success 200 "Successfully deleted"
// @Failure 400 {object} utils.ErrorStruct "Invalid space ID"
// @Failure 401 {object} utils.ErrorStruct "Unauthorized"
// @Failure 404 {object} utils.ErrorStruct "Space not found"
// @Failure 500 {object} utils.ErrorStruct "Internal Server Error"
// @Router /space/{spaceId} [delete]
func (ctrl *SpaceController) DeleteSpace(c *gin.Context) {
	spaceID, err := uuid.Parse(c.Param("spaceId"))
	if err != nil {
		c.Error(utils.NewErrorStruct(400, "Invalid space ID"))
		return
	}

	userID, err := uuid.Parse(c.GetString("userID"))
	if err != nil {
		c.Error(utils.NewErrorStruct(401, "Invalid user ID"))
		return
	}

	err = ctrl.spaceUseCase.DeleteSpace(c.Request.Context(), spaceID, userID)
	if err != nil {
		c.Error(err)
		return
	}

	c.Status(http.StatusOK)
}

// GetAllSpaces godoc
// @Summary Get all spaces
// @Description Get all spaces for the authenticated user
// @Tags spaces
// @Produce json
// @Security BearerAuth
// @Success 200 {object} schemas.GetAllSpacesResponse
// @Example {object} Response-Example
//
//	{
//	  "spaces": [
//	    {
//	      "id": "123e4567-e89b-12d3-a456-426614174000",
//	      "name": "My Space",
//	      "dimensions": "100x200",
//	      "thumbnail": "https://example.com/thumbnail.png"
//	    },
//	    {
//	      "id": "123e4567-e89b-12d3-a456-426614174001",
//	      "name": "Another Space",
//	      "dimensions": "300x400",
//	      "thumbnail": "https://example.com/another-thumbnail.png"
//	    }
//	  ]
//	}
//
// @Failure 401 {object} utils.ErrorStruct "Unauthorized"
// @Failure 500 {object} utils.ErrorStruct "Internal Server Error"
// @Router /space/all [get]
func (ctrl *SpaceController) GetAllSpaces(c *gin.Context) {
	userID, err := uuid.Parse(c.GetString("userID"))
	if err != nil {
		c.Error(utils.NewErrorStruct(401, "Invalid user ID"))
		return
	}

	response, err := ctrl.spaceUseCase.GetAllSpaces(c.Request.Context(), userID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response)
}
