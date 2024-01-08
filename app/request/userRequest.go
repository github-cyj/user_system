package request

type UserListRequest struct {
	Page int `form:"page" json:"page"`
	Size int `form:"size" json:"size"`
}

type UserAddRequest struct {
	Username string `form:"username" json:"username" binding:"required"`
	Avatar   string `form:"avatar" json:"avatar"`
}

type UserEditRequest struct {
	Username string `form:"username" json:"username"`
	Avatar   string `form:"avatar" json:"avatar"`
}

func NewUserListRequest() *UserListRequest {
	return &UserListRequest{
		Page: 1,
		Size: 10,
	}
}

func NewUserAddRequest() *UserAddRequest {
	return &UserAddRequest{}
}

func NewUserEditRequest() *UserEditRequest {
	return &UserEditRequest{}
}
