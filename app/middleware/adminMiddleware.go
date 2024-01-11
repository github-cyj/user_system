package middleware

import (
	"github.com/gin-gonic/gin"
	"user_system/app/cache"
	"user_system/app/constants"
	"user_system/app/models"
	"user_system/app/response"
)

var User = &models.User{}
var Token string

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		Token = c.GetHeader("X-Token")
		if Token == "" {
			response.NewErrorResponse(constants.InvalidToken).Return(c)
			c.Abort()
		}
		cache.RedisCache{}.Get(Token, User)
		c.Next()
	}
}
