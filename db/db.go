package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"oasis/config"
	"sync"
)

var mutex sync.Mutex
var mysqlURI string

//func getMySQLURI() string {
//	mutex.Lock()
//	defer mutex.Unlock()
//	if mysqlURI != "" {
//		return mysqlURI
//	}
//
//}

func OpenOasis() (db *gorm.DB, err error) {
	return openOasis()
}

func openOasis() (db *gorm.DB, err error) {

	config := config.NewConfig()
	user := config.MySQL.User
	password := config.MySQL.Password
	host := config.MySQL.Host
	port := config.MySQL.Port
	database := config.MySQL.Database

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, database)

	mysqlConfig := mysql.Config{
		DriverName:                "mysql",
		DSN:                       dsn,
		DefaultStringSize:         255,
		SkipInitializeWithVersion: false,
		DisableDatetimePrecision:  true,
	}
	db, err = gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	DB, err := db.DB()
	if err != nil {
		return nil, err
	}

	maxOpen := config.MySQL.MaxOpenConns
	maxIdle := config.MySQL.MaxIdleConns
	DB.SetMaxOpenConns(maxOpen)
	DB.SetMaxIdleConns(maxIdle)
	if err := DB.Ping(); err != nil {
		return nil, err
	}

	return db, err
}

func OpenInstance() (db *gorm.DB, err error) {
	return openInstance()
}

func openInstance() (db *gorm.DB, err error) {
	return db, err
}
