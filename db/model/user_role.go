package model

type UserRole struct {
	Model
	Name string `json:"name" gorm:"index:uk_name,unique;not null;"`
}
