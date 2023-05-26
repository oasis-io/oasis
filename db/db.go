package db

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"oasis/config"
	"oasis/db/model"
	"oasis/pkg/log"
)

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

	maxOpen := config.MySQL.MaxOpenConn
	maxIdle := config.MySQL.MaxIdleConn
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

func GetMenuTree() ([]model.Menu, error) {
	var menus []model.Menu
	// 获取所有的菜单项
	db := config.DB
	if err := db.Find(&menus).Error; err != nil {
		return nil, err
	}

	// 构建菜单树
	return buildMenuTree(menus, "0"), nil
}

func buildMenuTree(menus []model.Menu, parentId string) []model.Menu {
	var result []model.Menu
	for _, menu := range menus {
		// 将 menu.ID 转换为 string 来比较
		if menu.ParentID == parentId {
			children := buildMenuTree(menus, fmt.Sprintf("%d", menu.ID))
			menu.Children = children
			result = append(result, menu)
		}
	}

	return result
}

func Login(username, password string) (*model.User, error) {
	db := config.DB

	user := model.User{}

	if err := db.Where("username = ? AND password = ?", username, password).Preload("Roles").First(&user).Error; err != nil {
		log.Error("获取用户角色错误", zap.Error(err))
		return nil, err
	}

	return &user, nil
}
