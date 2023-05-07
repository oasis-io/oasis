package middleware

import (
	"github.com/gin-gonic/gin"
	"oasis/app/response"
	"oasis/pkg/utils"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("x-token")

		// 请求头没有token就提示
		if token == "" {
			response.Error(c, "No login")
			c.Abort()
		}

		// 解析token
		j := utils.NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			response.Error(c, err.Error())
			c.Abort()
		}

		c.Set("claims", claims)
		c.Next()
	}
}
