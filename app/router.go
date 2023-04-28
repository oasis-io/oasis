package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
	v1 "oasis/api/v1"
	"oasis/middleware"
	"oasis/pkg/logs"
)

func HttpRequests() {
	r := gin.Default()
	r.Use(middleware.Cors())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "alive",
		})
	})

	r.POST("/v1/user/login", v1.Login)

	// 登陆后拿到token才能访问
	v1Router := r.Group("v1")
	v1Router.Use(middleware.JWTAuth())
	{
		// Instance
		v1Router.POST("/instance", v1.GetInstance)
		v1Router.DELETE("/instance", v1.DeleteInstance)
		v1Router.PATCH("/instance", v1.UpdateInstance)
		v1Router.POST("/instance/add", v1.CreateInstance)
		v1Router.POST("/instance/list", v1.GetInstanceList)
		v1Router.POST("/instance/ping", v1.CheckInstanceConn)

		// User List
		v1Router.POST("/user", v1.GetUser)
		v1Router.DELETE("/user", v1.DeleteUser)
		v1Router.PATCH("/user", v1.UpdateUser)
		v1Router.POST("/user/add", v1.CreateUser)
		v1Router.POST("/user/list", v1.GetUserList)
		//v1Router.GET("/user/info", v1.GetUserInfo)

		// User Role
		v1Router.POST("/user/role", v1.GetRole)
		v1Router.DELETE("/user/role", v1.DeleteRole)
		v1Router.PATCH("/user/role", v1.UpdateRole)
		v1Router.POST("/user/role/add", v1.CreateRole)
		v1Router.POST("/user/role/list", v1.GetRoleList)

		// User Group
		v1Router.POST("/user/group", v1.GetUserGroup)
		v1Router.DELETE("/user/group", v1.DeleteUserGroup)
		v1Router.PATCH("/user/group", v1.UpdateUserGroup)
		v1Router.POST("/user/group/add", v1.CreateUserGroup)
		v1Router.POST("/user/group/list", v1.GetUserGroupList)

	}

	logs.Logger().Info("Run Oasis Server, Port: 9590")
	r.Run(":9590")
}
