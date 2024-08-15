package repository

import (
	"fmt"
	"time"
	"user_system/app/baseType"
	"user_system/app/cache"
	"user_system/app/constants"
	"user_system/app/middleware"
	"user_system/app/models"
	"user_system/app/request"
	"user_system/app/response"
	"user_system/utils"
)

type AuthRepository struct {
}

func (repository AuthRepository) Login(params *request.AuthLoginRequest) (authLogin response.AuthLoginResponse, r *response.Response) {
	var user *models.User
	if params.Username != "admin" {
		var password = utils.MD5(params.Password)
		result := models.Db.Where("( tel = ? or email= ? ) and password = ?", params.Username, params.Username, password).First(&user)
		if result.RowsAffected == 0 {
			r = response.NewErrorResponse(constants.ErrorUsernameOrPassword)
			return authLogin, r
		}
	} else {
		var now = baseType.JsonTime(time.Now())
		if params.Password != "123456" {
			r = response.NewErrorResponse(constants.ErrorUsernameOrPassword)
			return authLogin, r
		}
		user = &models.User{
			Username: params.Username,
			Avatar:   "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
			Tel:      "10000000000",
			Sex:      1,
			Email:    "10000000000@163.com",
			BaseModel: models.BaseModel{
				CreateTime: now,
				UpdateTime: now,
			},
		}
	}
	authLogin.Token = utils.MD5(user.Tel)
	fmt.Println(user)
	cache.RedisCache{}.Set(authLogin.Token, user, 0)
	return authLogin, r
}

func (repository AuthRepository) Info() (userInfo response.UserInfoResponse, r *response.Response) {
	var avatar = middleware.User.Avatar
	if avatar == "" {
		avatar = "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif"
	}
	userInfo = response.UserInfoResponse{
		Roles:        []string{"admin"},
		Introduction: "I am a super administrator",
		Avatar:       avatar,
		Name:         middleware.User.Username,
	}
	return userInfo, r
}

func (repository AuthRepository) Logout() (r *response.Response) {
	r = response.NewSuccessResponse(
		cache.RedisCache{}.Del(middleware.Token),
	)
	return r
}
