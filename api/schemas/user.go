package schemas

type UpdateMetadataRequest struct {
	AvatarID string `json:"avatarId" binding:"required"`
}

type UpdateMetadataResponse struct {
	Success bool `json:"success"`
}

type GetUserMetadataBulkResponse struct {
	Avatars []UserMetadataResponse `json:"avatars"`
}

type UserMetadataResponse struct {
	UserID   string `json:"userId"`
	ImageURL string `json:"imageUrl"`
}
