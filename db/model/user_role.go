package model

import (
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"oasis/config"
	"strings"
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

func AddDefaultRolePermission(role *UserRole) error {
	defaultPermission := &gormadapter.CasbinRule{
		Ptype: "p",
		V0:    strings.ToUpper(role.Name),
		V1:    "/v1/menu",
		V2:    "POST",
	}

	db := config.DB
	result := db.Create(defaultPermission)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *UserRole) CreateRole() (err error) {
	db := config.DB

	// 角色名称统一大写
	r.Name = strings.ToUpper(r.Name)

	result := db.Create(r)
	if result.Error != nil {
		return result.Error
	}

	// Add default permission for the new role.
	err = AddDefaultRolePermission(r)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRole) DeleteRole() error {
	db := config.DB

	// Start a new transaction
	tx := db.Begin()

	// Clear all user associations
	err := tx.Model(&r).Association("Users").Clear()
	if err != nil {
		tx.Rollback()
		return err
	}

	// Then, delete the role itself
	result := tx.Delete(&r)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	// Delete the casbin rules associated with the role
	err = tx.Where("v0 = ?", r.Name).Delete(&gormadapter.CasbinRule{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// If everything went well, commit the transaction
	tx.Commit()

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
