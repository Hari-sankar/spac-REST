package repository

import (
	"context"
	"strconv"

	"spac-REST/api/schemas"
	"spac-REST/api/utils"
	"spac-REST/pkg/logger"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	UpdateMetadata(ctx context.Context, userID int, avatarID string) error
	GetUserMetadataBulk(ctx context.Context, userIDs []string) ([]schemas.UserMetadataResponse, error)
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
	query := `UPDATE "User" SET avatarId = $1 WHERE id = $2`

	commandTag, err := r.db.Exec(ctx, query, avatarID, userID)
	if err != nil {
		return utils.NewErrorStruct(500, "Failed to update user metadata")
	}

	if commandTag.RowsAffected() == 0 {
		return utils.NewErrorStruct(404, "User not found")
	}

	return nil
}

func (r *userRepository) GetUserMetadataBulk(ctx context.Context, userIDs []string) ([]schemas.UserMetadataResponse, error) {
	query := `
		SELECT u.id, a.imageUrl 
		FROM "User" u 
		LEFT JOIN "Avatar" a ON u.avatar_id = a.id 
		WHERE u.id = ANY($1::integer[])
	`

	var intIDs []int
	for _, id := range userIDs {
		intID, err := strconv.Atoi(id)
		if err != nil {
			logger.New().Error(err)
			return nil, utils.NewErrorStruct(400, "Invalid user ID format")
		}
		intIDs = append(intIDs, intID)
	}

	rows, err := r.db.Query(ctx, query, intIDs)
	if err != nil {
		logger.New().Error(err)
		return nil, utils.NewErrorStruct(500, "Failed to fetch user metadata")
	}
	defer rows.Close()

	var metadata []schemas.UserMetadataResponse
	for rows.Next() {
		var response schemas.UserMetadataResponse
		var id int
		if err := rows.Scan(&id, &response.ImageURL); err != nil {
			logger.New().Error(err)
			return nil, utils.NewErrorStruct(500, "Failed to scan user metadata")
		}
		response.UserID = strconv.Itoa(id)
		metadata = append(metadata, response)
	}
	return metadata, nil
}
