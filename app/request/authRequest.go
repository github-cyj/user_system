package request

type AuthLoginRequest struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func NewAuthLoginRequest() *AuthLoginRequest {
	return &AuthLoginRequest{}
}
