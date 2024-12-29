package repository

import (
	"context"
	"spac-REST/api/utils"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	UpdateMetadata(ctx context.Context, userID int, avatarID string) error
}

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) UpdateMetadata(ctx context.Context, userID int, avatarID string) error {
	query := `UPDATE users SET avatar_id = $1 WHERE id = $2`

	commandTag, err := r.db.Exec(ctx, query, avatarID, userID)
	if err != nil {
		return utils.NewErrorStruct(500, "Failed to update user metadata")
	}

	if commandTag.RowsAffected() == 0 {
		return utils.NewErrorStruct(404, "User not found")
	}

	return nil
}
