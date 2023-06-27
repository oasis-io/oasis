package model

type DatabaseManager struct {
	Model
	Name         string `json:"name" gorm:"column:name;index:uk_name,unique;not null;"`
	CharacterSet string `json:"character_set"  gorm:"column:character_set;"`
	Collation    string `json:"collation"  gorm:"column:collation;"`
	Desc         string `json:"desc" gorm:"column:desc"`
}
