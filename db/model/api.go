package model

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"oasis/config"
)

type Api struct {
	Model
	Group  string `json:"group" gorm:"column:group;not null;"`  // API业务划分
	Path   string `json:"path" gorm:"column:path;not null;"`    // api路径
	Method string `json:"method" gorm:"column:method;not null"` // 方法:POST、GET、PUT、DELETE
	Desc   string `json:"desc" gorm:"column:desc;"`             // api中文描述
}

func (api *Api) CreateApi() error {
	db := config.DB

	result := db.Create(api)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (api *Api) CreateMultipleApi(apis []Api) error {
	db := config.DB

	tx := db.Begin()
	for _, data := range apis {
		result := tx.Create(&data)
		if result.Error != nil {
			tx.Rollback()
			return result.Error
		}
	}
	tx.Commit()

	return nil
}

// 修改一条数据
func (api *Api) UpdateApi() error {
	db := config.DB

	result := db.Model(api).Updates(api)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// 删除一条数据
func (api *Api) DeleteApi() error {
	db := config.DB

	result := db.Delete(api)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (api *Api) DeleteAllApis() error {
	db := config.DB

	//result := db.Exec("DELETE FROM apis")
	result := db.Exec("truncate table apis")
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// 查询一条数据
func (api *Api) GetApi() (*Api, error) {
	db := config.DB

	result := db.First(api)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("no api found")
		}
		return nil, result.Error
	}

	return api, nil
}

func (api *Api) GetAllApi() ([]Api, error) {
	db := config.DB

	var apiList []Api
	result := db.Find(&apiList)
	if result.Error != nil {
		return nil, result.Error
	}

	return apiList, nil
}
