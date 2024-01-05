package repository

import (
	"user_system/app/constants"
	"user_system/app/request"
	"user_system/app/response"
	"user_system/utils"
)

type AuthRepository struct {
}

func (repository AuthRepository) Login(params *request.AuthLoginRequest) (authLogin response.AuthLoginResponse, r *response.Response) {
	if params.Username != "admin" {
		r = response.NewErrorResponse(constants.ErrorUsername)
		return authLogin, r
	}

	if params.Password != "123456" {
		r = response.NewErrorResponse(constants.ErrorPassword)
		return authLogin, r
	}
	authLogin.Token = utils.MD5("123456")
	return authLogin, r
}

func (repository AuthRepository) Info() (userInfo response.UserInfoResponse, r *response.Response) {
	userInfo = response.UserInfoResponse{
		Roles:        []string{"admin"},
		Introduction: "I am a super administrator",
		Avatar:       "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
		Name:         "Super Admin",
	}
	return userInfo, r
}
