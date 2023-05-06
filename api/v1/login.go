package v1

import (
	"github.com/gin-gonic/gin"
	"oasis/app/request"
	"oasis/app/response"
	"oasis/config"
	"oasis/db/model"
	"oasis/pkg/utils"
)

type LoginResponse struct {
	User  model.User `json:"user"`
	Token string     `json:"token"`
	//ExpirationTime int64      `json:"expiration"` // token到期时间
}

func Login(c *gin.Context) {
	var req request.Login

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, response.ParamsError, err.Error())
		return
	}

	db := config.DB

	user := model.User{}
	err := db.Where("username = ? AND password = ?", req.Username, req.Password).Find(&user).Error
	if err != nil {
		response.Error(c, response.SQLError, err.Error())
	}

	j := utils.NewJWT()
	token, err := j.CreateToken(req.Username)
	if err != nil {
		response.Error(c, response.TokenError, err.Error())
	}

	response.Success(c, "登录成功", LoginResponse{
		User:  user,
		Token: token,
	})
}
