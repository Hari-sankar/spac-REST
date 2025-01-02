package models

import "github.com/google/uuid"

type MapElement struct {
	ID       uuid.UUID `json:"id"`
	ImageURL string    `json:"imageUrl"`
	Static   bool      `json:"static"`
	Height   int       `json:"height"`
	Width    int       `json:"width"`
}

// TableName specifies the table name for the Element model
func (MapElement) TableName() string {
	return "MapElement"
}
