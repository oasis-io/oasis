package v1

import "github.com/gin-gonic/gin"

func GetUserGroupList(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "OK",
	})
}

func GetUserGroup(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "OK",
	})
}

func CreateUserGroup(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "OK",
	})
}

func UpdateUserGroup(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "OK",
	})
}

func DeleteUserGroup(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "OK",
	})
}
