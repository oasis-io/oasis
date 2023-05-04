package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"oasis/app/request"
	"oasis/app/response"
	"oasis/config"
	"oasis/db/model"
	"oasis/pkg/utils"
)

type LoginResponse struct {
	//User  model.User `json:"user"`
	Token string `json:"token"`
	//ExpirationTime int64      `json:"expiration"` // token到期时间
}

func Login(c *gin.Context) {
	var req request.Login
	var user []model.User

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, response.ParamsError, err.Error())
		return
	}

	db := config.DB

	username := req.Username
	password := req.Password

	db.Where("username = ? AND password = ?", username, password).Find(&user)

	j := utils.NewJWT()
	token, err := j.CreateToken(username)
	if err != nil {
		fmt.Errorf("token create fail error: %v", err)
	}

	response.Success(c, "登录成功", LoginResponse{
		Token: token,
	})
}
