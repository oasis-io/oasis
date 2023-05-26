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

		if username == "admin" {
			c.Next()
			return
		}

		// 从数据库中获取角色名或用户组名
		names, err := utils.GetRolesAndGroupsByUsername(username)
		if err != nil {
			response.Error(c, err.Error())
			return
		}

		if len(names) == 0 {
			response.Error(c, "no role found!")
			c.Abort()
			return
		}

		obj := c.Request.URL.Path // 将被访问的资源。
		act := c.Request.Method   // 用户对资源执行的操作。

		e := casbin.Casbin()
		authorized := false

		for _, name := range names {
			ok, err := e.Enforce(name, obj, act)
			if err != nil {
				response.Error(c, "Casbin Enforce error")
				c.Abort()
				return
			}
			if ok {
				authorized = true
				break
			}
		}

		if !authorized {
			response.Error(c, "permission denied")
			c.Abort()
			return
		}

		c.Next()
	}
}
