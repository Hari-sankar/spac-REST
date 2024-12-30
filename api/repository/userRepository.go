package repository

import (
	"context"
	"fmt"
	"spac-REST/api/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, id uint) (*models.User, error)
	CreateUser(ctx context.Context, user *models.User) error
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, id uint) error
	GetAllUsers(ctx context.Context) ([]models.User, error)
}

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetUserByID(ctx context.Context, id uint) (*models.User, error) {
	var user models.User
	// if err := r.db.WithContext(ctx).First(&user, id).Error; err != nil {
	// 	return nil, err
	// }
	return &user, nil

}

func (r *userRepository) CreateUser(ctx context.Context, user *models.User) error {
	// return r.db.WithContext(ctx).Create(user).Error
	query := `INSERT INTO users (id, email, password, name, profile_picture_url, phone_number, role, created_at, updated_at)
	          VALUES (@ID, @Email, @Password, @Name, @ProfilePictureURL, @PhoneNumber, @Role, @CreatedAt, @UpdatedAt)`

	args := pgx.NamedArgs{
		"ID":                user.ID,
		"Password":          user.Password,
		"Name":              user.Username,
		"ProfilePictureURL": user.Avatar,
		"Role":              user.Type,
		"CreatedAt":         user.CreatedAt,
		"UpdatedAt":         user.UpdatedAt,
	}

	_, err := r.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}
func (r *userRepository) UpdateUser(ctx context.Context, user *models.User) error {
	query := `UPDATE users SET email = @Email, password = @Password, name = @Name, 
	          profile_picture_url = @ProfilePictureURL, phone_number = @PhoneNumber, 
	          role = @Role, updated_at = @UpdatedAt WHERE id = @ID`

	args := pgx.NamedArgs{
		"ID":                user.ID,
		"Password":          user.Password,
		"Name":              user.Username,
		"ProfilePictureURL": user.Avatar,
		"Role":              user.Type,
		"UpdatedAt":         user.UpdatedAt,
	}

	_, err := r.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to update user: %w", err)
	}

	return nil
}

func (r *userRepository) DeleteUser(ctx context.Context, id uint) error {
	query := `DELETE FROM users WHERE id = @ID`

	args := pgx.NamedArgs{
		"ID": id,
	}

	_, err := r.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to delete user: %w", err)
	}

	return nil
}

func (r *userRepository) GetAllUsers(ctx context.Context) ([]models.User, error) {
	query := `SELECT id, email, password, name, profile_picture_url, phone_number, role, created_at, updated_at FROM users`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve users: %w", err)
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(
			&user.ID, &user.Password, &user.Username,
			&user.Avatar, &user.Type,
			&user.CreatedAt, &user.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		users = append(users, user)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("row iteration error: %w", rows.Err())
	}

	return users, nil
}
