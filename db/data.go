package db

import (
	"oasis/db/model"
)

func InsertData() error {
	errOne := insertUser()
	if errOne != nil {
		return errOne
	}

	errTwo := insertApi()
	if errTwo != nil {
		return errTwo
	}

	err := insertMenu()
	if err != nil {
		return err
	}

	return nil
}

func insertUser() error {
	user := model.User{
		Username: "admin",
		Password: "123456",
	}

	// 查询用户名是否存在
	foundUser, err := user.GetUserByUsername()
	if err != nil {
		return err
	}

	if foundUser == nil {
		err := user.CreateUser()
		if err != nil {
			return err
		}
	}

	return nil
}

func insertApi() error {

	table := []model.Api{
		{Group: "Menu", Path: "/v1/menu", Method: "POST"},
		{Group: "UserList", Path: "/v1/user", Method: "POST"},
		{Group: "UserList", Path: "/v1/user", Method: "DELETE"},
	}

	api := model.Api{}
	if err := api.DeleteAllApis(); err != nil {
		return err
	} else {
		err := api.CreateMultipleApi(table)
		if err != nil {
			return err
		}
	}

	return nil
}

func insertMenu() error {
	table := []model.Menu{
		{
			ParentID:  "0",
			Name:      "Home",
			Path:      "home",
			Component: "views/home/index.vue",
			Meta: model.Meta{
				Title: "首页",
				Icon:  "HomeFilled",
			},
		},
		{
			ParentID:  "0",
			Name:      "User",
			Path:      "user",
			Component: "views/user/index.vue",
			Meta: model.Meta{
				Title: "用户中心",
				Icon:  "User",
			},
		},
		{
			ParentID:  "3",
			Name:      "UserList",
			Path:      "list",
			Component: "views/user/UserList/index.vue",
			Meta: model.Meta{
				Title: "用户管理",
			},
		},
		{
			ParentID:  "3",
			Name:      "UserRole",
			Path:      "role",
			Component: "views/user/UserRole/index.vue",
			Meta: model.Meta{
				Title: "角色管理",
			},
		},
		{
			ParentID:  "3",
			Name:      "UserGroup",
			Path:      "group",
			Component: "views/user/UserGroup/index.vue",
			Meta: model.Meta{
				Title: "用户组管理",
			},
		},
	}

	menu := model.Menu{}
	err := menu.CreateMultipleMenu(table)
	if err != nil {
		return err
	}

	return nil
}
