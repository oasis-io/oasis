package model

type AccountManager struct {
	Model
	Name        string `json:"name" gorm:"column:name;index:uk_name,unique;not null;"`
	Permissions string `json:"permissions" gorm:"column:permissions"`
	Desc        string `json:"desc" gorm:"column:desc"`
}
