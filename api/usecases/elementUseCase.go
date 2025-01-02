package usecases

import (
	"context"
	"spac-REST/api/repository"
	"spac-REST/api/schemas"

	"github.com/google/uuid"
)

type ElementUseCase struct {
	elementRepo repository.ElementRepository
}

func NewElementUseCase(elementRepo repository.ElementRepository) *ElementUseCase {
	return &ElementUseCase{
		elementRepo: elementRepo,
	}
}

func (u *ElementUseCase) CreateElement(ctx context.Context, req *schemas.CreateElementRequest) (*schemas.CreateElementResponse, error) {
	elementID, err := u.elementRepo.CreateElement(ctx, req)
	if err != nil {
		return nil, err
	}

	return &schemas.CreateElementResponse{ID: elementID}, nil
}

func (u *ElementUseCase) UpdateElement(ctx context.Context, elementID uuid.UUID, req *schemas.UpdateElementRequest) error {
	return u.elementRepo.UpdateElement(ctx, elementID, req)
}

func (u *ElementUseCase) GetAllElements(ctx context.Context) (*schemas.GetAllElementsResponse, error) {
	elements, err := u.elementRepo.GetAllElements(ctx)
	if err != nil {
		return nil, err
	}

	return &schemas.GetAllElementsResponse{Elements: elements}, nil
}
