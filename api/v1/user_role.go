package v1

import "github.com/gin-gonic/gin"

func GetRoleList(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "OK",
	})
}

func GetRole(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "OK",
	})
}

func CreateRole(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "OK",
	})
}

func UpdateRole(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "OK",
	})
}

func DeleteRole(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "OK",
	})
}
