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

//func (u *User) UpdateUser() error {
//	db := config.DB
//
//	updates := make(map[string]interface{})
//
//	//updates["email"] = u.Email
//	//updates["phone"] = u.Phone
//	if u.Password != "" {
//		// 对新密码进行加密
//		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
//		if err != nil {
//			return err
//		}
//		updates["password"] = string(hashedPassword)
//	}
//
//	if err := db.Model(u).Where("username = ?", u.Username).Updates(updates).Error; err != nil {
//		return err
//	}
//
//	return nil
//}
//
//// UpdateRoles 更新用户的角色
//func (u *User) UpdateRoles(roles []string) error {
//	db := config.DB
//	var userRoles []*UserRole
//
//	// 查找所有的角色
//	if err := db.Where("name IN ?", roles).Find(&userRoles).Error; err != nil {
//		return err
//	}
//
//	// 获取用户对象
//	var user User
//	if err := db.Where("username = ?", u.Username).First(&user).Error; err != nil {
//		return err
//	}
//
//	// 获取用户的角色关联对象
//	association := db.Model(&user).Association("Roles")
//	//// 判断关联表是否已经有数据
//	//count := association.Count()
//	//if count > 0 {
//	//	// 如果关联表已经有数据，直接返回
//	//	return nil
//	//}
//
//	// 更新用户的角色
//	if err := association.Replace(userRoles); err != nil {
//		return err
//	}
//
//	return nil
//}

func (u *User) CreateUser() error {
	db := config.DB

	// bcrypt加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// 使用加密后的密码替换原始密码
	u.Password = string(hashedPassword)

	// 创建用户记录
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
