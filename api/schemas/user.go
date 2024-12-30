package schemas

type UpdateMetadataRequest struct {
	AvatarID string `json:"avatarId" binding:"required"`
}

type UpdateMetadataResponse struct {
	Success bool `json:"success"`
}

type GetAvatarsResponse struct {
	Avatars []AvatarResponse `json:"avatars"`
}

type AvatarResponse struct {
	ID       string `json:"id"`
	ImageURL string `json:"imageUrl"`
	Name     string `json:"name"`
}

type GetUserMetadataBulkResponse struct {
	Avatars []UserMetadataResponse `json:"avatars"`
}

type UserMetadataResponse struct {
	UserID   string `json:"userId"`
	ImageURL string `json:"imageUrl"`
}
