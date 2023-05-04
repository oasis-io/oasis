package model

import "time"

type Model struct {
	ID         uint      `json:"id" gorm:"primary_key"`
	CreateTime time.Time `json:"create_time" gorm:"not null;default:current_timestamp;comment:创建时间"`
	UpdateTime time.Time `json:"update_time" gorm:"not null;default:current_timestamp on update current_timestamp;comment:修改时间"`
}
