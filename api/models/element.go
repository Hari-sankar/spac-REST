package models

import "github.com/google/uuid"

type Element struct {
	ID       uuid.UUID `json:"id"`
	ImageURL string    `json:"imageUrl"`
	Static   bool      `json:"static"`
	Height   int       `json:"height"`
	Width    int       `json:"width"`
}

// TableName specifies the table name for the Element model
func (Element) TableName() string {
	return "Element"
}
