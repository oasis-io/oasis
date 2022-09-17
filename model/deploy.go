package model

type Deploy struct {
	Model
	Name    string `json:"name" gorm:"index;not null;comment:实例名称"`
	DBType  string `json:"db_type" gorm:"column:db_type;not null;comment:数据库类型" example:"mysql"`
	Version string `json:"version" gorm:"not null;comment:版本"`
}
