// All database tables do not use foreign keys
// The suffix of the associated table of the two tables is unified with relation

package model

import (
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
}

type Meta struct {
	KeepAlive   bool   `json:"keepAlive" gorm:"column:keep_alive;default:0"`     // 是否缓存
	Title       string `json:"title" gorm:"column:title;not null"`               // 菜单名
	Icon        string `json:"icon" gorm:"column:icon;default:'';"`              // 菜单图标
	DefaultMenu bool   `json:"defaultMenu" gorm:"column:default_menu;default:0"` // 是否是基础路由
	CloseTab    bool   `json:"closeTab" gorm:"column:close_tab;default:0"`       // 自动关闭tab
}

func (menu *Menu) CreateMenu() error {
	db := config.DB

	//menu.ID = uuid.New()

	// 创建菜单记录
	result := db.Create(&menu)
	if result.Error != nil {
		return result.Error
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
