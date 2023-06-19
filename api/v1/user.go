package v1

import (
	"github.com/gin-gonic/gin"
	"oasis/app/response"
	"oasis/db/model"
	"oasis/pkg/log"
	"oasis/pkg/utils"
)

func GetUserInfo(c *gin.Context) {
	name, err := utils.GetTokenUserName(c)
	if err != nil {
		response.Error(c, "解析token错误")
		return
	}

	user := model.User{
		Username: name,
	}
	userInfo, err := user.GetUserByUsername()
	if err != nil {
		response.Error(c, err.Error())
		return
	}

	response.SendSuccessData(c, "获取用户信息成功", UserResponse{
		User: *userInfo,
	})
}

func GetUserList(c *gin.Context) {
	var req PageInfo
	var userRes []UserRes

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("parameter binding errors: " + err.Error())
		response.Error(c, "parameter binding errors")
		return
	}

	user := model.User{}
	userList, count, err := user.GetUserList(req.PageSize, req.CurrentPage)
	if err != nil {
		response.Error(c, err.Error())
		return
	}

	if len(userList) <= 0 {
		response.Error(c, "No user found!")
		return
	}

	for _, v := range userList {
		roleResponses := make([]RoleResponse, len(v.Roles))
		for i, role := range v.Roles {
			roleResponses[i] = RoleResponse{Name: role.Name}
		}
		userRes = append(userRes,
			UserRes{
				Username: v.Username,
				Email:    v.Email,
				Phone:    v.Phone,
				Roles:    roleResponses,
			})
	}

	response.SendSuccessData(c, "获取用户列表成功", PageResponse{
		Data:        userRes,
		Total:       count,
		PageSize:    req.PageSize,
		CurrentPage: req.CurrentPage,
	})
}

func GetUser(c *gin.Context) {
	var req model.User
	var userRes []UserRes

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("parameter binding errors: " + err.Error())
		response.Error(c, "parameter binding errors")
		return
	}

	user := model.User{
		Username: req.Username,
	}

	userInfo, err := user.QueryUserAndRolesByUsername()
	if err != nil {
		response.Error(c, err.Error())
		return
	}

	roles := userInfo.Roles

	roleResponses := make([]RoleResponse, len(roles))
	for i, role := range roles {
		roleResponses[i] = RoleResponse{Name: role.Name}
	}

	userRes = append(userRes, UserRes{
		Username: userInfo.Username,
		Email:    userInfo.Email,
		Phone:    userInfo.Phone,
		Password: userInfo.Password,
		Roles:    roleResponses,
	})

	response.SendSuccessData(c, "获取用户成功", PageResponse{
		Data: userRes,
	})
}

func GetUsers(c *gin.Context) {
	user := model.User{}
	userNames, err := user.GetUserNames()
	if err != nil {
		response.Error(c, err.Error())
		return
	}

	response.SendSuccessData(c, "获取所有用户成功", PageResponse{
		Data: userNames,
	})
}

func CreateUser(c *gin.Context) {
	var req struct {
		model.User
		RoleNames []string `json:"roles"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("parameter binding errors: " + err.Error())
		response.Error(c, "parameter binding errors")
		return
	}

	var roles []*model.UserRole

	if len(req.RoleNames) != 0 {
		for _, roleName := range req.RoleNames {
			role, err := new(model.UserRole).GetRoleName(roleName)
			if err != nil {
				response.Error(c, "Role not found: "+roleName)
				return
			}
			roles = append(roles, role)
		}
	} else {
		// 默认添加CONNECT角色
		defaultRole, err := new(model.UserRole).GetRoleName("CONNECT")
		if err != nil {
			response.Error(c, "Default role CONNECT not found")
			return
		}
		roles = append(roles, defaultRole)
	}

	user := model.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		Phone:    req.Phone,
		Roles:    roles,
	}

	err := user.CreateUser()
	if err != nil {
		log.Error(err.Error())
		response.Error(c, err.Error())
		return
	}

	response.Success(c)
}

func UpdateUser(c *gin.Context) {
	var req UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("parameter binding errors: " + err.Error())
		response.Error(c, "parameter binding errors")
		return
	}

	user := model.User{
		Username: req.Username,
		Email:    req.Email,
		Phone:    req.Phone,
	}

	if req.Email != "" || req.Phone != "" {
		if err := user.UpdateUser(); err != nil {
			log.Error("database update error：" + err.Error())
			response.Error(c, "database update error")
			return
		}
	}

	if err := user.UpdateRoles(req.Roles); err != nil {
		log.Error("database update error：" + err.Error())
		response.Error(c, "database update error")
		return
	}

	response.Success(c)
}

func UpdateUserPassword(c *gin.Context) {
	var req UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("parameter binding errors: " + err.Error())
		response.Error(c, "parameter binding errors")
		return
	}

	user := model.User{
		Username: req.Username,
		Password: req.Password,
	}

	if err := user.UpdateUserPassword(); err != nil {
		log.Error("database update error：" + err.Error())
		response.Error(c, "database update error")
		return
	}

	response.Success(c)

}

func DeleteUser(c *gin.Context) {
	var req model.User

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("parameter binding errors: " + err.Error())
		response.Error(c, "parameter binding errors")
		return
	}

	user := model.User{
		Username: req.Username,
	}

	err := user.DeleteUserByUsername()
	if err != nil {
		response.Error(c, err.Error())
	}

	response.Success(c)
}
