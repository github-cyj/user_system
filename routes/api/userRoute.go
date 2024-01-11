package api

import (
	"github.com/gin-gonic/gin"
	"user_system/app/controllers"
	"user_system/app/middleware"
)

func UserRoute(engine *gin.RouterGroup) {
	user := engine.Group("/user", middleware.AdminMiddleware())
	{
		user.GET("/:id", controllers.UserController{}.Get)
		user.GET("", controllers.UserController{}.PageList)
		user.POST("", controllers.UserController{}.Add)
		user.PUT("/:id", controllers.UserController{}.Edit)
		user.DELETE("/:id", controllers.UserController{}.Delete)
		user.PUT("/:id/password", controllers.UserController{}.UpdatePassword)
	}
}
