package schemas

import (
	"spac-REST/api/models"

	"github.com/google/uuid"
)

type CreateElementRequest struct {
	ImageURL string `json:"imageURL" binding:"required"`
	Width    int    `json:"width" binding:"required"`
	Height   int    `json:"height" binding:"required"`
	Static   bool   `json:"static" binding:"required"`
}

type CreateElementResponse struct {
	ID uuid.UUID `json:"id"`
}

type UpdateElementRequest struct {
	ImageURL string `json:"imageUrl" binding:"required"`
}

type GetAllElementsResponse struct {
	Elements []models.Element `json:"elements"`
}
