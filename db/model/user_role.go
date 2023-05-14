package model

type UserRole struct {
	Model
	Name       string       `json:"name" gorm:"column:name;index:uk_name,unique;not null;"`
	Desc       string       `json:"desc" gorm:"column:desc"`
	Users      []*User      `gorm:"many2many:user_role_relation;"`
	UserGroups []*UserGroup `gorm:"many2many:user_group_role_relation;"`
}

// RoleMenuRelation have better scalability
type RoleMenuRelation struct {
	Model
	RoleID string `json:"roleId" gorm:"column:role_id;not null;"`
	MenuID string `json:"menuId" gorm:"column:menu_id;not null;"`
}
