package schemas

type UpdateMetadataRequest struct {
	AvatarID string `json:"avatarId" binding:"required"`
}

type UpdateMetadataResponse struct {
	Success bool `json:"success"`
}
