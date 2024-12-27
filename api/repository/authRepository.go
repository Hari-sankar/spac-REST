package repository

import (
	"context"
	"fmt"
	"spac-REST/api/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthRepository interface {
	SignIn(ctx context.Context, id uint) (*models.User, error)
	SignUp(ctx context.Context, user *models.User) error
}

type authRepository struct {
	db *pgxpool.Pool
}

func NewAuthRepository(db *pgxpool.Pool) AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) SignIn(ctx context.Context, id uint) (*models.User, error) {
	var user models.User
	// if err := r.db.WithContext(ctx).First(&user, id).Error; err != nil {
	// 	return nil, err
	// }
	return &user, nil

}

func (r *authRepository) SignUp(ctx context.Context, user *models.User) error {
	// return r.db.WithContext(ctx).Create(user).Error
	query := `INSERT INTO users (id, email, password, name, profile_picture_url, phone_number, role, created_at, updated_at)
	          VALUES (@ID, @Email, @Password, @Name, @ProfilePictureURL, @PhoneNumber, @Role, @CreatedAt, @UpdatedAt)`

	args := pgx.NamedArgs{
		"ID":                user.ID,
		"Email":             user.Email,
		"Password":          user.Password,
		"Name":              user.Name,
		"ProfilePictureURL": user.ProfilePictureURL,
		"PhoneNumber":       user.PhoneNumber,
		"Role":              user.Role,
		"CreatedAt":         user.CreatedAt,
		"UpdatedAt":         user.UpdatedAt,
	}

	_, err := r.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}
