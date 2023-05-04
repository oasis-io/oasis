package model

type Instance struct {
	Model
	Name     string `json:"name" gorm:"index:uk_name,unique;not null;"`
	DBType   string `json:"db_type" gorm:"column:db_type;not null;" example:"mysql"`
	Host     string `json:"host" gorm:"not null;"`
	Port     string `json:"port" gorm:"not null;" example:"3306"`
	User     string `json:"user" gorm:"not null;" example:"root"`
	Password string `json:"password"  gorm:"not null;"`
}
