package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenOasis(user, password, host, port, database string) (db *gorm.DB, err error) {
	return openOasis(user, password, host, port, database)
}

func openOasis(user, password, host, port, database string) (db *gorm.DB, err error) {
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

	DB.SetMaxOpenConns(5)
	DB.SetMaxIdleConns(3)
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
