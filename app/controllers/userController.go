package controllers

import (
	"github.com/gin-gonic/gin"
	"user_system/app/constants"
	"user_system/app/request"
	"user_system/app/response"
)

type UserController struct {
}

func (controller UserController) List(c *gin.Context) {
	r := response.Gin{C: c}
	userListRequest := request.NewUserListRequest()
	err := c.ShouldBind(userListRequest)
	if err != nil {
		r.Error(constants.ErrorBind)
		return
	}
	r.Success(userListRequest)
}

func (controller UserController) Get(c *gin.Context) {
	r := response.Gin{C: c}
	r.Success(c.Param("id"))
}
