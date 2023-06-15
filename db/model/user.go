// All database tables do not use foreign keys
// The suffix of the associated table of the two tables is unified with relation

package model

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"oasis/config"
)

type User struct {
	Model
	Username   string       `json:"username" gorm:"column:username;index:uk_name,unique;not null;"`
	Password   string       `json:"password"  gorm:"column:password;not null;"`
	Email      string       `json:"email" gorm:"column:email;"`
	Phone      string       `json:"phone" gorm:"column:phone;"`
	IsEnable   bool         `json:"is_enable" gorm:"column:is_enable;type:tinyint(1);default:0;comment:0:enable,1:disabled"`
	Roles      []*UserRole  `gorm:"many2many:user_role_relation;"`
	UserGroups []*UserGroup `gorm:"many2many:user_group_relation"`
}

// GetUserNames 返回所有用户信息
func (u *User) GetUserNames() ([]string, error) {
	db := config.DB

	var users []User
	result := db.Select("username").Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	var userNames []string
	for _, user := range users {
		userNames = append(userNames, user.Username)
	}

	return userNames, nil
}

func (u *User) GetUserByUsername() (*User, error) {
	var user User
	db := config.DB

	result := db.Model(u).Where("username = ?", u.Username).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}

func (u *User) UpdateUser() error {
	db := config.DB
	updates := map[string]interface{}{
		"email":    u.Email,
		"phone":    u.Phone,
		"password": u.Password,
	}

	if u.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		updates["password"] = string(hashedPassword)
	}

	result := db.Model(&User{}).Where("username = ?", u.Username).Updates(updates)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}

func (u *User) UpdateUserPassword() error {
	db := config.DB
	updates := map[string]interface{}{
		"password": u.Password,
	}

	if u.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		updates["password"] = string(hashedPassword)
	}

	result := db.Model(&User{}).Where("username = ?", u.Username).Updates(updates)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}

func (u *User) UpdateRoles(roles []string) error {
	db := config.DB
	var userRoles []*UserRole
	if err := db.Where("name IN ?", roles).Find(&userRoles).Error; err != nil {
		return err
	}

	var user User
	if err := db.Where("username = ?", u.Username).First(&user).Error; err != nil {
		return err
	}

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Model(&user).Association("Roles").Replace(userRoles); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (u *User) CreateUser() error {
	db := config.DB

	// bcrypt加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)

	// 创建用户
	result := db.Create(&u)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// DeleteUserByUsername 删除用户，并且删除用户关联的角色、用户组关联信息
func (u *User) DeleteUserByUsername() (err error) {

	db := config.DB

	db.Where("username = ?", u.Username).Preload("Roles").Preload("UserGroups").Find(&u)

	if u.ID != 0 {
		db.Model(&u).Association("Roles").Delete(u.Roles)
		db.Model(&u).Association("UserGroups").Delete(u.UserGroups)
	}

	db.Delete(&u)
	var checkUser User
	db.Where("username = ?", u.Username).First(&checkUser)

	if checkUser.ID != 0 {
		return fmt.Errorf("user not deleted")
	}

	return nil
}

func (u *User) QueryUserAndRolesByUsername() (*User, error) {
	var user User
	db := config.DB

	result := db.Where("username = ?", u.Username).Preload("Roles").First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (u *User) GetUserList(pageSize, currentPage int) ([]User, int64, error) {
	var userList []User
	var count int64

	db := config.DB

	limit := pageSize
	offset := pageSize * (currentPage - 1)

	db.Preload("Roles").Limit(limit).Offset(offset).Find(&userList)

	// 获取总记录数
	db.Model(&User{}).Count(&count)

	return userList, count, nil
}
