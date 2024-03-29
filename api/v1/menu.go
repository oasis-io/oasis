package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"oasis/app/response"
	"oasis/config"
	"oasis/db"
	"oasis/db/model"
	"oasis/pkg/log"
	"oasis/pkg/utils"
	"sort"
)

type MenuResponse struct {
	Menus []model.Menu `json:"menus"`
}

// GetMenuTree 返回用户的所有菜单栏
// admin管理员直接返回所有菜单
func GetMenuTree(c *gin.Context) {
	username, err := utils.GetTokenUserName(c)
	if err != nil {
		log.Error("parsing token error", zap.Error(err))
		response.Error(c, "Parsing token error")
		return
	}

	allMenus, err := db.GetMenuTree()
	if err != nil {
		log.Error("Failed to get menu", zap.Error(err))
		response.Error(c, "Failed to get menu")
		return
	}

	if username == config.DefaultAdminUsername {
		response.SendSuccessData(c, "Get the menu successfully", MenuResponse{
			Menus: allMenus,
		})
		return
	}

	user := model.User{
		Username: username,
	}
	foundUser, err := user.QueryUserAndRolesByUsername()
	if err != nil {
		log.Error(err.Error())
		response.Error(c, "Failed to build menu, please check if permissions are added!")
		return
	}

	roles := foundUser.Roles
	var menus []model.Menu
	for _, role := range roles {
		// 查询角色关联的菜单信息
		menus, err = db.GetMenuTreeMapForRole(role.Name)
		if err != nil {
			log.Error("Failed to get menu", zap.Error(err))
			response.Error(c, "Failed to build menu, please check if permissions are added!")
			return
		}

	}

	response.SendSuccessData(c, "Get the menu successfully", MenuResponse{
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

	// 获取角色信息
	role, err := (&model.UserRole{}).GetRoleName(req.Name)
	if err != nil {
		response.Error(c, fmt.Sprintf("Unable to get role ID for name: %s", req.Name))
		return
	}

	menuIDs := make([]uint, len(req.Menus))
	for i, menu := range req.Menus {
		menuIDs[i] = menu
	}

	// 对menuIDs进行排序
	sort.Slice(menuIDs, func(i, j int) bool {
		return menuIDs[i] < menuIDs[j]
	})

	log.Info(fmt.Sprintf("menuIDs: %v", menuIDs))

	roleID := role.ID
	if err := model.UpdateRoleMenuRelations(roleID, menuIDs); err != nil {
		response.Error(c, fmt.Sprintf("Unable to update role-menu relations: %v", err))
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

	err := db.UpdateApiPermissions(req.Name, req.Apis)
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
