package v1

import "github.com/gin-gonic/gin"

func Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "OK",
	})
}

func GetUserList(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "OK",
	})
}

func GetUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "OK",
	})
}

func CreateUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "OK",
	})
}

func UpdateUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "OK",
	})
}

func DeleteUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "OK",
	})
}
