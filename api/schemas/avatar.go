package schemas

import "github.com/google/uuid"

type GetAvatarsResponse struct {
	Avatars []AvatarResponse `json:"avatars"`
}

type AvatarResponse struct {
	ID       string `json:"id"`
	ImageURL string `json:"imageUrl"`
	Name     string `json:"name"`
}

type CreateAvatarRequest struct {
	ImageURL string `json:"imageUrl" binding:"required"`
	Name     string `json:"name" binding:"required"`
}

type CreateAvatarResponse struct {
	AvatarID uuid.UUID `json:"avatarId"`
}
