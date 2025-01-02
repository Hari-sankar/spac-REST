package controllers

import (
	"net/http"
	"spac-REST/api/schemas"
	"spac-REST/api/usecases"
	"spac-REST/api/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MapController struct {
	mapUseCase *usecases.MapUseCase
}

func NewMapController(mapUseCase *usecases.MapUseCase) *MapController {
	return &MapController{
		mapUseCase: mapUseCase,
	}
}

// CreateMap godoc
// @Summary Create a new map
// @Description Create a map with specified properties
// @Tags maps
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body schemas.CreateMapRequest true "Map creation request"
// @Success 200 {object} schemas.CreateMapResponse
// @Router /map [post]
func (ctrl *MapController) CreateMap(c *gin.Context) {
	var req schemas.CreateMapRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(utils.NewErrorStruct(400, err.Error()))
		return
	}

	response, err := ctrl.mapUseCase.CreateMap(c.Request.Context(), &req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// AddMapElements godoc
// @Summary Add elements to a map
// @Description Add multiple elements to an existing map
// @Tags maps
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param mapId path string true "Map ID"
// @Param request body schemas.AddMapElementRequest true "Map elements to add"
// @Success 200
// @Router /map/{mapId}/elements [post]
func (ctrl *MapController) AddMapElements(c *gin.Context) {
	mapID, err := uuid.Parse(c.Param("mapId"))
	if err != nil {
		c.Error(utils.NewErrorStruct(400, "Invalid map ID"))
		return
	}

	var req schemas.AddMapElementRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(utils.NewErrorStruct(400, err.Error()))
		return
	}

	err = ctrl.mapUseCase.AddMapElements(c.Request.Context(), mapID, req.Elements)
	if err != nil {
		c.Error(err)
		return
	}

	c.Status(http.StatusOK)
}

// GetAllMaps godoc
// @Summary Get all maps
// @Description Get a list of all available maps
// @Tags maps
// @Produce json
// @Success 200 {object} schemas.GetAllMapsResponse
// @Router /map/all [get]
func (ctrl *MapController) GetAllMaps(c *gin.Context) {
	response, err := ctrl.mapUseCase.GetAllMaps(c.Request.Context())
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response)
}
