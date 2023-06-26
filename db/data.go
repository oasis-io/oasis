package db

import (
	"oasis/config"
	"oasis/db/model"
	"oasis/pkg/log"
)

func InsertData() error {
	if err := CreateAdminUser(); err != nil {
		return err
	}

	if err := insertApi(); err != nil {
		return err
	}

	if err := insertMenu(); err != nil {
		return err
	}

	return nil
}

func CreateDefaultRoles() error {

	data := []model.UserRole{
		{
			Name: "CONNECT",
			Desc: "允许用户连接",
		},
		{
			Name: "DBA",
			Desc: "拥有全部特权，是系统最高权限",
		},
	}

	role := model.UserRole{}

	err := role.CreateMultipleRoles(data)
	if err != nil {
		return err
	}
	return nil
}

func CreateAdminUser() error {
	user := model.User{
		Username: config.DefaultAdminUsername,
		Password: config.DefaultAdminPassword,
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
		log.Info("admin administrator user does not exist, initialize data.")
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
		{
			Group:  "首页",
			Desc:   "首页",
			Path:   "/v1/home",
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
			Path:   "/v1/user",
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
			Sort: 3,
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
			Sort: 4,
		},
		{
			ParentID:  "4",
			Name:      "InstanceList",
			Path:      "list",
			Component: "views/instance/instance/index.vue",
			Meta: model.Meta{
				Title: "实例列表",
			},
			Sort: 401,
		},
		{
			ParentID:  "4",
			Name:      "Session",
			Path:      "session",
			Component: "views/instance/session/index.vue",
			Meta: model.Meta{
				Title: "会话管理",
			},
			Sort: 402,
		},
		{
			ParentID:  "4",
			Name:      "Database",
			Path:      "database",
			Component: "views/instance/database/index.vue",
			Meta: model.Meta{
				Title: "数据库管理",
			},
			Sort: 403,
		},
		{
			ParentID:  "4",
			Name:      "Account",
			Path:      "account",
			Component: "views/instance/account/index.vue",
			Meta: model.Meta{
				Title: "帐号管理",
			},
			Sort: 404,
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
			Sort: 5,
		},
		{
			ParentID:  "5",
			Name:      "UserList",
			Path:      "list",
			Component: "views/user/user/index.vue",
			Meta: model.Meta{
				Title: "用户管理",
			},
			Sort: 501,
		},
		{
			ParentID:  "5",
			Name:      "UserRole",
			Path:      "role",
			Component: "views/user/role/index.vue",
			Meta: model.Meta{
				Title: "角色管理",
			},
			Sort: 502,
		},
		{
			ParentID:  "5",
			Name:      "UserGroup",
			Path:      "group",
			Component: "views/user/group/index.vue",
			Meta: model.Meta{
				Title: "用户组管理",
			},
			Sort: 503,
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
