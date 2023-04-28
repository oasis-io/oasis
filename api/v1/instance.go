package v1

import "github.com/gin-gonic/gin"

func GetInstanceList(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "OK",
	})
}

func GetInstance(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "OK",
	})
}

func CreateInstance(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "OK",
	})
}

func UpdateInstance(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "OK",
	})
}

func DeleteInstance(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "OK",
	})
}

func CheckInstanceConn(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "OK",
	})
}
