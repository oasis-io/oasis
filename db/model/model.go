package model

import (
	"github.com/google/uuid"
	"time"
)

type Model struct {
	ID         uint      `json:"id" gorm:"primary_key"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time;not null;default:current_timestamp;comment:创建时间"`
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time;not null;default:current_timestamp on update current_timestamp;comment:修改时间"`
}

type UUIDModel struct {
	ID         uuid.UUID `json:"id" gorm:"primary_key"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time;not null;default:current_timestamp;comment:创建时间"`
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time;not null;default:current_timestamp on update current_timestamp;comment:修改时间"`
}
