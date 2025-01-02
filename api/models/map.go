package models

import "github.com/google/uuid"

type Map struct {
	ID        uuid.UUID `json:"id"`
	Thumbnail string    `json:"thumbnail"`
	Name      string    `json:"name"`
	Height    int       `json:"height"`
	Width     int       `json:"width"`
}

// TableName specifies the table name for the Element model
func (Map) TableName() string {
	return "Map"
}
