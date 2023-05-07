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

		// migrate table
		err := db.AutoMigrate(
			&model.Instance{},
			&model.User{},
			&model.UserRole{},
			&model.UserGroup{},
			&model.Menu{},
		)
		if err != nil {
			panic(err)
		}
		log.Info("migrate table success")
	})
}
