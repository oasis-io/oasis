package app

import (
	"oasis/config"
	db2 "oasis/db"
	"oasis/pkg/log"
)

var err error

func RunServer() {
	// mysql connection pool
	log.Info("Connecting to MySQL")
	config.DB, err = db2.OpenOasis()
	if err != nil {
		panic(err)
	}

	log.Info("Starting Oasis server")
	startHttp()
}

func startHttp() {
	HttpRequests()
}
