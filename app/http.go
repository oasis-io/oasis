package app

import (
	"oasis/config"
	"oasis/db"
	"oasis/pkg/casbin"
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

	casbin.InitCasbin()

	err := db.InsertData()
	if err != nil {
		log.Error("初始化数据失败:" + err.Error())
		panic(err)
	}

	log.Info("Starting Oasis server")
	startHttp()
}

func startHttp() {
	HttpRequests()
}
