package model

import (
	"errors"
	"gorm.io/gorm"
	"oasis/config"
	"oasis/pkg/utils"
)

type Instance struct {
	Model
	Name     string `json:"name" gorm:"index:uk_name,unique;not null;"`
	DBType   string `json:"db_type" gorm:"column:db_type;not null;" example:"mysql"`
	Host     string `json:"host" gorm:"not null;"`
	Port     string `json:"port" gorm:"not null;" example:"3306"`
	User     string `json:"user" gorm:"not null;" example:"root"`
	Password string `json:"password"  gorm:"not null;"`
}

// CreateInstance 创建单条实例
func (ins *Instance) CreateInstance() error {
	db := config.DB

	// AES encryption
	encryptedPassword, err := utils.EncryptWithAES(ins.Password)
	if err != nil {
		return err
	}

	ins.Password = encryptedPassword

	// 创建用户
	result := db.Create(&ins)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// CreateMultipleInstance 批量创建实例
func (ins *Instance) CreateMultipleInstance() error {
	return nil
}

// UpdateUserGroupByID 更新单条实例
func (ins *Instance) UpdateUserGroupByID() error {
	db := config.DB

	updates := map[string]interface{}{
		"name": ins.Name,
		"host": ins.Host,
		"port": ins.Port,
		"user": ins.User,
	}

	if ins.Password != "" {
		// AES encryption
		encryptedPassword, err := utils.EncryptWithAES(ins.Password)
		if err != nil {
			return err
		}
		updates["password"] = encryptedPassword
	}

	result := db.Model(&Instance{}).Where("id = ?", ins.ID).Updates(updates)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (ins *Instance) UpdateInstancePassword() error {
	db := config.DB

	// Check if password is not empty
	if ins.Password == "" {
		return errors.New("password cannot be empty")
	}

	// AES encryption
	encryptedPassword, err := utils.EncryptWithAES(ins.Password)
	if err != nil {
		return err
	}

	// Only update password
	result := db.Model(&Instance{}).Where("name = ?", ins.Name).Update("password", encryptedPassword)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// DeleteInstance 删除单条实例
func (ins *Instance) DeleteInstance() error {
	db := config.DB

	if err := db.Where("name = ?", ins.Name).Delete(&Instance{}).Error; err != nil {
		return err
	}

	return nil
}

// GetAllInstance 获取所有实例列表并且返回
func (ins *Instance) GetAllInstance() ([]Instance, error) {
	db := config.DB

	var insList []Instance
	result := db.Find(&insList)
	if result.Error != nil {
		return nil, result.Error
	}

	return insList, nil
}

// GetInstanceByName 根据实例名查询实例信息并返回
func (ins *Instance) GetInstanceByName() (*Instance, error) {
	var instance Instance
	db := config.DB

	result := db.Model(ins).Where("name = ?", ins.Name).First(&instance)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &instance, nil
}

func (ins *Instance) GetInstanceList(pageSize, currentPage int) ([]Instance, int64, error) {
	var insList []Instance
	var count int64

	db := config.DB

	limit := pageSize
	offset := pageSize * (currentPage - 1)

	// 获取分页数据
	result := db.Limit(limit).Offset(offset).Find(&insList)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	// 获取总记录数
	countResult := db.Model(&Instance{}).Count(&count)
	if result.Error != nil {
		return nil, 0, countResult.Error
	}

	return insList, count, nil
}
