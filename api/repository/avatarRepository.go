package repository

import (
	"context"
	"spac-REST/api/models"
	"spac-REST/api/schemas"
	"spac-REST/api/utils"
	"spac-REST/pkg/logger"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AvatarRepository interface {
	GetAvailableAvatars(ctx context.Context) ([]models.Avatar, error)
	CreateAvatar(ctx context.Context, req *schemas.CreateAvatarRequest) (uuid.UUID, error)
}

type avatarRepository struct {
	db *pgxpool.Pool
}

func NewAvatarRepository(db *pgxpool.Pool) AvatarRepository {
	return &avatarRepository{db: db}
}

func (r *avatarRepository) GetAvailableAvatars(ctx context.Context) ([]models.Avatar, error) {
	query := `SELECT id, image_url, name FROM "Avatar"`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		logger.New().Error(err)
		return nil, utils.NewErrorStruct(500, "Failed to fetch avatars")
	}
	defer rows.Close()

	var avatars []models.Avatar
	for rows.Next() {
		var avatar models.Avatar
		if err := rows.Scan(&avatar.ID, &avatar.ImageURL, &avatar.Name); err != nil {
			return nil, utils.NewErrorStruct(500, "Failed to scan avatar data")
		}
		avatars = append(avatars, avatar)
	}
	return avatars, nil
}

func (r *avatarRepository) CreateAvatar(ctx context.Context, req *schemas.CreateAvatarRequest) (uuid.UUID, error) {
	query := `
        INSERT INTO "Avatar" (image_url, name)
        VALUES ($1, $2)
        RETURNING id`

	var avatarID uuid.UUID
	err := r.db.QueryRow(ctx, query, req.ImageURL, req.Name).Scan(&avatarID)
	if err != nil {
		return uuid.Nil, utils.NewErrorStruct(500, "Failed to create avatar")
	}

	return avatarID, nil
}
