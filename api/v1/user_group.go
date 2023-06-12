package v1

import (
	"github.com/gin-gonic/gin"
	"oasis/app/response"
	"oasis/db/model"
	"oasis/pkg/log"
)

func GetUserGroupList(c *gin.Context) {
	var req PageInfo
	var groupRes []GroupResponse

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("parameter binding errors: " + err.Error())
		response.Error(c, "parameter binding errors")
		return
	}

	group := model.UserGroup{}
	groupList, count, err := group.GetGroupList(req.PageSize, req.CurrentPage)
	if err != nil {
		response.Error(c, err.Error())
		return
	}

	if len(groupList) <= 0 {
		response.Error(c, "no user group")
		return
	}

	for _, v := range groupList {
		groupRes = append(groupRes,
			GroupResponse{
				Name: v.Name,
				Desc: v.Desc,
			})
	}

	response.SendSuccessData(c, "获取角色列表成功", PageResponse{
		Data:        groupRes,
		Total:       count,
		PageSize:    req.PageSize,
		CurrentPage: req.CurrentPage,
	})

}

// GetUserGroup 获取单个用户组信息
func GetUserGroup(c *gin.Context) {
	var req model.UserGroup
	var groupRes []GroupRes

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("parameter binding errors: " + err.Error())
		response.Error(c, "parameter binding errors")
		return
	}

	group := model.UserGroup{
		Name: req.Name,
	}

	info, err := group.QueryGroupAndUsersRolesByName()
	if err != nil {
		response.Error(c, err.Error())
		return
	}

	roles := info.Roles

	roleResponses := make([]RoleResponse, len(roles))
	for i, role := range roles {
		roleResponses[i] = RoleResponse{Name: role.Name}
	}

	users := info.Users
	userResponses := make([]UserRes, len(users))
	for i, user := range users {
		userResponses[i] = UserRes{Username: user.Username}
	}

	groupRes = append(groupRes, GroupRes{
		ID:    info.ID,
		Name:  info.Name,
		Desc:  info.Desc,
		Users: userResponses,
		Roles: roleResponses,
	})

	response.SendSuccessData(c, "获取用户组成功", PageResponse{
		Data: groupRes,
	})

}

// GetUserGroups 获取所有用户组信息
func GetUserGroups(c *gin.Context) {
	response.Success(c)

}

func CreateUserGroup(c *gin.Context) {
	var req struct {
		model.UserGroup
		RoleNames []string `json:"roles"`
		UserNames []string `json:"users"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("parameter binding errors: " + err.Error())
		response.Error(c, "parameter binding errors")
		return
	}

	var roles []*model.UserRole
	for _, roleName := range req.RoleNames {
		role, err := new(model.UserRole).GetRoleName(roleName)
		if err != nil {
			response.Error(c, "Role not found: "+roleName)
			return
		}
		roles = append(roles, role)
	}

	var users []*model.User
	for _, userName := range req.UserNames {
		u := model.User{
			Username: userName,
		}
		user, err := u.GetUserByUsername()
		if err != nil {
			response.Error(c, "User not found: "+userName)
			return
		}
		users = append(users, user)
	}

	group := model.UserGroup{
		Name:  req.Name,
		Desc:  req.Desc,
		Users: users,
		Roles: roles,
	}

	err := group.CreateUserGroup()
	if err != nil {
		response.Error(c, err.Error())
		return
	}

	response.Success(c)
}

func UpdateUserGroup(c *gin.Context) {
	var req GroupRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("parameter binding errors: " + err.Error())
		response.Error(c, "parameter binding errors")
		return
	}

	group := model.UserGroup{
		Model: model.Model{
			ID: req.ID,
		},
		Name: req.Name,
		Desc: req.Desc,
	}

	err := group.UpdateUserGroupByID()
	if err != nil {
		log.Error("database update error：" + err.Error())
		response.Error(c, "database update error")
		return
	}

	response.Success(c)
}

func DeleteUserGroup(c *gin.Context) {
	var req model.UserGroup

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("parameter binding errors: " + err.Error())
		response.Error(c, "parameter binding errors")
		return
	}

	if req.Name == "" {
		response.Error(c, "name 字段不能为空")
		return
	}

	role := new(model.UserGroup)

	foundGroup, err := role.GetGroupByName(req.Name)
	if err != nil {
		response.Error(c, err.Error())
		return
	}

	err = foundGroup.DeleteUserGroup()
	if err != nil {
		response.Error(c, err.Error())
		return
	}

	response.Success(c)

}
