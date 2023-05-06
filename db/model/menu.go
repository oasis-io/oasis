package model

type Menu struct {
	Model
	Level     uint   `json:"level"`    // 菜单等级
	ParentId  string `json:"parentId"` // 顶级菜单ID
	Name      string `json:"name"`     // 菜单名称
	Path      string `json:"path"`     // 菜单路径
	Meta      `json:"meta" gorm:"embedded;"`
	Component string `json:"component"`         // 前端文件路径
	Children  []Menu `json:"children" gorm:"-"` // 子菜单，不存储在数据库
}

type Meta struct {
	ActiveName string `json:"activeName" ` //高亮菜单
	KeepAlive  bool   `json:"keepAlive"`   // 是否缓存
	Title      string `json:"title"`       // 菜单名
	Icon       string `json:"icon"`        // 菜单图标
}
