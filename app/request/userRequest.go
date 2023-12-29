package request

type UserListRequest struct {
	Page int `form:"page" json:"page"`
	Size int `form:"size" json:"size"`
}

type UserAddRequest struct {
	Username string `form:"username" json:"username" binding:"required"`
}

type UserEditRequest struct {
	Username string `form:"username" json:"username"`
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
