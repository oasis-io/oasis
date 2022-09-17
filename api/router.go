package api

import (
	"github.com/gin-gonic/gin"
	v1 "oasis/api/v1"
)

func Routers() {

	// gin 日志写入文件
	//gin.DisableConsoleColor() // 禁用控制台颜色
	//f, _ := os.Create("access.log")
	//gin.DefaultWriter = io.MultiWriter(f)

	r := gin.Default()

	// 健康检测
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})

	v1Router := r.Group("v1")
	{
		// instance
		v1Router.GET("/instances", v1.SelectInstance)
		v1Router.POST("/instances/add", v1.CreateInstance)
		v1Router.DELETE("/instances/:instance_name/", v1.DeleteInstance)
		v1Router.PATCH("/instances/:instance_name/", v1.UpdateInstance)
	}

	r.Run()
}
