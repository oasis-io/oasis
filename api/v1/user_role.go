package v1

import (
	"github.com/gin-gonic/gin"
	"oasis/app/response"
	"oasis/config"
	"oasis/db/model"
)

func GetRoleList(c *gin.Context) {
	var req PageInfo
	var count int64
	var _roleRes []RoleResponse
	var roleList []model.UserRole

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, err.Error())
		return
	}

	limit := req.PageSize
	offset := req.PageSize * (req.CurrentPage - 1)

	db := config.DB

	db.Limit(limit).Offset(offset).Find(&roleList)

	// total row
	user := db.Model(&model.UserRole{})
	user.Count(&count)
	if len(roleList) <= 0 {
		response.Error(c, "没有找到用户")
		return
	}

	for _, v := range roleList {
		_roleRes = append(_roleRes,
			RoleResponse{
				Name: v.Name,
			})
	}

	response.SendSuccessData(c, "获取角色列表成功", PageResponse{
		Data:        _roleRes,
		Total:       count,
		PageSize:    req.PageSize,
		CurrentPage: req.CurrentPage,
	})

}

func GetRole(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "OK",
	})
}

func GetRoles(c *gin.Context) {
	role := model.UserRole{}
	roleNames, err := role.GetRoleNames()
	if err != nil {
		response.Error(c, err.Error())
		return
	}

	response.SendSuccessData(c, "获取所有角色名成功", PageResponse{
		Data: roleNames,
	})
}

func CreateRole(c *gin.Context) {
	var req model.UserRole

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, err.Error())
		return
	}

	role := model.UserRole{
		Name: req.Name,
	}

	err := role.CreateRole()
	if err != nil {
		response.Error(c, err.Error())
		return
	}

	response.Success(c)

}

func UpdateRole(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "OK",
	})
}

func DeleteRole(c *gin.Context) {
	var req model.UserRole

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, err.Error())
		return
	}

	if req.Name == "" {
		response.Error(c, "name 字段不能为空")
		return
	}

	role := new(model.UserRole)

	foundRole, err := role.FindByName(req.Name)
	if err != nil {
		response.Error(c, err.Error())
		return
	}

	err = foundRole.DeleteRole()
	if err != nil {
		response.Error(c, err.Error())
		return
	}

	response.Success(c)
}
