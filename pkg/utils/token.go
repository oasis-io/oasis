package utils

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"oasis/config"
	"oasis/db/model"
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

func GetRolesAndGroupsByUsername(username string) ([]string, error) {
	var user model.User
	var rolesAndGroups []string

	db := config.DB

	// 使用预加载来获取用户的角色和用户组
	err := db.Preload("Roles").Preload("UserGroups").Where("username = ?", username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 用户不存在
			return nil, nil
		}
		// 数据库错误
		return nil, err
	}

	// 将角色名和用户组名添加到rolesAndGroups数组中
	for _, role := range user.Roles {
		rolesAndGroups = append(rolesAndGroups, role.Name)
	}
	for _, group := range user.UserGroups {
		rolesAndGroups = append(rolesAndGroups, group.Name)
	}

	return rolesAndGroups, nil
}
