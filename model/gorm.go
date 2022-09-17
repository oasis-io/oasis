package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"oasis/config"
	"oasis/log"
	"os"
	"time"
)

type Model struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	CreateTime time.Time      `json:"create_time" gorm:"index;not null;default:current_timestamp;comment:创建时间"`
	UpdateTime time.Time      `json:"update_time" gorm:"not null;default:current_timestamp on update current_timestamp;comment:修改时间"`
	DeleteTime gorm.DeletedAt `json:"delete_time gorm:"index""`
}

func Migrate(config *config.Config) {
	db := myConn(config)
	AutoMigrate(db)
}

func AutoMigrate(db *gorm.DB) {
	// 迁移表结构
	err := db.AutoMigrate(
		&Instance{})

	if err != nil {
		os.Exit(1)
	}
}

func myConn(config *config.Config) *gorm.DB {
	user := config.DBConfig.User
	password := config.DBConfig.Password
	host := config.DBConfig.Host
	port := config.DBConfig.Port
	database := config.DBConfig.Database

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, database)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DriverName:                "mysql",
		DSN:                       dsn,
		DefaultStringSize:         255,
		SkipInitializeWithVersion: false,
		DisableDatetimePrecision:  true,
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Logger().Errorf("failed to connect database, error: %v", err)
	} else {
		log.Logger().Infoln("connect mysql success")
	}

	return db
}
