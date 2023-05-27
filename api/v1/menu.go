package v1

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"oasis/app/response"
	"oasis/db"
	"oasis/db/model"
	"oasis/pkg/log"
	"oasis/pkg/utils"
)

type MenuResponse struct {
	Menus []model.Menu `json:"menus"`
}

func GetMenu(c *gin.Context) {
	username, err := utils.GetTokenUserName(c)
	if err != nil {
		response.Error(c, "解析token错误")
		return
	}

	allMenus, err := db.GetMenuTree()
	if err != nil {
		log.Error("获取菜单失败!", zap.Error(err))
		response.Error(c, err.Error())
		return
	}

	if username == "admin" {
		response.SendSuccessData(c, "获取菜单成功", MenuResponse{
			Menus: allMenus,
		})
		return
	}

	user := model.User{
		Username: username,
	}
	foundUser, err := user.QueryUserAndRolesByUsername()
	if err != nil {
		response.Error(c, err.Error())
		return
	}

	roles := foundUser.Roles
	var menus []model.Menu
	for _, role := range roles {
		// 成功得到角色名
		log.Info(role.Name)
		casbinRules, err := db.GetCasbinRulesByRole(role.Name)
		if err != nil {
			response.Error(c, err.Error())
			return
		}

		// 成功得到API
		roleAPIs := make(map[string]struct{})
		for _, rule := range casbinRules {

			roleAPIs[rule.V1] = struct{}{}
			log.Info(rule.V1)
		}

		for _, menu := range allMenus {
			if _, ok := roleAPIs[menu.Path]; ok {
				menus = append(menus, menu)
			}
		}
	}

	response.SendSuccessData(c, "获取菜单成功", MenuResponse{
		Menus: menus,
	})
}

func contains(slice []string, item string) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}
