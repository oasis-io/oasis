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
	RoleID uint `json:"roleId" gorm:"column:role_id;not null;uniqueIndex:idx_role_menu"`
	MenuID uint `json:"menuId" gorm:"column:menu_id;not null;uniqueIndex:idx_role_menu"`
}

func CreateRoleMenuRelations(roleID uint, menuIDs []uint) error {
	db := config.DB
	db = db.Begin()

	for _, menuID := range menuIDs {
		relation := RoleMenuRelation{
			RoleID: roleID,
			MenuID: menuID,
		}
		if result := db.Create(&relation); result.Error != nil {
			db.Rollback()
			return result.Error
		}
	}

	db.Commit()
	return nil
}

func (r *UserRole) addDefaultRolePermission() error {
	defaultPermissions := []gormadapter.CasbinRule{
		// Menu
		{Ptype: "p", V0: strings.ToUpper(r.Name), V1: "/v1/menu", V2: "POST"},
		// Home
		{Ptype: "p", V0: strings.ToUpper(r.Name), V1: "/v1/home", V2: "POST"},
	}

	db := config.DB

	for _, permission := range defaultPermissions {
		result := db.Create(&permission)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}

func (r *UserRole) CreateRole() error {
	db := config.DB

	// 角色名称统一大写
	r.Name = strings.ToUpper(r.Name)

	result := db.Create(r)
	if result.Error != nil {
		return result.Error
	}

	// Add default permission for the new role.
	//err := r.addDefaultRolePermission()
	//if err != nil {
	//	return err
	//}

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

	// Clear all RoleMenuRelation
	err = tx.Where("role_id = ?", r.ID).Delete(&RoleMenuRelation{}).Error
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
	//err = tx.Where("v0 = ?", r.Name).Delete(&gormadapter.CasbinRule{}).Error
	//if err != nil {
	//	tx.Rollback()
	//	return err
	//}

	// If everything went well, commit the transaction
	tx.Commit()

	return nil
}

// GetRoleNames 返回所有角色信息
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

// GetRoleName 返回指定角色信息
func (r *UserRole) GetRoleName(name string) (*UserRole, error) {
	var role UserRole
	db := config.DB

	result := db.Where("name = ?", name).First(&role)
	if result.Error != nil {
		return nil, result.Error
	}

	return &role, nil
}

func (r *UserRole) GetRoleList(pageSize, currentPage int) ([]UserRole, int64, error) {
	var roleList []UserRole
	var count int64

	db := config.DB

	limit := pageSize
	offset := pageSize * (currentPage - 1)

	result := db.Limit(limit).Offset(offset).Find(&roleList)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	// 获取总记录数
	result = db.Model(&UserRole{}).Count(&count)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	return roleList, count, nil
}
