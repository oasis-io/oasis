package model

type Instance struct {
	Model
	Name     string `json:"name" gorm:"index;not null;comment:实例名"`
	DBType   string `json:"db_type" gorm:"column:db_type;not null;comment:数据库类型" example:"mysql"`
	Host     string `json:"host" gorm:"not null;comment:数据库连接地址"`
	Port     string `json:"port" gorm:"not null;comment:数据库端口" example:"3306"`
	User     string `json:"user" gorm:"not null;comment:用户名" example:"root"`
	Password string `json:"-"  gorm:"not null;comment:密码"`
}
