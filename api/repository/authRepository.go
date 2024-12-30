package repository

import (
	"context"
	"fmt"
	"spac-REST/api/models"
	"spac-REST/api/schemas"
	"spac-REST/api/utils"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthRepository interface {
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	CreateUser(ctx context.Context, user *models.User) (int, error)
}

type authRepository struct {
	db *pgxpool.Pool
}

func NewAuthRepository(db *pgxpool.Pool) AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) CreateUser(ctx context.Context, user *models.User) (int, error) {
	query := `INSERT INTO users (username, password, role)
			  VALUES (@Username, @Password, @Role)
			  RETURNING id`

	args := pgx.NamedArgs{
		"Username": user.Username,
		"Password": user.Password,
		"Role":     user.Type,
	}

	var id int
	err := r.db.QueryRow(ctx, query, args).Scan(&id)
	if err != nil {

		// Check for unique constraint violation
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr.Code == "23505" { // PostgreSQL unique violation code
				return 0, utils.NewErrorStruct(400, fmt.Sprintf("username %s already exists", user.Username))
			}
		}
		return 0, utils.NewErrorStruct(500, fmt.Sprintf("unable to create user: %v", err))
	}

	return id, nil
}

func (r *authRepository) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	query := `SELECT id, username, password, role FROM users WHERE username = $1`

	var user models.User
	err := r.db.QueryRow(ctx, query, username).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Type,
	)

	if err == pgx.ErrNoRows {
		return nil, schemas.ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}
