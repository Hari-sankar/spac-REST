package usecases

import (
	"context"
	"spac-REST/api/repository"
	"spac-REST/api/schemas"
)

type AvatarUseCase struct {
	avatarRepo repository.AvatarRepository
}

func NewAvatarUseCase(avatarRepo repository.AvatarRepository) *AvatarUseCase {
	return &AvatarUseCase{
		avatarRepo: avatarRepo,
	}
}

func (u *AvatarUseCase) GetAvailableAvatars(ctx context.Context) (*schemas.GetAvatarsResponse, error) {
	avatars, err := u.avatarRepo.GetAvailableAvatars(ctx)
	if err != nil {
		return nil, err
	}

	response := &schemas.GetAvatarsResponse{
		Avatars: make([]schemas.AvatarResponse, len(avatars)),
	}
	for i, avatar := range avatars {
		response.Avatars[i] = schemas.AvatarResponse{
			ID:       avatar.ID,
			ImageURL: avatar.ImageURL,
			Name:     avatar.Name,
		}
	}
	return response, nil
}

func (u *AvatarUseCase) CreateAvatar(ctx context.Context, req *schemas.CreateAvatarRequest) (*schemas.CreateAvatarResponse, error) {
	avatarID, err := u.avatarRepo.CreateAvatar(ctx, req)
	if err != nil {
		return nil, err
	}

	return &schemas.CreateAvatarResponse{AvatarID: avatarID}, nil
}
