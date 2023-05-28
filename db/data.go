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
		{
			Group:  "Menu",
			Path:   "/v1/menu",
			Method: "POST",
		},
		{
			Group:  "UserList",
			Path:   "/v1/user",
			Method: "POST",
		},
		{
			Group:  "UserList",
			Path:   "/v1/user",
			Method: "DELETE",
		},
		{
			Group:  "UserList",
			Path:   "/v1/user",
			Method: "PATCH",
		},
		{
			Group:  "UserList",
			Path:   "/v1/user/add",
			Method: "POST",
		},
		{
			Group:  "UserList",
			Path:   "/v1/user/list",
			Method: "POST",
		},
		{
			Group:  "UserList",
			Path:   "/v1/user/info",
			Method: "GET",
		},
		{
			Group:  "UserRole",
			Path:   "/v1/user/role",
			Method: "POST",
		},
		{
			Group:  "UserRole",
			Path:   "/v1/user/role",
			Method: "DELETE",
		},
		{
			Group:  "UserRole",
			Path:   "/v1/user/role",
			Method: "PATCH",
		},
		{
			Group:  "UserRole",
			Path:   "/v1/user/role/add",
			Method: "POST",
		},
		{
			Group:  "UserRole",
			Path:   "/v1/user/role/list",
			Method: "POST",
		},
		{
			Group:  "UserGroup",
			Path:   "/v1/user/group",
			Method: "POST",
		},
		{
			Group:  "UserGroup",
			Path:   "/v1/user/group",
			Method: "DELETE",
		},
		{
			Group:  "UserGroup",
			Path:   "/v1/user/group",
			Method: "PATCH",
		},
		{
			Group:  "UserGroup",
			Path:   "/v1/user/group/add",
			Method: "POST",
		},
		{
			Group:  "UserGroup",
			Path:   "/v1/user/group/list",
			Method: "POST",
		},
		{
			Group:  "Instance",
			Path:   "/v1/instance",
			Method: "POST",
		},
		{
			Group:  "Instance",
			Path:   "/v1/instance",
			Method: "DELETE",
		},
		{
			Group:  "Instance",
			Path:   "/v1/instance",
			Method: "PATCH",
		},
		{
			Group:  "Instance",
			Path:   "/v1/instance/list",
			Method: "POST",
		},
		{
			Group:  "Instance",
			Path:   "/v1/instance/add",
			Method: "POST",
		},
		{
			Group:  "Instance",
			Path:   "/v1/instance/ping",
			Method: "POST",
		},
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

// 一级菜单100, 二级菜单在父ID后面+1,示例1001，依次类推，所以一个等级菜单最多9个
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
			Sort: 100,
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
			Sort: 101,
		},
		{
			ParentID:  "101",
			Name:      "UserList",
			Path:      "list",
			Component: "views/user/UserList/index.vue",
			Meta: model.Meta{
				Title: "用户管理",
			},
			Sort: 1011,
		},
		{
			ParentID:  "101",
			Name:      "UserRole",
			Path:      "role",
			Component: "views/user/UserRole/index.vue",
			Meta: model.Meta{
				Title: "角色管理",
			},
			Sort: 1012,
		},
		{
			ParentID:  "101",
			Name:      "UserGroup",
			Path:      "group",
			Component: "views/user/UserGroup/index.vue",
			Meta: model.Meta{
				Title: "用户组管理",
			},
			Sort: 1013,
		},
	}

	menu := model.Menu{}

	if err := menu.DeleteAllMenu(); err != nil {
		return err
	} else {
		err := menu.CreateMultipleMenu(table)
		if err != nil {
			return err
		}

	}

	return nil
}
