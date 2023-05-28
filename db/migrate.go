package db

import (
	"oasis/config"
	"oasis/db/model"
	"oasis/pkg/log"
	"sync"
)

var once sync.Once

func AutoMigrate() {
	once.Do(func() {
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
			log.Error("migrate table fail:" + err.Error())
			panic(err)
		}
		log.Info("migrate table success")
	})
}
