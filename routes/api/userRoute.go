package api

import (
	"github.com/gin-gonic/gin"
	"user_system/app/controllers"
)

func UserRoute(engine *gin.RouterGroup) {
	user := engine.Group("/user")
	{
		user.GET("/:id", controllers.UserController{}.Get)
		user.GET("", controllers.UserController{}.List)
	}
}
