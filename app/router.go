package app

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
	"net/http"
	v1 "oasis/api/v1"
	"oasis/config"
	"oasis/middleware"
	"oasis/pkg/log"
	"os"
)

func HttpRequests() {
	//gin.SetMode(gin.ReleaseMode)

	accessLog := config.NewOasisConfig().Server.LogAccess
	if accessLog == "on" {
		//gin.SetMode(gin.DebugMode)
		logPath := config.NewOasisConfig().Server.LogAccessPath
		f, _ := os.Create(logPath)
		// 将访问log写入文件
		//gin.DefaultWriter = io.MultiWriter(f)

		// 如果需要同时将日志写入文件和控制台，请使用以下代码。
		gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	}

	r := gin.Default()
	r.Use(middleware.Cors())

	// 将前端build后的dist文件夹复制到web文件夹
	// gin框架去哪里找static静态文件与index.html模板文件
	//r.Static("/assets", "./web/dist/assets")
	//r.LoadHTMLGlob("web/dist/*.html")
	//
	//r.GET("/", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "index.html", nil)
	//})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	r.POST("/v1/user/login", v1.Login)

	// 登陆后拿到token才能访问
	v1Router := r.Group("v1")
	v1Router.Use(middleware.JWTAuth()).Use(middleware.CasbinAuth())
	{
		// Menu
		v1Router.POST("/menu", v1.GetMenuTree)
		v1Router.POST("/menu/getBaseMenuTree", v1.GetBaseMenuTree)
		v1Router.POST("/menu/permissions", v1.MenuPermissions)
		v1Router.POST("/menu/api/permissions", v1.MenuApiPermissions)
		v1Router.POST("/menu/getBaseMenuApi", v1.GetBaseMenuApi)
		v1Router.POST("/menu/getMenuAuthorized", v1.GetMenuAuthorized)
		v1Router.POST("/menu/getMenuApiAuthorized", v1.GetMenuApiAuthorized)

		// SQL
		v1Router.POST("/sql/query", v1.QueryData)

		// Instance
		v1Router.POST("/instance", v1.GetInstance)
		v1Router.DELETE("/instance", v1.DeleteInstance)
		v1Router.PATCH("/instance", v1.UpdateInstance)
		v1Router.PATCH("/instance/password", v1.UpdateInstancePassword)
		v1Router.POST("/instance/add", v1.CreateInstance)
		v1Router.POST("/instance/list", v1.GetInstanceList)
		v1Router.POST("/instance/ping", v1.PingInstance)
		v1Router.GET("/instance/name", v1.GetInstanceName)
		v1Router.POST("/instance/database", v1.GetInstanceDatabase)

		// User List
		v1Router.POST("/user", v1.GetUser)
		v1Router.DELETE("/user", v1.DeleteUser)
		v1Router.PATCH("/user", v1.UpdateUser)
		v1Router.PATCH("/user/password", v1.UpdateUserPassword)
		v1Router.POST("/user/add", v1.CreateUser)
		v1Router.POST("/user/list", v1.GetUserList)
		v1Router.GET("/user", v1.GetUserInfo)
		v1Router.GET("/user/all", v1.GetUsers)

		// User Role
		v1Router.GET("/user/role", v1.GetRoles)
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

	bind := config.NewOasisConfig().Server.Bind
	port := config.NewOasisConfig().Server.Port
	address := bind + ":" + port

	log.Info("Start HTTP Listener", zap.String("address", address))
	r.Run(address)
}
