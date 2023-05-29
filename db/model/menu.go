// All database tables do not use foreign keys
// The suffix of the associated table of the two tables is unified with relation

package model

import (
	"fmt"
	"oasis/config"
)

type Menu struct {
	Model
	ParentID  string `json:"parentID" gorm:"column:parent_id;not null;comment:父菜单ID"`
	Name      string `json:"name" gorm:"column:name;unique;not null"`
	Path      string `json:"path" gorm:"column:path;not null"`
	Component string `json:"component" gorm:"column:component;not null"` // 前端对应views路径
	Meta      `json:"meta" gorm:"embedded"`
	Hidden    bool   `json:"hidden" gorm:"column:hidden;default:0;"` //隐藏某些菜单项
	Sort      int    `json:"sort" gorm:"column:sort;default:0;"`     // 定义菜单项的显示顺序
	Children  []Menu `json:"children" gorm:"-"`
	Apis      []Api  `json:"apis" gorm:"many2many:menu_api_relation;"`
}

type Meta struct {
	KeepAlive   bool   `json:"keepAlive" gorm:"column:keep_alive;default:0"`     // 是否缓存
	Title       string `json:"title" gorm:"column:title;not null"`               // 菜单名
	Icon        string `json:"icon" gorm:"column:icon;default:'';"`              // 菜单图标
	DefaultMenu bool   `json:"defaultMenu" gorm:"column:default_menu;default:0"` // 是否是基础路由
	CloseTab    bool   `json:"closeTab" gorm:"column:close_tab;default:0"`       // 自动关闭tab
}

func (menu *Menu) LinkApis() error {
	db := config.DB

	// 获取菜单关联的已存在的 API
	existingApis := make([]Api, 0)
	err := db.Model(&menu).Association("Apis").Find(&existingApis)
	if err != nil {
		return err
	}

	// 构建现有的 API map，用于快速检索
	existingApiMap := make(map[string]Api)
	for _, existingApi := range existingApis {
		key := fmt.Sprintf("%s:%s:%s", existingApi.Group, existingApi.Path, existingApi.Method)
		existingApiMap[key] = existingApi
	}

	// 处理每个菜单关联的 API
	//for _, api := range menu.Apis {
	//	// 检查 API 是否已存在
	//	key := fmt.Sprintf("%s:%s:%s", api.Group, api.Path, api.Method)
	//	existingApi, ok := existingApiMap[key]
	//	if !ok {
	//		// 如果 API 不存在，则创建新的 API 记录
	//		result := db.Create(&api)
	//		if result.Error != nil {
	//			return result.Error
	//		}
	//	} else {
	//		// 如果 API 已存在，直接使用现有的 API 记录
	//		api.ID = existingApi.ID
	//	}
	//}

	// 关联菜单与 API
	err = db.Model(&menu).Association("Apis").Replace(&menu.Apis)
	if err != nil {
		return err
	}

	return nil
}

func (menu *Menu) CreateMenu() error {
	db := config.DB

	//menu.ID = uuid.New()

	// 创建菜单记录
	result := db.Create(&menu)
	if result.Error != nil {
		return result.Error
	}

	// 关联菜单与 API
	err := menu.LinkApis()
	if err != nil {
		return err
	}

	return nil
}

func (menu *Menu) CreateMultipleMenu(menus []Menu) error {
	db := config.DB

	tx := db.Begin()

	for i := range menus {
		//menus[i].ID = uuid.New()
		result := tx.Create(&menus[i])
		if result.Error != nil {
			tx.Rollback()
			return result.Error
		}
	}

	tx.Commit()

	return nil
}

func (menu *Menu) DeleteAllMenu() error {
	db := config.DB

	result := db.Exec("truncate table menus")
	if result.Error != nil {
		return result.Error
	}

	return nil
}
