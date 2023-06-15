package v1

import (
	"github.com/gin-gonic/gin"
	"oasis/app/request"
	"oasis/app/response"
	"oasis/db"
	"oasis/db/model"
	"oasis/pkg/jwt"
	"oasis/pkg/log"
)

type LoginResponse struct {
	User  model.User `json:"user"`
	Token string     `json:"token"`
	//ExpirationTime int64      `json:"expiration"` // token到期时间
}

// Login 登陆帐号与密码验证并且返回token
func Login(c *gin.Context) {
	var req request.Login
	var claims jwt.CustomClaims

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("parameter binding errors: " + err.Error())
		response.Error(c, "parameter binding errors")
		return
	}

	// 数据库判断数据里是否一致
	user, err := db.Login(req.Username, req.Password)
	if err != nil {
		response.Error(c, "username or password error")
		return
	}

	if req.Username == "admin" {
		claims = jwt.CustomClaims{
			Username: user.Username,
		}
	} else {
		// 转换 user.Roles 为 []string 类型
		roleNames := make([]string, len(user.Roles))
		for i, role := range user.Roles {
			roleNames[i] = role.Name
		}

		claims = jwt.CustomClaims{
			Username: user.Username,
			Roles:    roleNames,
		}
	}

	j := jwt.NewJWT()
	token, err := j.CreateToken(claims)
	if err != nil {
		response.Error(c, err.Error())
		return
	}

	response.SendSuccessData(c, "Login successful", LoginResponse{
		User:  *user,
		Token: token,
	})
}
