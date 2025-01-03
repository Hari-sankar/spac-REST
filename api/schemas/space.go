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

type SpaceElement struct {
	ID      int            `json:"id"`
	Element models.Element `json:"element"`
	X       int            `json:"x"`
	Y       int            `json:"y"`
}

type GetSpaceResponse struct {
	Dimensions string         `json:"dimensions"`
	Elements   []SpaceElement `json:"elements"`
}

type SpaceDimension struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type SpaceValidationResponse struct {
	Valid    bool      `json:"valid"`
	Width    int       `json:"width"`
	Height   int       `json:"height"`
	UserID   uuid.UUID `json:"userId"`
	Username string    `json:"username"`
}
