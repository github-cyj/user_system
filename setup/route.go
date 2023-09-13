package setup

import (
	"github.com/gin-gonic/gin"
	"user_system/app/middleware"
	"user_system/routes/api"
)

func Route() *gin.Engine {
	r := gin.Default()
	// 跨域处理
	r.Use(middleware.CORSMiddleware())
	getApi(r)
	return r
}

func getApi(engine *gin.Engine) {
	apiRoute := engine.Group("api")
	{
		api.UserRoute(apiRoute)
		api.FileRoute(apiRoute)
	}
}
