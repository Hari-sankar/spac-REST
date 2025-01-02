package repository

import (
	"context"
	"spac-REST/api/schemas"
	"spac-REST/api/utils"
	"spac-REST/pkg/logger"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SpaceRepository interface {
	CreateSpace(ctx context.Context, space *schemas.CreateSpaceRequest, creatorID uuid.UUID) (uuid.UUID, error)
	DeleteSpace(ctx context.Context, spaceID, creatorID uuid.UUID) error
	GetAllSpaces(ctx context.Context, userID uuid.UUID) ([]schemas.SpaceResponse, error)
	GetSpace(ctx context.Context, spaceID uuid.UUID) (*schemas.GetSpaceResponse, error)
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

func (r *spaceRepository) DeleteSpace(ctx context.Context, spaceID uuid.UUID, creatorID uuid.UUID) error {
	query := `DELETE FROM "Space" WHERE id = $1 and "creatorId = $2"`

	commandTag, err := r.db.Exec(ctx, query, spaceID, creatorID)
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

//function to get a space and its element details

func (r *spaceRepository) GetSpace(ctx context.Context, spaceID uuid.UUID) (*schemas.GetSpaceResponse, error) {
	// First get space dimensions from map
	dimensionsQuery := `
        SELECT m.width || 'x' || m.height as dimensions
        FROM "Space" s
        JOIN "Map" m ON s.map_id = m.id
        WHERE s.id = $1`

	var response schemas.GetSpaceResponse
	err := r.db.QueryRow(ctx, dimensionsQuery, spaceID).Scan(&response.Dimensions)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, utils.NewErrorStruct(404, "Space not found")
		}
		return nil, utils.NewErrorStruct(500, "Failed to fetch space dimensions")
	}

	// Then get elements
	elementsQuery := `
        SELECT 
            me.id,
            e.id as element_id,
            e.image_url,
            e.static,
            e.height,
            e.width,
            me.x,
            me.y
        FROM "MapElements" me
        JOIN "Element" e ON me.element_id = e.id
        JOIN "Space" s ON s.map_id = me.map_id
        WHERE s.id = $1`

	rows, err := r.db.Query(ctx, elementsQuery, spaceID)
	if err != nil {
		return nil, utils.NewErrorStruct(500, "Failed to fetch space elements")
	}
	defer rows.Close()

	var elements []schemas.SpaceElement
	for rows.Next() {
		var element schemas.SpaceElement
		var elementID string
		err := rows.Scan(
			&element.ID,
			&elementID,
			&element.Element.ImageURL,
			&element.Element.Static,
			&element.Element.Height,
			&element.Element.Width,
			&element.X,
			&element.Y,
		)
		if err != nil {
			return nil, utils.NewErrorStruct(500, "Failed to scan element data")
		}
		element.Element.ID, _ = uuid.Parse(elementID)
		elements = append(elements, element)
	}

	response.Elements = elements
	return &response, nil
}
