// All database tables do not use foreign keys
// The suffix of the associated table of the two tables is unified with relation

package model

import (
	"errors"
	"oasis/config"
)

type UserGroup struct {
	Model
	Name  string      `json:"name" gorm:"column:name;index:uk_name,unique;not null;"`
	Desc  string      `json:"desc" gorm:"column:desc"`
	Users []*User     `gorm:"many2many:user_group_relation"`
	Roles []*UserRole `gorm:"many2many:user_group_role_relation;"`
}

// GetGroupList 获取用户组列表
func (group *UserGroup) GetGroupList(pageSize, currentPage int) ([]UserGroup, int64, error) {
	var groupList []UserGroup
	var count int64

	db := config.DB

	limit := pageSize
	offset := pageSize * (currentPage - 1)

	result := db.Limit(limit).Offset(offset).Find(&groupList)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	// 获取总记录数
	result = db.Model(&UserGroup{}).Count(&count)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return groupList, count, nil
}

// GetGroupByName 根据用户组名称查询信息
func (group *UserGroup) GetGroupByName(name string) (*UserGroup, error) {
	var userGroup UserGroup
	db := config.DB

	result := db.Where("name = ?", name).First(&userGroup)
	if result.Error != nil {
		return nil, result.Error
	}

	return &userGroup, nil
}

// CreateUserGroup 创建一个用户组
func (group *UserGroup) CreateUserGroup() error {
	db := config.DB

	result := db.Create(group)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (group *UserGroup) DeleteUserGroup() error {
	db := config.DB

	// Start a new transaction
	tx := db.Begin()

	// Clear all user associations
	err := tx.Model(&group).Association("Users").Clear()
	if err != nil {
		tx.Rollback()
		return err
	}

	// Clear all role associations
	err = tx.Model(&group).Association("Roles").Clear()
	if err != nil {
		tx.Rollback()
		return err
	}

	// Delete UserGroup
	result := tx.Delete(&group)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	tx.Commit()

	return nil
}

// UpdateUserGroupByID 根据用户组ID去更新用户信息
func (group *UserGroup) UpdateUserGroupByID() error {
	db := config.DB

	updates := map[string]interface{}{
		"name": group.Name,
		"desc": group.Desc,
	}

	result := db.Model(&UserGroup{}).Where("id = ?", group.ID).Updates(updates)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no user group found with the specified ID")
	}

	return nil
}

func (group *UserGroup) QueryGroupAndUsersRolesByName() (*UserGroup, error) {
	var userGroup UserGroup
	db := config.DB

	result := db.Where("name = ?", group.Name).Preload("Users").Preload("Roles").First(&userGroup)
	if result.Error != nil {
		return nil, result.Error
	}

	return &userGroup, nil
}
