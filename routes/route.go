package routes

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"user_system/app/constants"
	"user_system/app/middleware"
	"user_system/app/response"
	"user_system/routes/api"
)

func Route() *gin.Engine {
	r := gin.Default()
	// 跨域处理
	r.Use(middleware.CORSMiddleware())
	// 异常处理
	r.Use(gin.Recovery())
	r.Use(gin.RecoveryWithWriter(os.Stderr, func(c *gin.Context, err any) {
		response.NewErrorResponseWithError(
			constants.Error,
			errors.New(fmt.Sprintf("%v", err)),
		).ReturnWithHttpCode(c, http.StatusInternalServerError)
	}))
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
