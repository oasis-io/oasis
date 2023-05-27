package model

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"oasis/config"
)

// All database tables do not use foreign keys
// The suffix of the associated table of the two tables is unified with relation

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

func (u *User) GetUserByUsername() (*User, error) {
	var user User
	db := config.DB

	result := db.Model(u).Where("username = ?", u.Username).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("no user")
		}
		return nil, result.Error
	}
	return &user, nil
}

func (u *User) UpdateUser() error {
	db := config.DB

	updates := make(map[string]interface{})

	if u.Email != "" {
		updates["email"] = u.Email
	}
	if u.Phone != "" {
		updates["phone"] = u.Phone
	}
	if u.Password != "" {
		updates["password"] = u.Password
	}

	if err := db.Model(u).Updates(updates).Error; err != nil {
		return err
	}

	return nil
}

func (u *User) CreateUser() (err error) {
	db := config.DB

	result := db.Create(&u)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (u *User) DeleteUser() (err error) {

	db := config.DB

	db.Where("username = ?", u.Username).Preload("Roles").Find(&u)

	if u.ID != 0 {
		db.Model(&u).Association("Roles").Delete(u.Roles)
	}

	db.Delete(&u)
	var checkUser User
	db.Where("username = ?", u.Username).First(&checkUser)
	if checkUser.ID != 0 {
		return fmt.Errorf("User delete failed!")
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
