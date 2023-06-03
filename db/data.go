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
		// 基础API
		{
			Group:  "基础API",
			Desc:   "查询菜单",
			Path:   "/v1/menu",
			Method: "POST",
		},
		// 用户管理
		{
			Group:  "用户管理",
			Desc:   "查询用户",
			Path:   "/v1/user",
			Method: "POST",
		},
		{
			Group:  "用户管理",
			Desc:   "删除用户",
			Path:   "/v1/user",
			Method: "DELETE",
		},
		{
			Group:  "用户管理",
			Desc:   "修改用户",
			Path:   "/v1/user",
			Method: "PATCH",
		},
		{
			Group:  "用户管理",
			Desc:   "创建用户",
			Path:   "/v1/user/add",
			Method: "POST",
		},
		{
			Group:  "用户管理",
			Desc:   "查询用户列表",
			Path:   "/v1/user/list",
			Method: "POST",
		},
		{
			Group:  "用户管理",
			Desc:   "查询用户信息",
			Path:   "/v1/user/info",
			Method: "GET",
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

// 一级菜单00、二级菜单00、三级菜单00
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
			Sort: 1,
		},
		{
			ParentID:  "0",
			Name:      "SQLQuery",
			Path:      "sql",
			Component: "views/sql/index.vue",
			Meta: model.Meta{
				Title: "SQL查询",
				Icon:  "Search",
			},
			Sort: 2,
		},
		{
			ParentID:  "0",
			Name:      "Instance",
			Path:      "instance",
			Component: "views/instance/index.vue",
			Meta: model.Meta{
				Title: "实例管理",
				Icon:  "Menu",
			},
			Sort: 3,
		},
		{
			ParentID:  "3",
			Name:      "InstanceList",
			Path:      "list",
			Component: "views/instance/InstanceList/index.vue",
			Meta: model.Meta{
				Title: "实例列表",
			},
			Sort: 301,
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
			Sort: 4,
		},
		{
			ParentID:  "4",
			Name:      "UserList",
			Path:      "list",
			Component: "views/user/UserList/index.vue",
			Meta: model.Meta{
				Title: "用户管理",
			},
			Sort: 401,
		},
		{
			ParentID:  "4",
			Name:      "UserRole",
			Path:      "role",
			Component: "views/user/UserRole/index.vue",
			Meta: model.Meta{
				Title: "角色管理",
			},
			Sort: 402,
		},
		{
			ParentID:  "4",
			Name:      "UserGroup",
			Path:      "group",
			Component: "views/user/UserGroup/index.vue",
			Meta: model.Meta{
				Title: "用户组管理",
			},
			Sort: 403,
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
