package utils

import (
	"spac-REST/pkg/logger"

	"golang.org/x/crypto/bcrypt"
)

// Function to hash password
func HashPassword(password string) (string, error) {
	if len(password) < 8 {
		return "", NewErrorStruct(400, "password must be at least 8 characters")
	}

	// Generate hash from password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", NewErrorStruct(500, "failed to hash password")
	}

	return string(hashedPassword), nil
}

// verify password provided by the user with the current user password
func VerifyPassword(password, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		logger.New().Error(err.Error())
		return NewErrorStruct(401, "invalid credentials")
	}
	return nil
}
