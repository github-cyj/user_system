package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		// 允许特定的请求头
		c.Header("Access-Control-Allow-Headers", "*")

		// 允许特定的来源请求
		c.Header("Access-Control-Allow-Origin", "*")

		// 允许特定的请求方法
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		// 允许请求头的键名称
		c.Header("Access-Control-Expose-Headers", "*")

		// 允许将对请求的响应暴露给页面
		c.Header("Access-Control-Allow-Credentials", "true")

		// 如果是预检请求（OPTIONS方法），则直接返回
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		// 继续处理其他请求
		c.Next()
	}
}
