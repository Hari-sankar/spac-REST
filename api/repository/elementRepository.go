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

type ElementRepository interface {
	CreateElement(ctx context.Context, req *schemas.CreateElementRequest) (uuid.UUID, error)
	UpdateElement(ctx context.Context, elementID uuid.UUID, req *schemas.UpdateElementRequest) error
	GetAllElements(ctx context.Context) ([]models.Element, error)
}

type elementRepository struct {
	db *pgxpool.Pool
}

func NewElementRepository(db *pgxpool.Pool) ElementRepository {
	return &elementRepository{db: db}
}

func (r *elementRepository) CreateElement(ctx context.Context, req *schemas.CreateElementRequest) (uuid.UUID, error) {
	query := `
        INSERT INTO "Element" ("imageUrl", width, height, static)
        VALUES ($1, $2, $3, $4)
        RETURNING id`

	var elementID uuid.UUID
	err := r.db.QueryRow(ctx, query, req.ImageURL, req.Width, req.Height, req.Static).Scan(&elementID)
	if err != nil {
		logger.New().Error(err)
		return uuid.Nil, utils.NewErrorStruct(500, "Failed to create element")
	}

	return elementID, nil
}

func (r *elementRepository) UpdateElement(ctx context.Context, elementID uuid.UUID, req *schemas.UpdateElementRequest) error {
	query := `
        UPDATE "Element"
        SET "imageUrl" = $1
        WHERE id = $2`

	result, err := r.db.Exec(ctx, query, req.ImageURL, elementID)
	if err != nil {
		logger.New().Error(err)
		return utils.NewErrorStruct(500, "Failed to update element")
	}

	if result.RowsAffected() == 0 {
		return utils.NewErrorStruct(404, "Element not found")
	}

	return nil
}

func (r *elementRepository) GetAllElements(ctx context.Context) ([]models.Element, error) {
	query := `
        SELECT id, image_url, width, height, static 
        FROM "Element"
        ORDER BY id`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, utils.NewErrorStruct(500, "Failed to fetch elements")
	}
	defer rows.Close()

	var elements []models.Element
	for rows.Next() {
		var element models.Element
		err := rows.Scan(
			&element.ID,
			&element.ImageURL,
			&element.Width,
			&element.Height,
			&element.Static,
		)
		if err != nil {
			return nil, utils.NewErrorStruct(500, "Failed to scan element data")
		}
		elements = append(elements, element)
	}

	return elements, nil
}
