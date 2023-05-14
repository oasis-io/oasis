package model

type UserGroup struct {
	Model
	Name  string      `json:"name" gorm:"column:name;index:uk_name,unique;not null;"`
	Desc  string      `json:"desc" gorm:"column:desc"`
	Users []*User     `gorm:"many2many:user_group_relation"`
	Roles []*UserRole `gorm:"many2many:user_group_role_relation;"`
}
