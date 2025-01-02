package repository

import (
	"context"
	"spac-REST/api/models"
	"spac-REST/api/schemas"
	"spac-REST/api/utils"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MapRepository interface {
	CreateMap(ctx context.Context, req *schemas.CreateMapRequest) (uuid.UUID, error)
	AddMapElements(ctx context.Context, mapID uuid.UUID, elements []models.MapElement) error
	GetAllMaps(ctx context.Context) ([]schemas.MapResponse, error)
}

type mapRepository struct {
	db *pgxpool.Pool
}

func NewMapRepository(db *pgxpool.Pool) MapRepository {
	return &mapRepository{db: db}
}

func (r *mapRepository) CreateMap(ctx context.Context, req *schemas.CreateMapRequest) (uuid.UUID, error) {
	width, height, err := utils.ParseDimensions(req.Dimensions)
	if err != nil {
		return uuid.Nil, utils.NewErrorStruct(400, "Invalid dimensions format")
	}

	query := `
        INSERT INTO "Map" (name, width, height, thumbnail)
        VALUES ($1, $2, $3, $4)
        RETURNING id`

	var mapID uuid.UUID
	err = r.db.QueryRow(ctx, query, req.Name, width, height, req.Thumbnail).Scan(&mapID)
	if err != nil {
		return uuid.Nil, utils.NewErrorStruct(500, "Failed to create map")
	}

	return mapID, nil
}

func (r *mapRepository) AddMapElements(ctx context.Context, mapID uuid.UUID, elements []models.MapElement) error {
	query := `
        INSERT INTO "MapElements" (map_id, element_id, x, y)
        VALUES ($1, $2, $3, $4)`

	for _, element := range elements {
		_, err := r.db.Exec(ctx, query, mapID, element.ID, element.X, element.Y)
		if err != nil {
			return utils.NewErrorStruct(500, "Failed to add map elements")
		}
	}

	return nil
}

func (r *mapRepository) GetAllMaps(ctx context.Context) ([]schemas.MapResponse, error) {
	query := `
        SELECT id, thumbnail, width || 'x' || height as dimensions, name
        FROM "Map"
        ORDER BY name`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, utils.NewErrorStruct(500, "Failed to fetch maps")
	}
	defer rows.Close()

	var maps []schemas.MapResponse
	for rows.Next() {
		var mapResponse schemas.MapResponse
		err := rows.Scan(&mapResponse.ID, &mapResponse.Thumbnail, &mapResponse.Dimensions, &mapResponse.Name)
		if err != nil {
			return nil, utils.NewErrorStruct(500, "Failed to scan map data")
		}
		maps = append(maps, mapResponse)
	}

	return maps, nil
}
