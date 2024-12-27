package usecases

import (
	"context"
	"spac-REST/api/models"
	"spac-REST/api/repository"
)

type UserUseCase struct {
	userRepo repository.UserRepository
}

func NewUserUseCase(userRepo repository.UserRepository) *UserUseCase {
	return &UserUseCase{userRepo: userRepo}
}

// Register handles user registration
func (u *UserUseCase) Register(ctx context.Context, user *models.User) error {
	// Add registration logic, such as validation, business rules, etc.
	return u.userRepo.CreateUser(ctx, user)

}

// GetUserByID retrieves a user by their ID
func (u *UserUseCase) GetUserByID(ctx context.Context, id uint) (*models.User, error) {
	return u.userRepo.GetUserByID(ctx, id)

}

// UpdateUser updates an existing user
func (u *UserUseCase) UpdateUser(ctx context.Context, user *models.User) error {
	// Add update logic, such as validation, business rules, etc.
	return u.userRepo.UpdateUser(ctx, user)

}

// DeleteUser removes a user by their ID
func (u *UserUseCase) DeleteUser(ctx context.Context, id uint) error {
	return u.userRepo.DeleteUser(ctx, id)
}

// GetAllUsers retrieves all users
func (u *UserUseCase) GetAllUsers(ctx context.Context) ([]models.User, error) {
	return u.userRepo.GetAllUsers(ctx)
}
