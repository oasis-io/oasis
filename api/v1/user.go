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
		response.Error(c, err.Error())
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
		response.Error(c, err.Error())
		return
	}

	user := model.User{
		Username: req.Username,
	}

	userInfo, err := user.GetUserByUsername()
	if err != nil {
		response.Error(c, err.Error())
		return
	}

	userRes = append(userRes, UserRes{
		Username: userInfo.Username,
		Email:    userInfo.Email,
		Phone:    userInfo.Phone,
		Password: userInfo.Password,
	})

	response.SendSuccessData(c, "获取用户成功", PageResponse{
		Data: userRes,
	})
}

func CreateUser(c *gin.Context) {
	var req struct {
		model.User
		RoleNames []string `json:"roles"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, err.Error())
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
	var req model.User

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("parameter binding errors: " + err.Error())
		response.Error(c, "parameter binding errors")
		return
	}

	user := model.User{
		Username: req.Username,
		Email:    req.Email,
		Phone:    req.Phone,
		Password: req.Password,
	}

	if err := user.UpdateUser(); err != nil {
		log.Error("database update error：" + err.Error())
		response.Error(c, "database update error")
		return
	}

	response.Success(c)
}

//func UpdateUser(c *gin.Context) {
//	var req UserRequest
//
//	if err := c.ShouldBindJSON(&req); err != nil {
//		log.Error("parameter binding errors: " + err.Error())
//		response.Error(c, "parameter binding errors")
//		return
//	}
//
//	user := model.User{
//		Username: req.Username,
//		Email:    req.Email,
//		Phone:    req.Phone,
//		Password: req.Password,
//	}
//
//	// 处理角色信息
//	roles := make([]*model.UserRole, len(req.Roles))
//	for i, roleName := range req.Roles {
//		role := &model.UserRole{Name: roleName}
//		roles[i] = role
//	}
//	user.Roles = roles
//
//	if err := user.UpdateUser(); err != nil {
//		log.Error("database update error：" + err.Error())
//		response.Error(c, "database update error")
//		return
//	}
//
//	response.Success(c)
//}

func DeleteUser(c *gin.Context) {
	var req model.User

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, err.Error())
		return
	}

	user := model.User{
		Username: req.Username,
	}

	err := user.DeleteUser()
	if err != nil {
		response.Error(c, err.Error())
	}

	response.Success(c)
}
