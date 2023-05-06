package app

import (
	"oasis/config"
	"oasis/db"
	"oasis/pkg/log"
)

var err error

func RunServer() {
	// mysql connection pool
	log.Info("Connecting to MySQL")
	config.DB, err = db.OpenOasis()
	if err != nil {
		panic(err)
	}

	// 初始化表格
	db.AutoMigrate()

	log.Info("Starting Oasis server")
	startHttp()
}

func startHttp() {
	HttpRequests()
}
