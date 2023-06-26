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
		log.Error(err.Error())
		response.Error(c, err.Error())
		return
	}

	if len(groupList) <= 0 {
		response.Error(c, "Failed to create user group")
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
		log.Error(err.Error())
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

// CreateUserGroup 创建一个新的用户组
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

// UpdateUserGroup 更新用户组信息以及关联的角色、用户
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

	if err := group.UpdateUserGroupByID(); err != nil {
		log.Error("database update error：" + err.Error())
		response.Error(c, "database update user group error")
		return
	}

	if err := group.UpdateUserGroupAssociations(req.Users, req.Roles); err != nil {
		log.Error("database update user group associations error：" + err.Error())
		response.Error(c, "database update user group associations error")
		return
	}

	response.Success(c)
}

// DeleteUserGroup 删除用户组信息以及用户组关联的用户、角色关系
func DeleteUserGroup(c *gin.Context) {
	var req model.UserGroup

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("parameter binding errors: " + err.Error())
		response.Error(c, "parameter binding errors")
		return
	}

	if req.Name == "" {
		response.Error(c, "name field is empty")
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
		log.Error(err.Error())
		response.Error(c, err.Error())
		return
	}

	response.Success(c)

}
