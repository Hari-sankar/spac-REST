package controllers

import (
	"net/http"
	"spac-REST/api/schemas"
	"spac-REST/api/usecases"
	"spac-REST/api/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ElementController struct {
	elementUseCase *usecases.ElementUseCase
}

func NewElementController(elementUseCase *usecases.ElementUseCase) *ElementController {
	return &ElementController{
		elementUseCase: elementUseCase,
	}
}

// CreateElement godoc
// @Summary Create a new element
// @Description Create an element with specified properties
// @Tags elements
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body schemas.CreateElementRequest true "Element creation request"
// @Success 200 {object} schemas.CreateElementResponse
// @Failure 400 {object} utils.ErrorStruct "Invalid request"
// @Failure 401 {object} utils.ErrorStruct "Unauthorized"
// @Failure 500 {object} utils.ErrorStruct "Internal Server Error"
// @Example {object} Request-Example
//
//	{
//	  "imageUrl": "https://example.com/image.png",
//	  "width": 1,
//	  "height": 1,
//	  "static": true
//	}
//
// @Router /element [post]
func (ctrl *ElementController) CreateElement(c *gin.Context) {
	var req schemas.CreateElementRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(utils.NewErrorStruct(400, err.Error()))
		return
	}

	response, err := ctrl.elementUseCase.CreateElement(c.Request.Context(), &req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// UpdateElement godoc
// @Summary Update an element
// @Description Update an element's image URL
// @Tags elements
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param elementId path string true "Element ID"
// @Param request body schemas.UpdateElementRequest true "Element update request"
// @Success 200 "Successfully updated"
// @Failure 400 {object} utils.ErrorStruct "Invalid request"
// @Failure 401 {object} utils.ErrorStruct "Unauthorized"
// @Failure 404 {object} utils.ErrorStruct "Element not found"
// @Failure 500 {object} utils.ErrorStruct "Internal Server Error"
// @Example {object} Request-Example
//
//	{
//	  "imageUrl": "https://example.com/new-image.png"
//	}
//
// @Router /element/{elementId} [put]
func (ctrl *ElementController) UpdateElement(c *gin.Context) {
	elementID, err := uuid.Parse(c.Param("elementId"))
	if err != nil {
		c.Error(utils.NewErrorStruct(400, "Invalid element ID"))
		return
	}

	var req schemas.UpdateElementRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(utils.NewErrorStruct(400, err.Error()))
		return
	}

	err = ctrl.elementUseCase.UpdateElement(c.Request.Context(), elementID, &req)
	if err != nil {
		c.Error(err)
		return
	}

	c.Status(http.StatusOK)
}

// GetAllElements godoc
// @Summary Get all elements
// @Description Get a list of all available elements
// @Tags elements
// @Produce json
// @Security BearerAuth
// @Success 200 {object} schemas.GetAllElementsResponse
// @Failure 500 {object} utils.ErrorStruct "Internal Server Error"
// @Router /element/all [get]
func (ctrl *ElementController) GetAllElements(c *gin.Context) {
	response, err := ctrl.elementUseCase.GetAllElements(c.Request.Context())
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response)
}
