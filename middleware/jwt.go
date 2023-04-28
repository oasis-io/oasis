package middleware

import (
	"github.com/gin-gonic/gin"
	"oasis/pkg/utils"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("x-token")

		// 请求头没有token就提示
		if token == "" {
			c.JSON(400, gin.H{
				"message": "not login!",
			})
			c.Abort()
		}

		// 解析token
		j := utils.NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			c.JSON(405, gin.H{
				"message": "token expired！",
			})
			c.Abort()
		}

		c.Set("claims", claims)
		c.Next()
	}
}
