package app

import (
	"go.uber.org/zap"
	"oasis/config"
	"oasis/db"
	"oasis/pkg/casbin"
	"oasis/pkg/log"
	"os"
)

func RunServer() {
	log.Info("Initializing the server")

	// Connect to database
	log.Info("Connecting to MySQL")
	var err error
	config.DB, err = db.OpenOasis()
	if err != nil {
		log.Error("Failed to connect to the database", zap.Error(err))
		os.Exit(1)
	}

	// 初始化表格
	db.AutoMigrate()

	casbin.InitCasbin()

	if err := db.InsertData(); err != nil {
		log.Error("初始化数据失败:" + err.Error())
		panic(err)
	}

	// Start the server
	log.Info("Starting Oasis server")
	HttpRequests()
}
