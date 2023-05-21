package v1

import (
	"github.com/gin-gonic/gin"
	"oasis/app/response"
	"oasis/config"
	"oasis/db/model"
	"oasis/pkg/utils"
)

func GetUserInfo(c *gin.Context) {
	var count int64
	var user model.User

	db := config.DB

	name := utils.GetTokenUserName(c)

	db.Where("username = ? ", name).Find(&user).Count(&count)
	if count == 0 {
		response.Error(c, "没有用户")
	}
	response.SendSuccessData(c, "获取用户信息成功", UserResponse{
		User: user,
	})
}

func GetUserList(c *gin.Context) {
	var req PageInfo
	var count int64
	var _userRes []UserRes
	var userList []model.User

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, err.Error())
		return
	}

	limit := req.PageSize
	offset := req.PageSize * (req.CurrentPage - 1)

	db := config.DB

	db.Limit(limit).Offset(offset).Find(&userList)

	// total row
	user := db.Model(&model.User{})
	user.Count(&count)

	if len(userList) <= 0 {
		response.Error(c, "No user found!")
		return
	}

	for _, v := range userList {
		_userRes = append(_userRes,
			UserRes{
				Username: v.Username,
				Email:    v.Email,
				Phone:    v.Phone,
			})
	}

	response.SendSuccessData(c, "获取用户列表成功", PageResponse{
		Data:        _userRes,
		Total:       count,
		PageSize:    req.PageSize,
		CurrentPage: req.CurrentPage,
	})
}

func GetUser(c *gin.Context) {
	var req model.User
	var count int64
	var user []model.User
	var _userRes []UserRes

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, err.Error())
		return
	}

	db := config.DB

	name := req.Username

	db.Where("username = ? ", name).Find(&user).Count(&count)

	if len(user) <= 0 {
		response.Error(c, "No user data!")
	} else {
		for _, v := range user {
			_userRes = append(_userRes,
				UserRes{
					Username: v.Username,
					Email:    v.Email,
					Phone:    v.Phone,
					Password: v.Password,
				})
		}

		response.SendSuccessData(c, "获取用户成功", PageResponse{
			Data: _userRes,
		})
	}
}

func CreateUser(c *gin.Context) {
	var req model.User

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, err.Error())
		return
	}

	user := model.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		Phone:    req.Phone,
	}

	db := config.DB

	result := db.Save(&user)
	if result.Error != nil {
		response.Error(c, result.Error.Error())
	}

	response.Success(c)
}

func UpdateUser(c *gin.Context) {
	var req model.User
	var user model.User

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, err.Error())
		return
	}

	db := config.DB

	name := req.Username

	// updateMap := map[string]interface{}{}

	if req.Email != "" {
		db.Model(&user).Where("username = ?", name).Update("email", req.Email)
	}
	if req.Phone != "" {
		db.Model(&user).Where("username = ?", name).Update("phone", req.Phone)
	}
	if req.Password != "" {
		passwd := req.Password
		db.Model(&user).Where("username = ?", name).Update("password", passwd)
	}

	response.Success(c)
}

func DeleteUser(c *gin.Context) {
	var req model.User
	var user []model.User
	var count int64

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, err.Error())
		return
	}

	name := req.Username

	db := config.DB

	db.Where("username = ?", name).Delete(&user).Count(&count)
	if len(user) <= 0 {
		response.Success(c)
	} else {
		response.Error(c, "User delete fail!")
	}
}
