package controllers

import (
	"github.com/gin-gonic/gin"
	"user_system/app/constants"
	"user_system/app/repository"
	"user_system/app/request"
	"user_system/app/response"
)

type AuthController struct {
}

func (controller AuthController) Login(c *gin.Context) {
	authLoginRequest := request.NewAuthLoginRequest()
	err := c.ShouldBind(&authLoginRequest)
	if err != nil {
		response.NewErrorResponseWithError(constants.Error, err).Return(c)
		return
	}
	authLogin, r := repository.AuthRepository{}.Login(authLoginRequest)
	r.ReturnWithData(c, authLogin)
}

func (controller AuthController) Info(c *gin.Context) {
	userInfo, r := repository.AuthRepository{}.Info()
	r.ReturnWithData(c, userInfo)
}
