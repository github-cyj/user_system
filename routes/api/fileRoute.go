package api

import (
	"github.com/gin-gonic/gin"
	"user_system/app/controllers"
)

func FileRoute(engine *gin.RouterGroup) {
	user := engine.Group("/file")
	{
		user.POST("/upload", controllers.FileController{}.Upload)
	}
}
