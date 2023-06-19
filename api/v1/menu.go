package v1

import (
	"fmt"
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

type MenuApiResponse struct {
	Api []model.Api `json:"apis"`
}

// GetMenuTree 返回用户的所有菜单栏
// admin管理员直接返回所有菜单
func GetMenuTree(c *gin.Context) {
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
		// 查询角色关联的菜单信息
		menus, err = db.GetMenuTreeMapForRole(role.Name)
		if err != nil {
			log.Error("获取菜单失败!", zap.Error(err))
			response.Error(c, err.Error())
			return
		}

	}

	response.SendSuccessData(c, "获取菜单成功", MenuResponse{
		Menus: menus,
	})
}

// GetBaseMenuTree 前端设置权限需要返回菜单信息
func GetBaseMenuTree(c *gin.Context) {
	allMenus, err := db.GetMenuTreeMap()
	if err != nil {
		log.Error("获取菜单失败!", zap.Error(err))
		response.Error(c, err.Error())
		return
	}

	response.SendSuccessData(c, "获取菜单成功", MenuResponse{
		Menus: allMenus,
	})
}

// MenuPermissions 添加角色关联的菜单
func MenuPermissions(c *gin.Context) {
	var req MenuRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("parameter binding errors: " + err.Error())
		response.Error(c, "parameter binding errors")
		return
	}

	role, err := (&model.UserRole{}).GetRoleName(req.Name)
	if err != nil {
		response.Error(c, fmt.Sprintf("Unable to get role ID for name: %s", req.Name))
		return
	}

	menuIDs := make([]uint, len(req.Menus))
	for i, menu := range req.Menus {
		menuIDs[i] = menu.ID
	}

	roleID := role.ID
	if err := model.CreateRoleMenuRelations(roleID, menuIDs); err != nil {
		response.Error(c, fmt.Sprintf("Unable to create role-menu relations: %v", err))
		return
	}

	response.Success(c)
}

// MenuApiPermissions 添加角色关联的API权限并插入casbin_rule表
func MenuApiPermissions(c *gin.Context) {
	var req MenuRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("parameter binding errors: " + err.Error())
		response.Error(c, "parameter binding errors")
		return
	}

	err := db.AddApiPermissions(req.Name, req.Apis)
	if err != nil {
		response.Error(c, err.Error())
		return
	}

	response.Success(c)
}

func GetBaseMenuApi(c *gin.Context) {
	api := model.Api{}
	allApi, err := api.GetAllApi()
	if err != nil {
		log.Error("获取菜单API失败!", zap.Error(err))
		response.Error(c, err.Error())
		return
	}

	apiTree := db.BuildApiTree(allApi)

	response.SendSuccessData(c, "获取菜单API成功", apiTree)
}

func GetMenuAuthorized(c *gin.Context) {
	var req RoleRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("parameter binding errors: " + err.Error())
		response.Error(c, "parameter binding errors")
		return
	}

	// 查询角色关联的菜单信息
	menus, err := db.GetMenuTreeMapForRole(req.Name)
	if err != nil {
		log.Error("获取菜单失败!", zap.Error(err))
		response.Error(c, err.Error())
		return
	}

	response.SendSuccessData(c, "获取菜单成功", MenuResponse{
		Menus: menus,
	})
}

func GetMenuApiAuthorized(c *gin.Context) {
	var req RoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("parameter binding errors: " + err.Error())
		response.Error(c, "parameter binding errors")
		return
	}

	apis, err := db.GetApisByRole(req.Name)
	if err != nil {
		response.Error(c, err.Error())
		return
	}

	apiTree := db.BuildApiTree(apis)

	response.SendSuccessData(c, "获取菜单API成功", apiTree)
}
