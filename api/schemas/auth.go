package schemas

type SignUpRequest struct {
	Username string `json:"username" binding:"required,min=3,max=30"`
	Password string `json:"password" binding:"required,min=8,max=100"`
	Type     string `json:"type" binding:"required,oneof=Admin User"`
}

type SignUpResponse struct {
	UserID string `json:"userId"`
}

type SignInRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignInResponse struct {
	Token string `json:"token"`
}
