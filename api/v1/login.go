package v1

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"oasis/app/request"
	"oasis/app/response"
	"oasis/config"
	"oasis/db/model"
	"oasis/pkg/jwt"
	"oasis/pkg/log"
)

type LoginResponse struct {
	User  model.User `json:"user"`
	Token string     `json:"token"`
	//ExpirationTime int64      `json:"expiration"` // token到期时间
}

func Login(c *gin.Context) {
	var req request.Login

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, err.Error())
		return
	}

	db := config.DB

	user := model.User{}
	if err := db.Where("username = ? AND password = ?", req.Username, req.Password).Preload("Roles").First(&user).Error; err != nil {
		log.Error("获取用户角色错误", zap.Error(err))
		response.Error(c, err.Error())
		return
	}

	// 转换 user.Roles 为 []string 类型
	roleNames := make([]string, len(user.Roles))
	for i, role := range user.Roles {
		roleNames[i] = role.Name
	}

	claims := jwt.CustomClaims{
		Username: user.Username,
		Roles:    roleNames,
	}

	j := jwt.NewJWT()
	token, err := j.CreateToken(claims)
	if err != nil {
		response.Error(c, err.Error())
		return
	}

	response.SendSuccessData(c, "Login successful", LoginResponse{
		User:  user,
		Token: token,
	})
}
