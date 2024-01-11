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
		response.NewErrorResponseWithError(constants.Error, err).Return(c)
		return
	}
	userList, r := repository.UserRepository{}.GetList(userListRequest)
	r.ReturnWithData(c, userList)
}

func (controller UserController) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, r := repository.UserRepository{}.Get(uint(id))
	r.ReturnWithData(c, user)
}

func (controller UserController) Add(c *gin.Context) {
	userAddRequest := request.NewUserAddRequest()
	err := c.ShouldBind(&userAddRequest)
	if err != nil {
		response.NewErrorResponseWithError(constants.ErrorBind, err).Return(c)
		return
	}
	id, r := repository.UserRepository{}.Add(userAddRequest)
	r.ReturnWithData(c, id)

}

func (controller UserController) Edit(c *gin.Context) {
	userEditRequest := request.NewUserEditRequest()
	err := c.ShouldBind(&userEditRequest)
	if err != nil {
		response.NewErrorResponseWithError(constants.ErrorBind, err).Return(c)
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))
	updateCount, r := repository.UserRepository{}.Edit(uint(id), userEditRequest)
	r.ReturnWithData(c, updateCount)
}

func (controller UserController) Delete(c *gin.Context) {
	userEditRequest := request.NewUserEditRequest()
	err := c.ShouldBind(&userEditRequest)
	if err != nil {
		response.NewErrorResponseWithError(constants.ErrorBind, err).Return(c)
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))

	deleteCount, r := repository.UserRepository{}.Delete(uint(id))
	r.ReturnWithData(c, deleteCount)
}

func (controller UserController) UpdatePassword(c *gin.Context) {
	updatePasswordRequest := request.NewUpdatePasswordRequest()
	err := c.ShouldBind(&updatePasswordRequest)
	if err != nil {
		response.NewErrorResponseWithError(constants.ErrorBind, err).Return(c)
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))
	updateCount, r := repository.UserRepository{}.UpdatePassword(uint(id), updatePasswordRequest)
	r.ReturnWithData(c, updateCount)
}
