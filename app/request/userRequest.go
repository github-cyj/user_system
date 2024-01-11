package request

type UserListRequest struct {
	Page int `form:"page" json:"page"`
	Size int `form:"size" json:"size"`
}

type UserAddRequest struct {
	Username string `form:"username" json:"username" binding:"required"`
	Avatar   string `form:"avatar" json:"avatar"`
	Sex      uint16 `form:"sex" json:"sex" binding:"required"`
	Tel      string `form:"tel" json:"tel" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
}

type UserEditRequest struct {
	Username string `form:"username" json:"username"`
	Avatar   string `form:"avatar" json:"avatar"`
	Sex      uint16 `form:"sex" json:"sex" binding:"required"`
	Tel      string `form:"tel" json:"tel" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
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

type UpdatePasswordRequest struct {
	OldPassword    string `form:"old_password" json:"old_password" binding:"required"`
	NewPassword    string `form:"new_password" json:"new_password" binding:"required"`
	RepeatPassword string `form:"repeat_password" json:"repeat_password" binding:"required"`
}

func NewUpdatePasswordRequest() *UpdatePasswordRequest {
	return &UpdatePasswordRequest{}
}
