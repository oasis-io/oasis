// All database tables do not use foreign keys
// The suffix of the associated table of the two tables is unified with relation

package model

type UserGroup struct {
	Model
	Name  string      `json:"name" gorm:"column:name;index:uk_name,unique;not null;"`
	Desc  string      `json:"desc" gorm:"column:desc"`
	Users []*User     `gorm:"many2many:user_group_relation"`
	Roles []*UserRole `gorm:"many2many:user_group_role_relation;"`
}
