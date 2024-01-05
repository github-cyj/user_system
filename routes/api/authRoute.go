package api

import (
	"github.com/gin-gonic/gin"
	"user_system/app/controllers"
)

func AuthRoute(engine *gin.RouterGroup) {
	user := engine.Group("/auth")
	{
		user.POST("/login", controllers.AuthController{}.Login)
		user.GET("/info", controllers.AuthController{}.Info)
	}
}
