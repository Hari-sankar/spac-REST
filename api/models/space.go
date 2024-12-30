package models

import "github.com/google/uuid"

// User represents the spaces table in the database
type Space struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Dimensions string    `json:"dimensions"`
	Thumbnail  string    `json:"thumbnail"`
}

// TableName specifies the table name for the User model
func (Space) TableName() string {
	return "spaces"
}
