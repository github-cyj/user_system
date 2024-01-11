package api

import (
	"github.com/gin-gonic/gin"
	"user_system/app/controllers"
	"user_system/app/middleware"
)

func AuthRoute(engine *gin.RouterGroup) {
	user := engine.Group("/auth")
	{
		user.POST("/login", controllers.AuthController{}.Login)
		user.GET("/info", middleware.AdminMiddleware(), controllers.AuthController{}.Info)
		user.DELETE("/logout", middleware.AdminMiddleware(), controllers.AuthController{}.Logout)
	}
}
