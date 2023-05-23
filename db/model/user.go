package model

import (
	"fmt"
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
