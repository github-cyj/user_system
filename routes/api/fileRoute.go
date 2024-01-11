package api

import (
	"github.com/gin-gonic/gin"
	"user_system/app/controllers"
	"user_system/app/middleware"
)

func FileRoute(engine *gin.RouterGroup) {
	user := engine.Group("/file", middleware.AdminMiddleware())
	{
		user.POST("/upload_single", controllers.FileController{}.UploadSingle)
		user.POST("/upload_multi", controllers.FileController{}.UploadMulti)
	}
}
