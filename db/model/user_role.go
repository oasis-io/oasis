package model

import (
	"oasis/config"
)

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

func (r *UserRole) QueryRole() (*UserRole, error) {
	var role UserRole
	db := config.DB

	result := db.Where("name = ?", r.Name).First(&role)
	if result.Error != nil {
		return nil, result.Error
	}

	return &role, nil
}

func (r *UserRole) CreateRole() (err error) {
	db := config.DB

	result := db.Create(&r)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *UserRole) DeleteRole() error {
	db := config.DB

	result := db.Where("name = ?", r.Name).Delete(&r)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *UserRole) GetRoleNames() ([]string, error) {
	db := config.DB

	var roles []UserRole
	result := db.Select("name").Find(&roles)
	if result.Error != nil {
		return nil, result.Error
	}

	var roleNames []string
	for _, role := range roles {
		roleNames = append(roleNames, role.Name)
	}

	return roleNames, nil
}

func (r *UserRole) FindByName(name string) (*UserRole, error) {
	db := config.DB
	var role UserRole
	result := db.Where("name = ?", name).First(&role)
	if result.Error != nil {
		return nil, result.Error
	}
	return &role, nil
}
