package usecases

import (
	"context"
	"spac-REST/api/repository"
	"spac-REST/api/schemas"

	"github.com/google/uuid"
)

type SpaceUseCase struct {
	spaceRepo repository.SpaceRepository
}

func NewSpaceUseCase(spaceRepo repository.SpaceRepository) *SpaceUseCase {
	return &SpaceUseCase{
		spaceRepo: spaceRepo,
	}
}

func (u *SpaceUseCase) CreateSpace(ctx context.Context, req *schemas.CreateSpaceRequest, creatorID uuid.UUID) (*schemas.CreateSpaceResponse, error) {
	spaceID, err := u.spaceRepo.CreateSpace(ctx, req, creatorID)
	if err != nil {
		return nil, err
	}

	return &schemas.CreateSpaceResponse{SpaceID: spaceID}, nil
}

func (u *SpaceUseCase) DeleteSpace(ctx context.Context, spaceID uuid.UUID) error {
	return u.spaceRepo.DeleteSpace(ctx, spaceID)
}

func (u *SpaceUseCase) GetAllSpaces(ctx context.Context, userID uuid.UUID) (*schemas.GetAllSpacesResponse, error) {
	spaces, err := u.spaceRepo.GetAllSpaces(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &schemas.GetAllSpacesResponse{Spaces: spaces}, nil
}
