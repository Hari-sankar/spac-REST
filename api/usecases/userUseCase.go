package usecases

import (
	"context"
	"spac-REST/api/repository"
	"spac-REST/api/schemas"
)

type UserUseCase struct {
	userRepo repository.UserRepository
}

func NewUserUseCase(userRepo repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepo: userRepo,
	}
}

func (u *UserUseCase) UpdateMetadata(ctx context.Context, userID int, req *schemas.UpdateMetadataRequest) (*schemas.UpdateMetadataResponse, error) {
	if err := u.userRepo.UpdateMetadata(ctx, userID, req.AvatarID); err != nil {
		return nil, err
	}

	return &schemas.UpdateMetadataResponse{Success: true}, nil
}

func (u *UserUseCase) GetUserMetadataBulk(ctx context.Context, userIDs []string) (*schemas.GetUserMetadataBulkResponse, error) {
	metadata, err := u.userRepo.GetUserMetadataBulk(ctx, userIDs)
	if err != nil {
		return nil, err
	}

	return &schemas.GetUserMetadataBulkResponse{
		Avatars: metadata,
	}, nil
}
