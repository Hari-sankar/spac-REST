package schemas

import (
	"github.com/google/uuid"
)

type CreateSpaceRequest struct {
	Name  string    `json:"name" binding:"required"`
	MapID uuid.UUID `json:"mapId" binding:"required"`
}

type CreateSpaceResponse struct {
	SpaceID uuid.UUID `json:"spaceId"`
}

type SpaceResponse struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Dimensions string    `json:"dimensions"`
	Thumbnail  string    `json:"Thumbnail"`
	CreatorID  uuid.UUID `json:"creatorID"`
}
type GetAllSpacesResponse struct {
	Spaces []SpaceResponse `json:"spaces"`
}
