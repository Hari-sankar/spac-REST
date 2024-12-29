package models

type Avatar struct {
	ID       string `db:"id" json:"id"`
	ImageURL string `db:"image_url" json:"imageUrl"`
	Name     string `db:"name" json:"name"`
}

func (Avatar) TableName() string {
	return "avatars"
}
