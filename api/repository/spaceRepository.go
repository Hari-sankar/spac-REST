package repository

import (
	"context"
	"spac-REST/api/schemas"
	"spac-REST/api/utils"
	"spac-REST/pkg/logger"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SpaceRepository interface {
	CreateSpace(ctx context.Context, space *schemas.CreateSpaceRequest, creatorID uuid.UUID) (uuid.UUID, error)
	DeleteSpace(ctx context.Context, spaceID uuid.UUID) error
	GetAllSpaces(ctx context.Context, userID uuid.UUID) ([]schemas.SpaceResponse, error)
}

type spaceRepository struct {
	db *pgxpool.Pool
}

func NewSpaceRepository(db *pgxpool.Pool) SpaceRepository {
	return &spaceRepository{db: db}
}

func (r *spaceRepository) CreateSpace(ctx context.Context, space *schemas.CreateSpaceRequest, creatorID uuid.UUID) (uuid.UUID, error) {
	query := `
        INSERT INTO "Space" (name, "mapId", "creatorId")
        VALUES ($1, $2, $3)
        RETURNING id`

	var spaceID uuid.UUID
	err := r.db.QueryRow(ctx, query, space.Name, space.MapID, creatorID).Scan(&spaceID)
	if err != nil {
		logger.New().Error(err.Error())
		return uuid.Nil, utils.NewErrorStruct(500, "Failed to create space")
	}

	return spaceID, nil
}

func (r *spaceRepository) DeleteSpace(ctx context.Context, spaceID uuid.UUID) error {
	query := `DELETE FROM "Space" WHERE id = $1`

	commandTag, err := r.db.Exec(ctx, query, spaceID)
	if err != nil {
		return utils.NewErrorStruct(500, "Failed to delete space")
	}

	if commandTag.RowsAffected() == 0 {
		return utils.NewErrorStruct(404, "Space not found")
	}

	return nil
}

func (r *spaceRepository) GetAllSpaces(ctx context.Context, userID uuid.UUID) ([]schemas.SpaceResponse, error) {
	query := `
        SELECT s.id, s.name, m.width || 'x' || m.height as dimensions, m.thumbnail
        FROM "Space"  s JOIN "Map"  m ON (s."mapId" = m.id)  
        WHERE "creatorId" = $1`

	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		logger.New().Error(err.Error())
		return nil, utils.NewErrorStruct(500, "Failed to fetch spaces")
	}
	defer rows.Close()
	var spaces []schemas.SpaceResponse

	for rows.Next() {
		var space schemas.SpaceResponse
		if err := rows.Scan(&space.ID, &space.Name, &space.Dimensions, &space.Thumbnail); err != nil {
			return nil, utils.NewErrorStruct(500, "Failed to scan space data")
		}
		spaces = append(spaces, space)
	}

	return spaces, nil
}
