package utils

import (
	"github.com/gin-gonic/gin"
	"oasis/app/response"
	"oasis/pkg/jwt"
)

func GetTokenUserName(c *gin.Context) string {
	// 解析token
	token := c.Request.Header.Get("x-token")

	j := jwt.NewJWT()
	x, err := j.ParseToken(token)
	if err != nil {
		response.Error(c, err.Error())
		c.Abort()
	}

	name := x.Username

	return name
}
