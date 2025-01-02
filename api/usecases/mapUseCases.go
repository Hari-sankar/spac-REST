package usecases

import (
	"context"
	"spac-REST/api/models"
	"spac-REST/api/repository"
	"spac-REST/api/schemas"

	"github.com/google/uuid"
)

type MapUseCase struct {
	mapRepo repository.MapRepository
}

func NewMapUseCase(mapRepo repository.MapRepository) *MapUseCase {
	return &MapUseCase{
		mapRepo: mapRepo,
	}
}

func (u *MapUseCase) CreateMap(ctx context.Context, req *schemas.CreateMapRequest) (*schemas.CreateMapResponse, error) {
	mapID, err := u.mapRepo.CreateMap(ctx, req)
	if err != nil {
		return nil, err
	}

	return &schemas.CreateMapResponse{ID: mapID}, nil
}

func (u *MapUseCase) AddMapElements(ctx context.Context, mapID uuid.UUID, elements []models.MapElement) error {
	return u.mapRepo.AddMapElements(ctx, mapID, elements)
}

func (u *MapUseCase) GetAllMaps(ctx context.Context) (*schemas.GetAllMapsResponse, error) {
	maps, err := u.mapRepo.GetAllMaps(ctx)
	if err != nil {
		return nil, err
	}

	return &schemas.GetAllMapsResponse{Maps: maps}, nil
}
