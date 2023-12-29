package controllers

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"user_system/app/constants"
	"user_system/app/repository"
	"user_system/app/request"
	"user_system/app/response"
)

type UserController struct {
}

func (controller UserController) PageList(c *gin.Context) {
	userListRequest := request.NewUserListRequest()
	err := c.ShouldBind(&userListRequest)
	if err != nil {
		response.NewError(constants.Error, err).Error(c)
		return
	}
	result := repository.UserRepository{}.GetList(userListRequest)
	response.NewSuccess(result).Success(c)
}

func (controller UserController) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	result := repository.UserRepository{}.Get(uint(id))
	if result.ID == 0 {
		response.NewErrorWithData(constants.ErrorNotExits, "用户").Error(c)
		return
	}
	response.NewSuccess(result).Success(c)
}

func (controller UserController) Add(c *gin.Context) {
	userAddRequest := request.NewUserAddRequest()
	err := c.ShouldBind(&userAddRequest)
	if err != nil {
		response.NewError(constants.ErrorBind, err).Error(c)
		return
	}
	result := repository.UserRepository{}.Add(userAddRequest)
	response.NewSuccess(result).Success(c)

}

func (controller UserController) Edit(c *gin.Context) {
	userEditRequest := request.NewUserEditRequest()
	err := c.ShouldBind(&userEditRequest)
	if err != nil {
		response.NewError(constants.ErrorBind, err).Error(c)
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))
	userInfo := repository.UserRepository{}.Get(uint(id))
	if userInfo.ID == 0 {
		response.NewErrorWithData(constants.ErrorNotExits, "用户").Error(c)
		return
	}
	result := repository.UserRepository{}.Edit(uint(id), userEditRequest)
	response.NewSuccess(result).Success(c)
}

func (controller UserController) Delete(c *gin.Context) {
	userEditRequest := request.NewUserEditRequest()
	err := c.ShouldBind(&userEditRequest)
	if err != nil {
		response.NewError(constants.ErrorBind, err).Error(c)
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))
	userInfo := repository.UserRepository{}.Get(uint(id))
	if userInfo.ID == 0 {
		response.NewErrorWithData(constants.ErrorNotExits, "用户").Error(c)
		return
	}
	result := repository.UserRepository{}.Delete(uint(id))
	response.NewSuccess(result).Success(c)
}
