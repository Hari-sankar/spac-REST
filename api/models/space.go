package models

import "github.com/google/uuid"

// User represents the spaces table in the database
type Space struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	MapID     uuid.UUID `json:"mapID"`
	CreatorID uuid.UUID `json:"creatorID"`
}

// TableName specifies the table name for the User model
func (Space) TableName() string {
	return "spaces"
}
