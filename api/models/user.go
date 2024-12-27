package models

import (
	"time"
)

// User represents the users table in the database
type User struct {
	ID uint `gorm:"primaryKey" json:"id"`
	// Email             string    `gorm:"uniqueIndex;not null;type:varchar(255)" json:"email"`
	Password string  `gorm:"not null" json:"password"` // Password is omitted from JSON output
	Username string  `gorm:"not null;type:varchar(255)" json:"name"`
	Avatar   *string `gorm:"column:profile_picture_url" json:"profile_picture_url,omitempty"`
	// PhoneNumber       *string   `gorm:"type:varchar(20)" json:"phone_number,omitempty"`
	Type      string    `gorm:"type:varchar(50);default:USER" json:"type"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName specifies the table name for the User model
func (User) TableName() string {
	return "users"
}
