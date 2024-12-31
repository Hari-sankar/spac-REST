package schemas

import (
	"spac-REST/api/models"

	"github.com/google/uuid"
)

type CreateSpaceRequest struct {
	Name  string    `json:"name" binding:"required"`
	MapID uuid.UUID `json:"mapId" binding:"required"`
}

type CreateSpaceResponse struct {
	SpaceID uuid.UUID `json:"spaceId"`
}

type GetAllSpacesResponse struct {
	Spaces []models.Space `json:"spaces"`
}
