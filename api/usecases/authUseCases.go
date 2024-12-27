package usecases

import (
	"context"
	"spac-REST/api/models"
	"spac-REST/api/repository"
	"spac-REST/api/schemas"
)

type AuthUseCase struct {
	authRepo repository.AuthRepository
}

func NewAuthUseCase(authRepo repository.AuthRepository) *AuthUseCase {
	return &AuthUseCase{authRepo: authRepo}
}

// Register handles user registration
func (auth *AuthUseCase) SignUp(ctx context.Context, req *schemas.SignUpRequest) error {
	// Add registration logic, such as validation, business rules, etc.

	var user models.User

	user.Password = req.Password
	user.Username = req.Username
	user.Type = req.Type
	return auth.authRepo.SignUp(ctx, &user)

}

// GetUserByID retrieves a user by their ID
func (auth *AuthUseCase) SignIn(ctx context.Context, id uint) (*models.User, error) {
	return auth.authRepo.SignIn(ctx, id)

}
