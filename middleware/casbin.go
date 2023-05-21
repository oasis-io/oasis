package middleware

import (
	"github.com/gin-gonic/gin"
	"oasis/app/response"
	"oasis/pkg/casbin"
	"oasis/pkg/utils"
)

func CasbinAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := utils.GetTokenUserName(c)
		sub := username           // 想要访问资源的用户。
		obj := c.Request.URL.Path // 将被访问的资源。
		act := c.Request.Method   // 用户对资源执行的操作。

		e := casbin.Casbin()
		ok, err := e.Enforce(sub, obj, act)
		if err != nil {
			response.Error(c, "Casbin Enforce error")
			c.Abort()
			return
		}

		if !ok {
			response.Error(c, "permission denied")
			c.Abort()
			return
		}

		c.Next()
	}
}
