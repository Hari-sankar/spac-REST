package usecases

import (
	"context"
	"spac-REST/api/models"
	"spac-REST/api/repository"
	"spac-REST/api/schemas"
	"spac-REST/api/utils"
	"time"
)

type AuthUseCase struct {
	authRepo repository.AuthRepository
}

func NewAuthUseCase(authRepo repository.AuthRepository) *AuthUseCase {
	return &AuthUseCase{authRepo: authRepo}
}

// Register handles user registration
func (auth *AuthUseCase) SignUp(ctx context.Context, req *schemas.SignUpRequest) (*schemas.SignUpResponse, error) {
	// Add registration logic, such as validation, business rules, etc.

	var user models.User
	var userID int
	var err error

	user.Username = req.Username
	user.Type = req.Type

	if user.Password, err = utils.HashPassword(req.Password); err != nil {
		return nil, err
	}

	if userID, err = auth.authRepo.CreateUser(ctx, &user); err != nil {
		return nil, err
	}
	return &schemas.SignUpResponse{
		UserID: userID,
	}, nil

}

// GetUserByID retrieves a user by their ID
func (auth *AuthUseCase) SignIn(ctx context.Context, username, password string) (*schemas.SignInResponse, error) {
	user, err := auth.authRepo.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, schemas.ErrInvalidCredentials
	}

	if err := utils.VerifyPassword(password, user.Password); err != nil {
		return nil, schemas.ErrInvalidCredentials
	}

	claims := &utils.Claims{
		UserID:    user.ID,
		UserName:  user.Username,
		ExpiresAt: time.Now().Add(30 * time.Minute),
	}

	token, err := utils.GenerateAcessToken(claims)
	if err != nil {
		return nil, schemas.ErrInternalServer
	}

	return &schemas.SignInResponse{
		Token: token,
	}, nil
}
