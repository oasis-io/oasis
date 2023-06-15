package utils

import (
	"github.com/gin-gonic/gin"
	"oasis/pkg/jwt"
)

func GetTokenUserName(c *gin.Context) (string, error) {
	// 解析token
	token := c.Request.Header.Get("x-token")

	j := jwt.NewJWT()
	x, err := j.ParseToken(token)
	if err != nil {
		return "", err
	}

	return x.Username, nil
}
