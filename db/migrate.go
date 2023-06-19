package db

import (
	"oasis/config"
	"oasis/db/model"
	"oasis/pkg/log"
	"os"
	"sync"
)

var once sync.Once

func AutoMigrate() {
	db := config.DB

	err := db.AutoMigrate(
		&model.Instance{},
		&model.User{},
		&model.UserRole{},
		&model.UserGroup{},
		&model.Menu{},
		&model.RoleMenuRelation{},
		&model.Api{},
	)
	if err != nil {
		log.Error("Migrate table fail:" + err.Error())
		os.Exit(1)
	}
}
