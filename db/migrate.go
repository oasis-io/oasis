package db

import (
	"oasis/config"
	"oasis/db/model"
)

func AutoMigrate() {
	db := config.DB

	// migrate table
	err := db.AutoMigrate(
		&model.Instance{},
		&model.User{},
		&model.UserRole{},
		&model.UserGroup{},
	)
	if err != nil {
		panic(err)
	}
}
