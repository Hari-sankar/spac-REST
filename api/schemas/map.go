package schemas

import (
	"spac-REST/api/models"

	"github.com/google/uuid"
)

type CreateMapRequest struct {
	Thumbnail  string `json:"thumbnail" binding:"required"`
	Dimensions string `json:"dimensions" binding:"required"`
	Name       string `json:"name" binding:"required"`
}

type CreateMapResponse struct {
	ID uuid.UUID `json:"id"`
}

type AddMapElementRequest struct {
	Elements []models.MapElement `json:"elements" binding:"required"`
}

type MapResponse struct {
	ID         uuid.UUID `json:"id"`
	Thumbnail  string    `json:"thumbnail"`
	Dimensions string    `json:"dimensions"`
	Name       string    `json:"name"`
}

type GetAllMapsResponse struct {
	Maps []MapResponse `json:"maps"`
}
