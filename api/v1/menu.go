package v1

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"oasis/app/response"
	"oasis/db"
	"oasis/db/model"
	"oasis/pkg/log"
)

type MenuResponse struct {
	Menus []model.Menu `json:"menus"`
}

func GetMenu(c *gin.Context) {

	menus, err := db.GetMenuTree()
	if err != nil {
		log.Error("获取菜单失败!", zap.Error(err))
		response.Error(c, err.Error())
	}

	if menus == nil {
		menus = []model.Menu{}
	}

	response.SendSuccessData(c, "获取菜单成功", MenuResponse{
		Menus: menus,
	})
}
