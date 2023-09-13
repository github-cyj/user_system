package request

type UserListRequest struct {
	Page int `form:"page" json:"page"`
	Size int `form:"size" json:"size"`
}

func NewUserListRequest() *UserListRequest {
	return &UserListRequest{
		Page: 1,
		Size: 10,
	}
}
