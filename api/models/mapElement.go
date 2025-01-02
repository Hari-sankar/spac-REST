package models

import "github.com/google/uuid"

type MapElement struct {
	ID        uuid.UUID `json:"id"`
	MapID     uuid.UUID `json:"mapID"`
	ElementID uuid.UUID `json:"elementID"`
	X         int       `json:"X"`
	Y         int       `json:"Y"`
}

// TableName specifies the table name for the Element model
func (MapElement) TableName() string {
	return "MapElement"
}
